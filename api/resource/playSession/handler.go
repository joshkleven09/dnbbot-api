package playSession

import (
	validatorUtil "dnbbot-api/util/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (a *ApiService) HandleGetPlaySessions(context *gin.Context) {
	var guildId uuid.UUID
	var userId uuid.UUID

	guildIdStr := context.Query("guildId")
	userIdStr := context.Query("userId")

	if guildIdStr != "" {
		var parseErr error
		guildId, parseErr = uuid.Parse(guildIdStr)

		if parseErr != nil {
			context.AbortWithError(http.StatusBadRequest, parseErr)
			return
		}
	}

	if userIdStr != "" {
		var parseErr error
		userId, parseErr = uuid.Parse(userIdStr)

		if parseErr != nil {
			context.AbortWithError(http.StatusBadRequest, parseErr)
			return
		}
	}

	playSessions, err := a.GetPlaySessions(guildId, userId)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, playSessions.ToApi())
}

func (a *ApiService) HandleGetPlaySession(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	playSession, err := a.GetPlaySession(id)

	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	context.IndentedJSON(http.StatusOK, playSession.ToApi())
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
			context.AbortWithStatusJSON(http.StatusBadRequest, "{}")
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

func (a *ApiService) mergeUserGuild(playSession PlaySession) {
	
}
