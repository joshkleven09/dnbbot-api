package guildConfig

import (
	validatorUtil "dnbbot-api/util/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type ApiService struct {
	logger     *zerolog.Logger
	validator  *validator.Validate
	repository *Repository
}

func New(logger *zerolog.Logger, validator *validator.Validate, db *mongo.Database) *ApiService {
	return &ApiService{
		logger:     logger,
		validator:  validator,
		repository: NewRepository(db),
	}
}

func (a *ApiService) HandleGetGuildConfig(context *gin.Context) {
	externalGuildId := context.Query("externalGuildId")

	guildConfigs, err := a.repository.FindAll(externalGuildId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, guildConfigs.ToApi())
}

func (a *ApiService) HandleCreateGuildConfig(context *gin.Context) {
	createApi := &CreateApi{}

	if err := context.BindJSON(&createApi); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := a.validator.Struct(createApi); err != nil {
		err := validatorUtil.ToErrResponse(err)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, "error validating guild config creation")
			return
		}
	}

	createModel := createApi.ToModel()

	insertCount, err := a.repository.Create(createModel)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		a.logger.Info().Msg("new guild config created")
	}

	if insertCount == 0 {
		a.logger.Info().Msg("guild config updated")
		context.Status(http.StatusOK)
	} else {
		a.logger.Info().Msg("new guild config created")
		context.Status(http.StatusCreated)
	}
}

//func (a *ApiService) HandleUpdateGuildConfig(context *gin.Context) {
//
//	context.IndentedJSON(http.StatusOK, models.ToApi())
//}

func (a *ApiService) HandleDeleteGuildConfig(context *gin.Context) {
	guildConfigId := context.Param("id")

	err := a.repository.Delete(guildConfigId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	context.Status(http.StatusNoContent)
}
