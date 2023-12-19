package hello

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"net/http"
)

type ApiService struct {
	logger     *zerolog.Logger
	validator  *validator.Validate
	repository *Repository
}

func New(logger *zerolog.Logger, validator *validator.Validate, db *gorm.DB) *ApiService {
	return &ApiService{
		logger:     logger,
		validator:  validator,
		repository: NewRepository(db),
	}
}

func (a *ApiService) Get(context *gin.Context) {
	hellos, err := a.repository.List()

	if err != nil {
		a.logger.Error().Err(err).Msg("")
		//e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	a.logger.Info().Msg("test log")
	context.IndentedJSON(http.StatusOK, hellos.ToApi())
}
