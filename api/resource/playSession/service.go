package playSession

import (
	"dnbbot-api/api/resource/guild"
	"dnbbot-api/api/resource/userProfile"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type ApiService struct {
	logger     *zerolog.Logger
	validator  *validator.Validate
	repository *Repository
	userRepo   *userProfile.Repository
	guildRepo  *guild.Repository
}

func New(logger *zerolog.Logger, validator *validator.Validate, db *gorm.DB) *ApiService {
	return &ApiService{
		logger:     logger,
		validator:  validator,
		repository: NewRepository(db),
		userRepo:   userProfile.NewRepository(db),
		guildRepo:  guild.NewRepository(db),
	}
}

func (a *ApiService) GetPlaySessions(guildId uuid.UUID, userId uuid.UUID) (PlaySessions, error) {
	var playSessions PlaySessions
	var err error

	if guildId != uuid.Nil && userId != uuid.Nil {
		playSessions, err = a.repository.FindAllByGuildAndUserId(guildId, userId)
	} else if guildId != uuid.Nil {
		playSessions, err = a.repository.FindAllByGuildId(guildId)
	} else if userId != uuid.Nil {
		playSessions, err = a.repository.FindAllByUserId(userId)
	} else {
		playSessions, err = a.repository.FindAll()
	}

	return playSessions, err
}

func (a *ApiService) GetPlaySession(sessionId uuid.UUID) (*PlaySession, error) {
	return a.repository.FindById(sessionId)
}

func (a *ApiService) CreatePlaySession(playSessionCreateApi PlaySessionCreateApi) (*PlaySession, error) {
	u, err := a.userRepo.GetOrCreate(userProfile.UserProfile{
		ExternalId: playSessionCreateApi.ExternalUserId,
		Name:       playSessionCreateApi.UserName,
	})

	if err != nil {
		return nil, err
	}

	g, err := a.guildRepo.GetOrCreate(guild.Guild{
		ExternalId: playSessionCreateApi.ExternalGuildId,
		Name:       playSessionCreateApi.GuildName,
	})

	if err != nil {
		return nil, err
	}

	newPlaySession := playSessionCreateApi.ToModel(u.ID, g.ID)

	playSession, err := a.repository.Create(newPlaySession)

	if err == nil {
		a.logger.Info().Str("id", playSession.ID.String()).Msg("new play_session created")
	}

	return playSession, err
}
