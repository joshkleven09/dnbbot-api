package userProfile

import (
	"dnbbot-api/api/resource"
	"time"
)

type UserApi struct {
	UserId     string    `json:"user_id"`
	ExternalId string    `json:"external_id"`
	Name       string    `json:"name"`
	Timezone   string    `json:"timezone"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserCreateApi struct {
	ExternalId string `json:"external_id" form:"required"`
	Name       string `json:"name" form:"required"`
	Timezone   string `json:"timezone"`
}

type UserProfile struct {
	resource.Model
	ExternalId string
	Name       string
	Timezone   string
}

type UserProfiles []*UserProfile

func (u *UserProfile) ToApi() *UserApi {
	return &UserApi{
		UserId:     u.ID.String(),
		ExternalId: u.ExternalId,
		Name:       u.Name,
		Timezone:   u.Timezone,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (users UserProfiles) ToApi() []*UserApi {
	apis := make([]*UserApi, len(users))
	for i, v := range users {
		apis[i] = v.ToApi()
	}

	return apis
}

func (u *UserCreateApi) ToModel() *UserProfile {
	return &UserProfile{
		ExternalId: u.ExternalId,
		Name:       u.Name,
		Timezone:   u.Timezone,
	}
}
