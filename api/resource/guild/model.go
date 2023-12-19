package guild

import (
	"dnbbot-api/api/resource"
	"time"
)

type GuildApi struct {
	GuildId    string    `json:"guild_id"`
	ExternalId string    `json:"external_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GuildCreateApi struct {
	ExternalId string `json:"external_id" form:"required"`
	Name       string `json:"name" form:"required"`
}

type Guild struct {
	resource.Model
	ExternalId string
	Name       string
}

type Guilds []*Guild

func (g *Guild) ToApi() *GuildApi {
	return &GuildApi{
		GuildId:    g.ID.String(),
		ExternalId: g.ExternalId,
		Name:       g.Name,
		CreatedAt:  g.CreatedAt,
		UpdatedAt:  g.UpdatedAt,
	}
}

func (guilds Guilds) ToApi() []*GuildApi {
	apis := make([]*GuildApi, len(guilds))
	for i, v := range guilds {
		apis[i] = v.ToApi()
	}

	return apis
}

func (g *GuildCreateApi) ToModel() *Guild {
	return &Guild{
		ExternalId: g.ExternalId,
		Name:       g.Name,
	}
}
