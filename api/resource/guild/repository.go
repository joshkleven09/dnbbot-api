package guild

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

func (r *Repository) FindAll() (Guilds, error) {
	guilds := make([]*Guild, 0)
	if err := r.db.Find(&guilds).Error; err != nil {
		return nil, err
	}

	return guilds, nil
}

func (r *Repository) FindById(id uuid.UUID) (*Guild, error) {
	guild := &Guild{}
	if err := r.db.First(&guild, id).Error; err != nil {
		return nil, err
	}

	return guild, nil
}

func (r *Repository) Create(guild *Guild) (*Guild, error) {
	if err := r.db.Create(guild).Error; err != nil {
		return nil, err
	}

	return guild, nil
}

func (r *Repository) GetOrCreate(guild Guild) (*Guild, error) {
	if err := r.db.Where("external_id = ?", guild.ExternalId).FirstOrCreate(&guild).Error; err != nil {
		return nil, err
	}

	return &guild, nil
}
