package guildConfig

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Api struct {
	GuildConfigId   primitive.ObjectID `json:"guild_config_id"`
	ExternalGuildId string             `json:"external_guild_id"`
	GuildName       string             `json:"guild_name"`
	DefaultChannel  string             `json:"default_channel"`
	CreatedAt       time.Time          `json:"created_at"`
	LastUpdatedAt   time.Time          `json:"last_updated_at"`
}

type CreateApi struct {
	ExternalGuildId string `json:"external_guild_id" form:"required"`
	GuildName       string `json:"guild_name" form:"required"`
	DefaultChannel  string `json:"default_channel"`
}

type Model struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ExternalGuildId string             `bson:"external_guild_id"`
	GuildName       string             `bson:"guild_name"`
	DefaultChannel  string             `bson:"default_channel"`
	CreatedAt       time.Time          `bson:"created_at"`
	LastUpdatedAt   time.Time          `bson:"last_updated_at"`
}

type Models []*Model

func (p *Model) ToApi() *Api {
	return &Api{
		GuildConfigId:   p.ID,
		ExternalGuildId: p.ExternalGuildId,
		GuildName:       p.GuildName,
		DefaultChannel:  p.DefaultChannel,
		CreatedAt:       p.CreatedAt,
		LastUpdatedAt:   p.LastUpdatedAt,
	}
}

func (models Models) ToApi() []*Api {
	apis := make([]*Api, len(models))
	for i, v := range models {
		apis[i] = v.ToApi()
	}

	return apis
}

func (p *CreateApi) ToModel() *Model {
	now := time.Now()
	return &Model{
		ExternalGuildId: p.ExternalGuildId,
		GuildName:       p.GuildName,
		DefaultChannel:  p.DefaultChannel,
		CreatedAt:       now,
		LastUpdatedAt:   now,
	}
}
