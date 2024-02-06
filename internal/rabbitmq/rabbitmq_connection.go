package rabbitmq

import (
	"encoding/json"

	"github.com/leonardohklaus/goapi/internal/entity"
	"github.com/streadway/amqp"
)

func PublishRabbitMQMessage(product *entity.Product) error {
	connectRabbitMQ, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	product_json, _ := json.Marshal(product)

	messageMQ := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(string(product_json)),
	}

	// Attempt to publish a message to the queue.
	if err := channelRabbitMQ.Publish(
		"amq.direct",            // exchange
		"NewProduct", 			// Routing Key
		false,           		// mandatory
		false,           		// immediate
		messageMQ,         		// message to publish
	); err != nil {
		return err
	}
	return nil;
}	
