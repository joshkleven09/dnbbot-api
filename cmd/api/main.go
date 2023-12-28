package main

import (
	"context"
	"dnbbot-api/api/router"
	"dnbbot-api/util/logger"
	"dnbbot-api/util/properties"
	"dnbbot-api/util/validator"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	properties.New()
	l := logger.New()
	v := validator.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("mongo.connection")))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mongoDb := client.Database(viper.GetString("mongo.database"))

	l.Info().Msg("Starting server...")

	router.New(l, v, mongoDb)
}
