package router

import "github.com/cristiano-pacheco/go-modulith/internal/module/billing/handler"

func RegisterBillingHandlers(r *Router, billingHandler *handler.BillingHandler) {
	r.Get("/billings", billingHandler.Index)
	r.Post("/billings", billingHandler.Store)
	r.Get("/billings/:id", billingHandler.Show)
	r.Put("/billings/:id", billingHandler.Update)
	r.Delete("/billings/:id", billingHandler.Delete)
}
