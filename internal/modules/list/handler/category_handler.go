package handler

import (
	"github.com/cristiano-pacheco/go-modulith/internal/modules/list/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/modules/list/usecase/category/create_category_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/modules/list/usecase/category/find_category_usecase"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	createCategoryUseCase create_category_usecase.UseCaseI
	findCategoriesUseCase find_category_usecase.UseCaseI
}

func NewCategoryHandler(
	createCategoryUseCase create_category_usecase.UseCaseI,
	findCategoriesUseCase find_category_usecase.UseCaseI,
) *CategoryHandler {
	return &CategoryHandler{createCategoryUseCase, findCategoriesUseCase}
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var request dto.CreateCategoryRequest

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	input := create_category_usecase.Input{
		UserID: userID,
		Name:   request.Name,
	}

	output, err := h.createCategoryUseCase.Execute(c.UserContext(), input)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	res := dto.CreateCategoryResponse{
		ID:   output.CategoryModel.ID,
		Name: output.CategoryModel.Name,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	input := find_category_usecase.Input{
		UserID: userID,
	}

	output, err := h.findCategoriesUseCase.Execute(c.UserContext(), input)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
