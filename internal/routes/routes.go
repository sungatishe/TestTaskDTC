package routes

import (
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	r chi.Router
}

func NewRoutes(r chi.Router) *Routes {
	return &Routes{r: r}
}

func (rt *Routes) SetupOrderRoutes(orderHandler OrderHandlerInterface) {
	rt.r.Get("/orders", orderHandler.GetOrdersByFilters)
	rt.r.Get("/orders/{id}", orderHandler.GetOrderByID)
	rt.r.Post("/orders", orderHandler.CreateOrder)
	rt.r.Put("/orders/{id}", orderHandler.UpdateOrder)
	rt.r.Delete("/orders/{id}", orderHandler.DeleteOrder)
}

func (rt *Routes) SetupProductRoutes(productHandler ProductHandlerInterface) {
	rt.r.Get("/products", productHandler.GetAllProducts)
	rt.r.Get("/products/{id}", productHandler.GetProductByID)
	rt.r.Post("/products", productHandler.CreateProduct)
	rt.r.Put("/products/{id}", productHandler.UpdateProduct)
	rt.r.Delete("/products/{id}", productHandler.DeleteProduct)
}

func (rt *Routes) SetupAuthRoutes(authHandler AuthHandlerInterface) {
	rt.r.Post("/register", authHandler.RegisterUser)
	rt.r.Post("/login", authHandler.LoginUser)
}
