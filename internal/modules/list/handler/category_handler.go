package handler

import (
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/usecase/category_create"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/usecase/category_find"
	"github.com/cristiano-pacheco/shoplist/internal/shared/sdk/empty"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	createCategoryUseCase category_create.CategoryCreateUseCaseI
	findCategoriesUseCase category_find.CategoryFindUseCaseI
}

func NewCategoryHandler(
	createCategoryUseCase category_create.CategoryCreateUseCaseI,
	findCategoriesUseCase category_find.CategoryFindUseCaseI,
) *CategoryHandler {
	return &CategoryHandler{createCategoryUseCase, findCategoriesUseCase}
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

	input := category_create.Input{UserID: userID, Name: request.Name}
	output, err := h.createCategoryUseCase.Execute(c.UserContext(), input)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
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

	input := category_find.Input{UserID: userID}

	if !empty.IsEmpty(request.CategoryID) {
		input.CategoryID = &request.CategoryID
	}

	if !empty.IsEmpty(request.CategoryName) {
		input.CategoryName = &request.CategoryName
	}

	output, err := h.findCategoriesUseCase.Execute(c.UserContext(), input)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	res := dto.CategoryFindResponse{
		Categories: make([]dto.Category, len(output.Categories)),
	}

	for i, category := range output.Categories {
		res.Categories[i] = dto.Category{
			ID:   category.ID,
			Name: category.Name,
		}
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	return nil
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	return nil
}
