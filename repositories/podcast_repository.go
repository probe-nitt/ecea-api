package repositories

import (
	"github.com/ecea-nitt/ecea-server/schemas"
	"gorm.io/gorm"
)

type podcastRepository struct {
	db *gorm.DB
}

type PodcastRepository interface {
	InsertPodcast(podcast schemas.Podcast) error
	InsertAsset(name string) (uint, error)
	FindPodcastTypeByName(name string) (uint, error)
}

func NewPodcastRepository(db *gorm.DB) PodcastRepository {
	return &podcastRepository{db}
}

func (pr *podcastRepository) InsertPodcast(podcast schemas.Podcast) error {
	return pr.db.Create(&podcast).Error
}

func (pr *podcastRepository) InsertAsset(name string) (uint, error) {
	var assetType schemas.AssetType
	if err := pr.db.Where("name = ?", schemas.Image).First(&assetType).Error; err != nil {
		return 0, err
	}
	asset := schemas.Asset{
		Name:        name,
		AssetTypeID: assetType.ID,
	}
	if err := pr.db.Create(&asset).Error; err != nil {
		return 0, err
	}
	return asset.ID, nil
}

func (pr *podcastRepository) FindPodcastTypeByName(name string) (uint, error) {
	var podcastType schemas.PodcastType
	if err := pr.db.Where("name = ?", name).First(&podcastType).Error; err != nil {
		return 0, err
	}
	return podcastType.ID, nil
}
