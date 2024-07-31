package handlers

import (
	"github.com/AlexEr256/messageService/repositories"
	"github.com/gofiber/fiber/v2"
)

type IConsumerHandler interface {
	GetStat(c *fiber.Ctx) error
}

type ConsumerHandler struct {
	ConsumerRepository repositories.IConsumerRepository
}

func NewConsumerHandler(repository repositories.IConsumerRepository) IConsumerHandler {
	return &ConsumerHandler{ConsumerRepository: repository}
}

func (h ConsumerHandler) GetStat(c *fiber.Ctx) error {
	result, err := h.ConsumerRepository.Get()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get information about entities"})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
