package handler

import "github.com/cristiano-pacheco/go-modulith/internal/module/billing/usecase"

type BillingHandler struct {
	createBillingUseCase *usecase.CreateBillingUseCase
}

func NewBillingHandler(createBillingUseCase *usecase.CreateBillingUseCase) *BillingHandler {
	return &BillingHandler{createBillingUseCase}
}

func (h *BillingHandler) Index() {
}

func (h *BillingHandler) Store() {
}

func (h *BillingHandler) Update() {
}

func (h *BillingHandler) Delete() {
}

func (h *BillingHandler) Show() {
}
