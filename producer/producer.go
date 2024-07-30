package producer

import "github.com/gofiber/fiber"

func ProducerService() {
	app := fiber.New()

	app.Get("/producer", func(c *fiber.Ctx) {
		c.Send("Producer service")
	})

	app.Listen(3000)
}
