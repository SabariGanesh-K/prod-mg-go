package main

import (

	"github.com/SabariGanesh-K/prod-img-proc-service-go/imgprocessor"
	"github.com/SabariGanesh-K/prod-img-proc-service-go/util"
	"github.com/rs/zerolog/log"

	"github.com/streadway/amqp"
)

func main() {
    // Initialize logger
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
    // Initialize RabbitMQ connection
    rabbitMQConn, err := amqp.Dial(config.RabbitMQUrl)
    if err != nil {
        	log.Fatal().Err(err).Msg("Failed to connect to RabbitMQ")
    }
    defer rabbitMQConn.Close()

    ch, err := rabbitMQConn.Channel()
    if err != nil {
		log.Fatal().Err(err).Msg("Failed to open a channel" )
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "product_images", // name
        false,            // durable
        false,            // delete when unused
        false,            // exclusive
        false,            // no-wait
        nil,              // arguments
    )
    if err != nil {
        	log.Fatal().Err(err).Msg("Failed to declare a queue" )
    }




    // Consume messages from RabbitMQ
    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        	log.Fatal().Err(err).Msg("Failed to register a consumer" )
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            imgprocessor.ProcessImage(d.Body)
        }
    }()

    log.Info().Msg("Image processing microservice running")
    <-forever
}
