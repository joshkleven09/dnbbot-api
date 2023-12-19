package hello

import (
	"github.com/google/uuid"
)

type HelloApi struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Hello struct {
	ID   uuid.UUID `gorm:"primarykey"`
	Name string
}

type Hellos []*Hello

func (h *Hello) ToApi() *HelloApi {
	return &HelloApi{
		ID:   h.ID.String(),
		Name: h.Name,
	}
}

func (hs Hellos) ToApi() []*HelloApi {
	apis := make([]*HelloApi, len(hs))
	for i, v := range hs {
		apis[i] = v.ToApi()
	}

	return apis
}
