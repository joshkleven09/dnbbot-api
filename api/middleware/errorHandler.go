package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ErrorHandler(logger *zerolog.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		lastError := context.Errors.Last()
		if lastError == nil {
			return
		}

		for _, err := range context.Errors {
			logger.Error().Err(err).Msg("")
		}

		context.JSON(-1, "")
	}

}
