package router

import (
	"dnbbot-api/api/middleware"
	"dnbbot-api/api/resource/guild"
	"dnbbot-api/api/resource/hello"
	"dnbbot-api/api/resource/playSession"
	"dnbbot-api/api/resource/userProfile"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func New(logger *zerolog.Logger, validator *validator.Validate, db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(middleware.ErrorHandler(logger))

	helloApi := hello.New(logger, validator, db)
	userApi := userProfile.New(logger, validator, db)
	guildApi := guild.New(logger, validator, db)
	playSessionApi := playSession.New(logger, validator, db)

	router.GET("/v1/hello", helloApi.Get)

	router.GET("/v1/users", userApi.HandleGetUsers)
	router.GET("/v1/users/:id", userApi.HandleGetUser)
	router.POST("/v1/users", userApi.HandleCreateUser)

	router.GET("/v1/guilds", guildApi.HandleGetGuilds)
	router.GET("/v1/guilds/:id", guildApi.HandleGetGuild)
	router.POST("/v1/guilds", guildApi.HandleCreateGuild)

	router.GET("/v1/sessions", playSessionApi.HandleGetPlaySessions)
	router.GET("/v1/sessions/:id", playSessionApi.HandleGetPlaySession)
	router.POST("/v1/sessions", playSessionApi.HandleCreatePlaySession)

	err := router.Run(":8080")
	if err != nil {
		return nil
	}

	return router
}
