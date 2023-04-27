// Package main is the entry point for the backend of Mentor Management System application.
//
// This package is responsible for initializing the application and starting the server.
// It loads the configuration file, establishes the database and Redis connections,
// creates the Gin server and the task processor for asynchronous tasks.
//
// The main function calls the connectDB function to establish a connection to MongoDB,
// then it creates a new MongoDB store using the connection, and creates a new Redis task distributor.
// It then launches the task processor in a goroutine and starts the Gin server.
//
// The package also provides two helper functions: connectDB and closeDB, which are responsible for
// establishing and closing the database connection respectively.
package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os"

	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/cmd/api"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/mongodb"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/cache"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/mail"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/worker"
	"github.com/hibiken/asynq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//go:embed docs
var staticFiles embed.FS

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("could not load config file")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Caller().Logger()
	}

	conn, err := connectDB(config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	log.Info().Msg("database connection established")
	defer closeDB(conn)
	store := mongodb.NewMongoClient(conn)

	cache, err := cache.NewRedisCache(config.RedisAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to redis cache")
	}
	log.Info().Msg("redis cache connection established")

	asynqRedisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
	taskDistributor := worker.NewRedisTaskDistributor(asynqRedisOpt)

	go runTaskProcessor(config, asynqRedisOpt, store)
	runGinServer(config, store, taskDistributor, cache)
}

func runGinServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor, cache cache.Cache) {

	fsys, err := fs.Sub(staticFiles, "docs/swagger-ui")
	if err != nil {
		log.Fatal().Err(err).Msg("can get swagger-ui static files")
	}

	server, err := api.NewServer(config, store, taskDistributor, fsys, cache)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	if err := server.Start(fmt.Sprint(config.HTTPServerAddress)); err != nil {
		log.Fatal().Err(err).Msg("error occur starting server")
	}
}

func runTaskProcessor(config utils.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("starting task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

// connectDB establishes connection to MongoDB
func connectDB(connURI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(connURI).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, err
}

// closeDB close database connection.
func closeDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client.Disconnect(ctx)
}
