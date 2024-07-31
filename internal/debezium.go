package internal

import (
	"bytes"
	"net/http"
)

var plan = []byte(`{
    "name": "default_connector",
    "config": {
      "connector.class": "io.debezium.connector.postgresql.PostgresConnector",
      "tasks.max": "1",
      "database.hostname": "postgres",
      "topic.prefix": "postgres",
      "database.port": "5432",
      "database.user": "postgres",
      "database.password": "postgres",
      "database.dbname": "postgres",
      "database.server.name": "postgres",
      "schema.include.list": "public",
      "table.include.list": "public.messages",
      "key.converter": "org.apache.kafka.connect.storage.StringConverter",
      "key.converter.schemas.enable": "false",
      "database.history.kafka.bootstrap.servers": "kafka:9092",
      "value.converter": "org.apache.kafka.connect.json.JsonConverter",
      "value.converter.schemas.enable": "false",
      "include.schema.changes": "true"
    }
  })`)

func CheckDebeziumConnector() error {
	response, err := http.Get("http://debezium_connect:8083/connectors/default_connector")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {

		_, err := http.Post("http://debezium_connect:8083/connectors/", "application/json", bytes.NewBuffer(plan))

		if err != nil {
			panic(err)
		}

		return nil
	}

	return err
}
