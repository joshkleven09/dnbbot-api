package guild

import (
	validatorUtil "dnbbot-api/util/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

func (a *ApiService) HandleGetGuilds(context *gin.Context) {
	guilds, err := a.repository.FindAll()

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, guilds.ToApi())
}

func (a *ApiService) HandleGetGuild(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	guild, err := a.repository.FindById(id)

	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	context.IndentedJSON(http.StatusOK, guild.ToApi())
}

func (a *ApiService) HandleCreateGuild(context *gin.Context) {
	guildCreateApi := &GuildCreateApi{}

	if err := context.BindJSON(&guildCreateApi); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := a.validator.Struct(guildCreateApi); err != nil {
		err := validatorUtil.ToErrResponse(err)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, "{}")
			return
		}
	}

	newGuild := guildCreateApi.ToModel()

	guild, err := a.repository.Create(newGuild)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	a.logger.Info().Str("id", guild.ID.String()).Msg("new guild created")
	context.IndentedJSON(http.StatusCreated, guild.ToApi())
}
