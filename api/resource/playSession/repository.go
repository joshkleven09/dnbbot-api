package playSession

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindAll() (PlaySessions, error) {
	playSessions := make([]*PlaySession, 0)
	if err := r.db.Joins("UserProfile").Joins("Guild").Find(&playSessions).Error; err != nil {
		return nil, err
	}

	return playSessions, nil
}

func (r *Repository) FindAllByGuildAndUserId(guildId uuid.UUID, userId uuid.UUID) (PlaySessions, error) {
	playSessions := make([]*PlaySession, 0)
	if err := r.db.Joins("UserProfile").Joins("Guild").Where("guild_id = ? && user_profile_id = ?", guildId, userId).Find(&playSessions).Error; err != nil {
		return nil, err
	}

	return playSessions, nil
}

func (r *Repository) FindAllByGuildId(guildId uuid.UUID) (PlaySessions, error) {
	playSessions := make([]*PlaySession, 0)
	if err := r.db.Joins("UserProfile").Joins("Guild").Where("guild_id = ?", guildId).Find(&playSessions).Error; err != nil {
		return nil, err
	}

	return playSessions, nil
}

func (r *Repository) FindAllByUserId(userId uuid.UUID) (PlaySessions, error) {
	playSessions := make([]*PlaySession, 0)
	if err := r.db.Joins("UserProfile").Joins("Guild").Where("user_profile_id = ?", userId).Find(&playSessions).Error; err != nil {
		return nil, err
	}

	return playSessions, nil
}

func (r *Repository) FindById(id uuid.UUID) (*PlaySession, error) {
	playSession := &PlaySession{}

	if err := r.db.Joins("UserProfile").Joins("Guild").First(&playSession, id).Error; err != nil {
		return nil, err
	}

	return playSession, nil
}

func (r *Repository) Create(playSession *PlaySession) (*PlaySession, error) {
	if err := r.db.Create(playSession).Error; err != nil {
		return nil, err
	}

	return playSession, nil
}
