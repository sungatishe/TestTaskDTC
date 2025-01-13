package routes

import (
	"TestTask/internal/middleware"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	r chi.Router
}

func NewRoutes(r chi.Router) *Routes {
	return &Routes{r: r}
}

func (rt *Routes) SetupOrderRoutes(orderHandler OrderHandlerInterface) {
	rt.r.Route("/orders", func(r chi.Router) {
		// Применение миддлвары для авторизации
		r.Use(middleware.AuthMiddleware)

		// Эндпоинты для роли User
		r.With(middleware.RoleMiddleware("User", "Admin")).Post("/", orderHandler.CreateOrder)
		r.With(middleware.RoleMiddleware("User", "Admin")).Get("/", orderHandler.GetOrdersByFilters)
		r.With(middleware.RoleMiddleware("User", "Admin")).Get("/{id}", orderHandler.GetOrderByID)
		r.With(middleware.RoleMiddleware("User", "Admin")).Put("/{id}", orderHandler.UpdateOrder)

		// Эндпоинты для роли Admin
		r.With(middleware.RoleMiddleware("Admin")).Delete("/{id}", orderHandler.DeleteOrder)
	})
}

func (rt *Routes) SetupProductRoutes(productHandler ProductHandlerInterface) {
	rt.r.Route("/products", func(r chi.Router) {
		// Применение миддлвары для авторизации
		r.Use(middleware.AuthMiddleware)

		// Эндпоинты для роли Admin
		r.With(middleware.RoleMiddleware("Admin")).Post("/", productHandler.CreateProduct)
		r.With(middleware.RoleMiddleware("Admin")).Put("/{id}", productHandler.UpdateProduct)
		r.With(middleware.RoleMiddleware("Admin")).Delete("/{id}", productHandler.DeleteProduct)

		// Эндпоинты, доступные всем
		r.Get("/", productHandler.GetAllProducts)
		r.Get("/{id}", productHandler.GetProductByID)
	})
}

func (rt *Routes) SetupAuthRoutes(authHandler AuthHandlerInterface) {
	rt.r.Post("/register", authHandler.RegisterUser)
	rt.r.Post("/login", authHandler.LoginUser)
}
