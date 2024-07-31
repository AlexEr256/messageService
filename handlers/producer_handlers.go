package handlers

import (
	"github.com/AlexEr256/messageService/dto"
	"github.com/AlexEr256/messageService/repositories"
	"github.com/gofiber/fiber/v2"
)

type IProducerHandler interface {
	CreateMessage(c *fiber.Ctx) error
}

type ProducerHandler struct {
	ProducerRepository repositories.IProducerRepository
}

func NewProducerHandler(repository repositories.IProducerRepository) IProducerHandler {
	return &ProducerHandler{ProducerRepository: repository}
}

func (h ProducerHandler) CreateMessage(c *fiber.Ctx) error {
	request := &dto.MessageRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Invalid Entity"})
	}

	if request.Creator == "" || request.Recipient == "" || request.Mail == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Some fields are empty"})
	}

	response, err := h.ProducerRepository.Add(request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create message"})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
