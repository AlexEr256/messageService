package main

import (
	"fmt"
	"github.com/AlexEr256/messageService/database"
	"github.com/AlexEr256/messageService/handlers"
	"github.com/AlexEr256/messageService/internal"
	"github.com/AlexEr256/messageService/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := internal.CheckDebeziumConnector()

	if err != nil {
		panic(err)
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env file ", err)
		return
	}

	producerConnection := os.Getenv("PRODUCER_POSTGRES")
	con, err := database.NewConnection(producerConnection)
	if err != nil {
		fmt.Println("failed to connect to database ", err)
		return
	}
	r := repositories.NewProducerRepository(con.Db)
	h := handlers.NewProducerHandler(r)

	app := fiber.New()

	app.Group("/producer").Post("/messages", h.CreateMessage)

	app.Listen(":3000")

}
