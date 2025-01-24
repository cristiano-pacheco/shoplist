package handler

import (
	"strconv"

	"github.com/cristiano-pacheco/shoplist/internal/modules/list/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list/usecase"
	"github.com/gofiber/fiber/v2"
)

type ListHandler struct {
	createListUseCase *usecase.ListCreateUseCase
	findListUsecase   *usecase.ListFindUsecase
	updateListUseCase *usecase.ListUpdateUseCase
	deleteListUseCase *usecase.ListDeleteUseCase
}

func NewListHandler(
	createListUseCase *usecase.ListCreateUseCase,
	findListUsecase *usecase.ListFindUsecase,
	updateListUseCase *usecase.ListUpdateUseCase,
	deleteListUseCase *usecase.ListDeleteUseCase,
) *ListHandler {
	return &ListHandler{
		createListUseCase,
		findListUsecase,
		updateListUseCase,
		deleteListUseCase,
	}
}

func (h *ListHandler) Create(c *fiber.Ctx) error {
	var request dto.ListCreateRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	input := usecase.ListCreateInput{UserID: userID, Name: request.Name}
	output, err := h.createListUseCase.Execute(c.UserContext(), input)
	if err != nil {
		return err
	}

	res := dto.ListCreateResponse{
		List: dto.List{
			ID:   output.ListModel.ID,
			Name: output.ListModel.Name,
		},
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *ListHandler) FindByUserID(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var request dto.ListFindRequest
	err := c.BodyParser(&request)
	lists, err := h.findListUsecase.ExecuteByUserID(c.UserContext(), userID)
	if err != nil {
		return err
	}

	res := dto.ListFindResponse{
		Lists: make([]dto.List, len(lists)),
	}

	for i, list := range lists {
		res.Lists[i] = dto.List{
			ID:   list.ID,
			Name: list.Name,
		}
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ListHandler) Update(c *fiber.Ctx) error {
	listID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var request dto.ListUpdateRequest
	err = c.BodyParser(&request)
	if err != nil {
		return err
	}

	input := usecase.ListUpdateInput{
		ListID: uint64(listID),
		UserID: userID,
		Name:   request.Name,
	}

	err = h.updateListUseCase.Execute(c.UserContext(), input)
	if err != nil {
		return err
	}

	res := dto.ListUpdateResponse{
		List: dto.List{
			ID:   uint64(listID),
			Name: request.Name,
		},
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ListHandler) Delete(c *fiber.Ctx) error {
	listID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	input := usecase.ListDeleteInput{
		ListID: uint64(listID),
		UserID: userID,
	}

	err = h.deleteListUseCase.Execute(c.UserContext(), input)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *ListHandler) FindByID(c *fiber.Ctx) error {
	listID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	list, err := h.findListUsecase.ExecuteByIDAndUserID(c.UserContext(), listID, userID)
	if err != nil {
		return err
	}

	res := dto.List{
		ID:   list.ID,
		Name: list.Name,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
