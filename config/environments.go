package environments

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Environment      = ""
	Port             = 0
	ConnectionString = ""
	BootstrapServer  = ""
	GroupId          = ""
	Topic            = ""
	SlackBaseUrl     = ""
)

func NewConfig() {
	var err error

	viper.SetConfigName("../../config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Error during reading config", err)
		log.Fatal(err)
	}

	Environment = viper.GetString("environment")
	Port = viper.GetInt("producer.port")
	ConnectionString = viper.GetString("producer.connectionString")
}
