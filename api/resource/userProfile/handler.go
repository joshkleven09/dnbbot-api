package userProfile

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

func (a *ApiService) HandleGetUsers(context *gin.Context) {
	users, err := a.repository.FindAll()

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, users.ToApi())
}

func (a *ApiService) HandleGetUser(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := a.repository.FindById(id)

	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	context.IndentedJSON(http.StatusOK, user.ToApi())
}

func (a *ApiService) HandleCreateUser(context *gin.Context) {
	userCreateApi := &UserCreateApi{}

	if err := context.BindJSON(&userCreateApi); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := a.validator.Struct(userCreateApi); err != nil {
		err := validatorUtil.ToErrResponse(err)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, "{}")
			return
		}
	}

	newUserProfile := userCreateApi.ToModel()

	userProfile, err := a.repository.Create(newUserProfile)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	a.logger.Info().Str("id", userProfile.ID.String()).Msg("new user_profile created")
	context.IndentedJSON(http.StatusCreated, userProfile.ToApi())
}
