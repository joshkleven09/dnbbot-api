package main

import (
	"dnbbot-api/api/router"
	"dnbbot-api/util/logger"
	"dnbbot-api/util/properties"
	"dnbbot-api/util/validator"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	properties.New()
	l := logger.New()
	v := validator.New()

	var logLevel gormlogger.LogLevel
	if strings.ToLower(viper.GetString("db.logging-level")) == "debug" {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(
		fmtDBString,
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.db-name"),
		viper.GetInt("db.port"),
	)

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{
		Logger:         gormlogger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		l.Fatal().Err(err).Msg("DB connection start failure")
		return
	}

	l.Info().Msg("Starting server...")

	router.New(l, v, db)
}
