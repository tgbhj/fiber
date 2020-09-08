package router

import (
	"fiber/handler"
	"github.com/gofiber/fiber"
)

// Route 路由
func Route(app *fiber.App) {
	api := app.Group("/api")  // /api

	v1 := api.Group("/v1")   // /api/v1

	v1.Post("/signUp", handler.SignUp)
	v1.Post("/signIn", handler.SignIn)
	v1.Get("/info", handler.GetInfo)
	v1.Post("/info", handler.PostInfo)
}