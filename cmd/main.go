package main

import (
	"github.com/AlexEr256/messageService/environments"
	"github.com/AlexEr256/messageService/internal"
	"github.com/AlexEr256/messageService/producer"
)

func init() {
	err := internal.CheckDebeziumConnector()

	if err != nil {
		panic(err)
	}
}

func main() {
	environments.NewConfig()
	producer.ProducerService()

}
