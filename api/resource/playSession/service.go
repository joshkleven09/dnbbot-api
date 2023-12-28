package playSession

import (
	"dnbbot-api/api/resource"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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

func (a *ApiService) GetPlaySessions(guildId string, userId string, date string, timeFilterStartStr string, timeFilterEndStr string) (PlaySessions, error) {
	var playSessions PlaySessions
	var err error
	var timeFilterStart time.Time
	var timeFilterEnd time.Time

	if (timeFilterStartStr != "" && timeFilterEndStr == "") || (timeFilterStartStr == "" && timeFilterEndStr != "") {
		return nil, &resource.ValidationError{Message: "both start and end time filters must be provided"}
	}

	if timeFilterStartStr != "" && timeFilterEndStr != "" {
		timeFilterStart, _ = time.Parse(time.RFC3339, timeFilterStartStr)
		timeFilterEnd, _ = time.Parse(time.RFC3339, timeFilterEndStr)
	}

	if guildId != "" && userId != "" {
		playSessions, err = a.repository.FindAllByGuildAndUserId(guildId, userId)
	} else if guildId != "" {
		playSessions, err = a.repository.FindAllByGuildId(guildId, date, timeFilterStart, timeFilterEnd)
	} else if userId != "" {
		playSessions, err = a.repository.FindAllByUserId(userId)
	} else {
		playSessions, err = a.repository.FindAll()
	}

	return playSessions, err
}

func (a *ApiService) CreatePlaySession(playSessionCreateApi PlaySessionCreateApi) (*PlaySession, error) {
	newPlaySession := playSessionCreateApi.ToModel()

	playSession, err := a.repository.Create(newPlaySession)

	if err == nil {
		a.logger.Info().Msg("new play_session created")
	}

	return playSession, err
}
