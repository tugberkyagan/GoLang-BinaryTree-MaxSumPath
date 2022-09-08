package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-wrk/pkg/models"
	"go-wrk/pkg/services"
)

type BinaryTreeHandlerInterface interface {
	CalculateMaxPathSum(c *fiber.Ctx) error
}
type BinaryTreeHandlerStruct struct {
	Service services.BinaryTreeServiceInterface
}

func (h BinaryTreeHandlerStruct) CalculateMaxPathSum(c *fiber.Ctx) error {
	tree := &models.JsonTree{}

	if err := c.BodyParser(tree); err != nil {
		return err
	}

	response := h.Service.Main(*tree)

	return c.JSON(response)
}
