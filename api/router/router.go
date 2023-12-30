package router

import (
	"dnbbot-api/api/middleware"
	"dnbbot-api/api/resource/health"
	"dnbbot-api/api/resource/playSession"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(logger *zerolog.Logger, validator *validator.Validate, db *mongo.Database) *gin.Engine {
	router := gin.New()
	router.Use(middleware.ErrorHandler(logger))

	healthApi := health.New(logger, validator)
	playSessionApi := playSession.New(logger, validator, db)

	router.GET("/health", healthApi.HandleGetHealth)

	router.GET("/v1/sessions", playSessionApi.HandleGetPlaySessions)
	router.POST("/v1/sessions", playSessionApi.HandleCreatePlaySession)

	err := router.Run(":8080")
	if err != nil {
		return nil
	}

	return router
}
