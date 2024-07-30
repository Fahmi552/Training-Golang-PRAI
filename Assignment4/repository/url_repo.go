package repository

import (
	"context"

	"Assignment4/entity"

	"gorm.io/gorm"
)

// GormDBIface defines an interface for GORM DB methods used in the repository
type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type URLRepository interface {
	CreateURL(url *entity.URL) error
	GetURLByShort(shortURL string) (*entity.URL, error)
}

type urlRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{db}
}

func (r *urlRepository) CreateURL(url *entity.URL) error {
	return r.db.Create(url).Error
}

func (r *urlRepository) GetURLByShort(shortURL string) (*entity.URL, error) {
	var url entity.URL
	if err := r.db.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}
