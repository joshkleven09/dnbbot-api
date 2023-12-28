package playSession

import (
	"dnbbot-api/api/resource"
	validatorUtil "dnbbot-api/util/validator"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *ApiService) HandleGetPlaySessions(context *gin.Context) {
	playSessions, err := a.GetPlaySessions(
		context.Query("guildId"),
		context.Query("userId"),
		context.Query("date"),
		context.Query("timeFilterStart"),
		context.Query("timeFilterEnd"),
	)

	if err != nil {
		target := &resource.ValidationError{}
		if errors.As(err, &target) {
			context.AbortWithError(http.StatusBadRequest, err)
		} else {
			context.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	context.IndentedJSON(http.StatusOK, playSessions.ToApi())
}

func (a *ApiService) HandleCreatePlaySession(context *gin.Context) {
	playSessionCreateApi := &PlaySessionCreateApi{}

	if err := context.BindJSON(&playSessionCreateApi); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := a.validator.Struct(playSessionCreateApi); err != nil {
		err := validatorUtil.ToErrResponse(err)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, "error validating")
			return
		}
	}

	playSession, err := a.CreatePlaySession(*playSessionCreateApi)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusCreated, playSession.ToApi())
}
