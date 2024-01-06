package router

import (
	"dnbbot-api/api/middleware"
	"dnbbot-api/api/resource/guildConfig"
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
	guildConfigApi := guildConfig.New(logger, validator, db)

	router.GET("/health", healthApi.HandleGetHealth)

	router.GET("/v1/sessions", playSessionApi.HandleGetPlaySessions)
	router.POST("/v1/sessions", playSessionApi.HandleCreatePlaySession)
	router.DELETE("/v1/sessions/:id", playSessionApi.HandleDeletePlaySession)

	router.GET("/v1/guild_configs", guildConfigApi.HandleGetGuildConfig)
	router.POST("/v1/guild_configs", guildConfigApi.HandleCreateGuildConfig)
	router.DELETE("/v1/guild_configs/:id", guildConfigApi.HandleDeleteGuildConfig)

	err := router.Run(":8080")
	if err != nil {
		return nil
	}

	return router
}
