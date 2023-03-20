package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ethfaucet/models"
	"github.com/nikola43/ethfaucet/services"
)

func Claim(context *fiber.Ctx) error {
	claimRequest := new(models.ClaimRequest)

	err := context.BodyParser(claimRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	claimRequest.IPAddress = context.Get("X-Real-Ip")
	fmt.Println("claimRequest.IPAddress", claimRequest.IPAddress)
	claimResponse, err := services.Claim(claimRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.JSON(claimResponse)
}
