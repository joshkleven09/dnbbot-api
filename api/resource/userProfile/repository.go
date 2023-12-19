package userProfile

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

func (r *Repository) FindAll() (UserProfiles, error) {
	users := make([]*UserProfile, 0)
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) FindById(id uuid.UUID) (*UserProfile, error) {
	user := &UserProfile{}
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) Create(userProfile *UserProfile) (*UserProfile, error) {
	if err := r.db.Create(userProfile).Error; err != nil {
		return nil, err
	}

	return userProfile, nil
}

func (r *Repository) GetOrCreate(userProfile UserProfile) (*UserProfile, error) {
	if err := r.db.Where("external_id = ?", userProfile.ExternalId).FirstOrCreate(&userProfile).Error; err != nil {
		return nil, err
	}

	return &userProfile, nil
}
