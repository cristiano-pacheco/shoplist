package handler

import (
	"strconv"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/usecase"
	shared_errs "github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/response"
	"github.com/cristiano-pacheco/shoplist/internal/shared/sdk/empty"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	createCategoryUseCase *usecase.CategoryCreateUseCase
	findCategoriesUseCase *usecase.CategoryFindUseCase
	updateCategoryUseCase *usecase.CategoryUpdateUseCase
	deleteCategoryUseCase *usecase.CategoryDeleteUseCase
	errorMapper           shared_errs.ErrorMapper
}

func NewCategoryHandler(
	createCategoryUseCase *usecase.CategoryCreateUseCase,
	findCategoriesUseCase *usecase.CategoryFindUseCase,
	updateCategoryUseCase *usecase.CategoryUpdateUseCase,
	deleteCategoryUseCase *usecase.CategoryDeleteUseCase,
	errorMapper shared_errs.ErrorMapper,
) *CategoryHandler {
	return &CategoryHandler{
		createCategoryUseCase,
		findCategoriesUseCase,
		updateCategoryUseCase,
		deleteCategoryUseCase,
		errorMapper,
	}
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var request dto.CategoryCreateRequest

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	input := usecase.CategoryCreateInput{UserID: userID, Name: request.Name}
	output, err := h.createCategoryUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	res := dto.CategoryCreateResponse{
		Category: dto.Category{
			ID:   output.CategoryModel.ID,
			Name: output.CategoryModel.Name,
		},
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *CategoryHandler) Find(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var request dto.CategoryFindRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	input := usecase.CategoryFindInput{UserID: userID}

	if !empty.IsEmpty(request.CategoryID) {
		input.CategoryID = &request.CategoryID
	}

	if !empty.IsEmpty(request.CategoryName) {
		input.CategoryName = &request.CategoryName
	}

	output, err := h.findCategoriesUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	res := dto.CategoryFindResponse{
		Categories: make([]dto.Category, len(output.CategoryList)),
	}

	for i, category := range output.CategoryList {
		res.Categories[i] = dto.Category{
			ID:   category.ID,
			Name: category.Name,
		}
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	categoryID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var request dto.CategoryUpdateRequest
	err = c.BodyParser(&request)
	if err != nil {
		return err
	}

	input := usecase.CategoryUpdateInput{
		CategoryID: uint64(categoryID),
		UserID:     userID,
		Name:       request.Name,
	}

	err = h.updateCategoryUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	res := dto.CategoryUpdateResponse{
		Category: dto.Category{
			ID:   uint64(categoryID),
			Name: request.Name,
		},
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	categoryID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	input := usecase.CategoryDeleteInput{
		CategoryID: uint64(categoryID),
		UserID:     userID,
	}

	err = h.deleteCategoryUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
