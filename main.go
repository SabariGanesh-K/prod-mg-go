package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"database/sql"
 "github.com/streadway/amqp"
"fmt"
	// "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// "github.com/hibiken/asynq"
	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/SabariGanesh-K/prod-mgm-go/api"
	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"

	"github.com/SabariGanesh-K/prod-mgm-go/util"
	
	"golang.org/x/sync/errgroup"

)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	fmt.Printf("hello")
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	conn,err:= sql.Open(config.DBDriver,config.DBSource)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	rabbitMQConn, err := amqp.Dial(config.RabbitMQUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to RabbitMQ")


    }
	defer rabbitMQConn.Close()
	ch, err := rabbitMQConn.Channel()
    if err != nil {
		log.Fatal().Err(err).Msg("Failed to open a RabbitMQ  channel)")

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
		log.Fatal().Err(err).Msg("Failed to declare a RabbitMQ queue")

    }
	store := db.NewStore(conn)

	waitGroup, ctx := errgroup.WithContext(ctx)

	// runTaskProcessor(ctx, waitGroup, config, redisOpt, store)
	runGinServer(config,store,ch,q	)


	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}




func runGinServer(config util.Config, store db.Store ,rmqch *amqp.Channel,q amqp.Queue) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
