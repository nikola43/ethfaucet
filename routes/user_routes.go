package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ethfaucet/controllers"
)

func UserRoutes (router fiber.Router) {
	// /api/v1/auth/claim
	router.Post("/claim", controllers.Claim)
}
