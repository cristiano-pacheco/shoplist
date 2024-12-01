package handler

import (
	"github.com/cristiano-pacheco/go-modulith/internal/buylist/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/buylist/usecase/category/create_category_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/buylist/usecase/category/find_category_usecase"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	createCategoryUseCase create_category_usecase.UseCaseI
}

func NewCategoryHandler(
	createCategoryUseCase create_category_usecase.UseCaseI,
) *CategoryHandler {
	return &CategoryHandler{createCategoryUseCase}
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var request dto.CreateCategoryRequestDTO

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

	response := dto.CreateCategoryResponseDTO{
		ID:   output.CategoryModel.ID,
		Name: output.CategoryModel.Name,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	input := find_category_usecase.Input{
		UserID: userID,
	}

	output, err := h.getCategoriesUseCase.Execute(c.UserContext(), input)
}
