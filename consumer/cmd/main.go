package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AlexEr256/messageService/database"
	"github.com/AlexEr256/messageService/dto"
	"github.com/AlexEr256/messageService/handlers"
	"github.com/AlexEr256/messageService/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env file ", err)
		return
	}

	topic := os.Getenv("CONSUMER_TOPIC")
	consumerConnection := os.Getenv("CONSUMER_POSTGRES")

	con, err := database.NewConnection(consumerConnection)
	if err != nil {
		fmt.Println("failed to connect to database ", err)
		return
	}

	repo := repositories.NewConsumerRepository(con.Db)
	h := handlers.NewConsumerHandler(repo)

	app := fiber.New()

	app.Group("/consumer").Get("/total", h.GetStat)

	go app.Listen(":3100")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9093", "kafka:9092"},
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			continue
		}

		var message *dto.Payload
		if err = json.Unmarshal(m.Value, &message); err != nil {
			fmt.Println("Failed to parse incoming message")
			continue
		}

		messageRequest := &dto.MessageRequest{
			Creator:   message.After.Creator,
			Recipient: message.After.Recipient,
			Mail:      message.After.Mail,
		}

		_, err = repo.Add(messageRequest)
		if err != nil {
			fmt.Println("Failed to insert message in database")
		}

		fmt.Println("Inserted", messageRequest)
	}

	if err = reader.Close(); err != nil {
		fmt.Println("failed to close reader:", err)
	}
}
