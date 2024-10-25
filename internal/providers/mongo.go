package providers

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func initMongo(quit chan bool) {
	// Init Mongo
	uri := os.Getenv("MONGODBURI")
	log.Debug().Msgf("Mongo URI: %v", uri)

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	var err error = nil
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		// Kill the server if mongo connection fails
		log.Fatal().Msg("Failed to connect to MongoDB")
		return
	}
	DB = client.Database(os.Getenv("DATABASE"))

	log.Info().Msg("Connected to MongoDB")
	defer client.Disconnect(context.Background())

	<-quit
}
