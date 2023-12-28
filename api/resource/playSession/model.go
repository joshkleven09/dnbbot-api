package playSession

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PlaySessionApi struct {
	SessionId primitive.ObjectID `json:"session_id"`
	UserId    string             `json:"user_id"`
	Username  string             `json:"username"`
	GuildId   string             `json:"guild_id"`
	GuildName string             `json:"guild_name"`
	Date      string             `json:"date"`
	TimeRange string             `json:"time_range"`
	StartTime time.Time          `json:"start_time"`
	EndTime   time.Time          `json:"end_time"`
	Game      string             `json:"game"`
	IsPlayer  *bool              `json:"is_player"`
	CreatedAt time.Time          `json:"created_at"`
}

type PlaySessionCreateApi struct {
	ExternalUserId  string    `json:"external_user_id" form:"required"`
	Username        string    `json:"username" form:"required"`
	ExternalGuildId string    `json:"external_guild_id" form:"required"`
	GuildName       string    `json:"guild_name" form:"required"`
	Date            string    `json:"date" form:"required"`
	TimeRange       string    `json:"time_range" form:"required"`
	StartTime       time.Time `json:"start_time" form:"required"`
	EndTime         time.Time `json:"end_time" form:"required"`
	Game            string    `json:"game"`
	IsPlayer        *bool     `json:"is_player" form:"required"`
}

type PlaySession struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    string             `bson:"user_id"`
	Username  string             `bson:"username"`
	GuildId   string             `bson:"guild_id"`
	GuildName string             `bson:"guild_name"`
	Date      string             `bson:"date"`
	TimeRange string             `bson:"time_range"`
	StartTime time.Time          `bson:"start_time"`
	EndTime   time.Time          `bson:"end_time"`
	Game      string             `bson:"game"`
	IsPlayer  *bool              `bson:"is_player"`
	CreatedAt time.Time          `bson:"created_at"`
}

type PlaySessions []*PlaySession

func (p *PlaySession) ToApi() *PlaySessionApi {
	return &PlaySessionApi{
		SessionId: p.ID,
		UserId:    p.UserId,
		Username:  p.Username,
		GuildId:   p.GuildId,
		GuildName: p.GuildName,
		Date:      p.Date,
		TimeRange: p.TimeRange,
		StartTime: p.StartTime,
		EndTime:   p.EndTime,
		Game:      p.Game,
		IsPlayer:  p.IsPlayer,
		CreatedAt: p.CreatedAt,
	}
}

func (playSessions PlaySessions) ToApi() []*PlaySessionApi {
	apis := make([]*PlaySessionApi, len(playSessions))
	for i, v := range playSessions {
		apis[i] = v.ToApi()
	}

	return apis
}

func (p *PlaySessionCreateApi) ToModel() *PlaySession {
	return &PlaySession{
		UserId:    p.ExternalUserId,
		Username:  p.Username,
		GuildId:   p.ExternalGuildId,
		GuildName: p.GuildName,
		Date:      p.Date,
		TimeRange: p.TimeRange,
		StartTime: p.StartTime,
		EndTime:   p.EndTime,
		Game:      p.Game,
		IsPlayer:  p.IsPlayer,
		CreatedAt: time.Now(),
	}
}
