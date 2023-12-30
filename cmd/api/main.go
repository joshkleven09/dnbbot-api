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
	"os"
)

func main() {
	env := os.Getenv("DNBBOT_ENV")
	l := logger.New()
	properties.New(env, l)

	v := validator.New()

	var mongoConnectionStr string

	if env == "prod" {
		l.Info().Msg("Connecting to prod DB...")
		mongoConnectionStr = os.Getenv("DNBBOT_MONGO_CONN_STR")
	} else {
		l.Info().Msg("Connecting to local DB...")
		mongoConnectionStr = viper.GetString("mongo.connection")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConnectionStr))
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
