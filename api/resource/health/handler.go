package health

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

type ApiService struct {
	logger    *zerolog.Logger
	validator *validator.Validate
}

func New(logger *zerolog.Logger, validator *validator.Validate) *ApiService {
	return &ApiService{
		logger:    logger,
		validator: validator,
	}
}

func (a *ApiService) HandleGetHealth(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Health{
		Status: "up",
	})
}
