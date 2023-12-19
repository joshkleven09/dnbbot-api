package playSession

import (
	"dnbbot-api/api/resource"
	"dnbbot-api/api/resource/guild"
	"dnbbot-api/api/resource/userProfile"
	"github.com/google/uuid"
	"time"
)

type PlaySessionApi struct {
	SessionId string              `json:"session_id"`
	User      userProfile.UserApi `json:"user"`
	Guild     guild.GuildApi      `json:"guild"`
	StartTime time.Time           `json:"start_time"`
	EndTime   time.Time           `json:"end_time"`
	Game      string              `json:"game"`
	IsPlayer  *bool               `json:"is_player"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

type PlaySessionCreateApi struct {
	ExternalUserId  string    `json:"external_user_id" form:"required"`
	UserName        string    `json:"user_name" form:"required"`
	ExternalGuildId string    `json:"external_guild_id" form:"required"`
	GuildName       string    `json:"guild_name" form:"required"`
	StartTime       time.Time `json:"start_time" form:"required"`
	EndTime         time.Time `json:"end_time" form:"required"`
	Game            string    `json:"game"`
	IsPlayer        *bool     `json:"is_player" form:"required"`
}

type PlaySession struct {
	resource.Model
	UserProfileID uuid.UUID
	UserProfile   userProfile.UserProfile `gorm:"foreignKey:UserProfileID"`
	GuildID       uuid.UUID
	Guild         guild.Guild `gorm:"foreignKey:GuildID"`
	StartTime     time.Time
	EndTime       time.Time
	Game          string
	IsPlayer      *bool
}

type PlaySessions []*PlaySession

func (p *PlaySession) ToApi() *PlaySessionApi {
	return &PlaySessionApi{
		SessionId: p.ID.String(),
		User:      *p.UserProfile.ToApi(),
		Guild:     *p.Guild.ToApi(),
		StartTime: p.StartTime,
		EndTime:   p.EndTime,
		Game:      p.Game,
		IsPlayer:  p.IsPlayer,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (playSessions PlaySessions) ToApi() []*PlaySessionApi {
	apis := make([]*PlaySessionApi, len(playSessions))
	for i, v := range playSessions {
		apis[i] = v.ToApi()
	}

	return apis
}

func (p *PlaySessionCreateApi) ToModel(userUuid uuid.UUID, guildUuid uuid.UUID) *PlaySession {
	return &PlaySession{
		UserProfileID: userUuid,
		GuildID:       guildUuid,
		StartTime:     p.StartTime,
		EndTime:       p.EndTime,
		Game:          p.Game,
		IsPlayer:      p.IsPlayer,
	}
}
