package repositories

import (
	"github.com/ecea-nitt/ecea-server/config"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/schemas"
	"gorm.io/gorm"
)

type podcastRepository struct {
	db *gorm.DB
}

type PodcastRepository interface {
	InsertPodcast(podcast schemas.Podcast) error
	InsertAsset(name string) (uint, error)
	FindPodcastByName(name string) (schemas.Podcast, error)
	FindPodcastTypeByName(name string) (schemas.PodcastType, error)
	FindPodcastByEpisodeNoAndType(episodeNo uint, podcastType string) (schemas.Podcast, error)
	FetchThumbnail(podcastID uint) (schemas.Asset, error)
	UpdatePodcastThumbnail(podcastID uint, assetID uint) error
	UpdatePodcastURL(podcastID uint, url string) error
	UpdatePodcastDescription(podcastID uint, description string) error
	DeletePodcast(podcastID uint) error
	GetAllPodcasts() ([]models.PodcastRequest, error)
	GetPodcastByType(podcastType string) ([]models.PodcastRequest, error)
	GetPodcastByEpisodeNoAndType(episodeNo uint, podcastType string) (models.PodcastRequest, error)
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

func (pr *podcastRepository) FindPodcastByName(name string) (schemas.Podcast, error) {
	var podcast schemas.Podcast
	if err := pr.db.Where("name = ?", name).First(&podcast).Error; err != nil {
		return schemas.Podcast{}, err
	}

	return podcast, nil
}

func (pr *podcastRepository) FindPodcastByEpisodeNoAndType(episodeNo uint, podcastType string) (schemas.Podcast, error) {
	var podcastTypeSchema schemas.PodcastType
	if err := pr.db.Where("name = ?", podcastType).First(&podcastTypeSchema).Error; err != nil {
		return schemas.Podcast{}, err
	}
	var podcast schemas.Podcast
	if err := pr.db.Where("episode_no = ? AND type_id = ?", episodeNo, podcastTypeSchema.ID).First(&podcast).Error; err != nil {
		return schemas.Podcast{}, err
	}
	return podcast, nil
}

func (pr *podcastRepository) FindPodcastTypeByName(name string) (schemas.PodcastType, error) {
	var podcastType schemas.PodcastType
	if err := pr.db.Where("name = ?", name).First(&podcastType).Error; err != nil {
		return schemas.PodcastType{}, err
	}

	return podcastType, nil
}

func (pr *podcastRepository) FetchThumbnail(podcastID uint) (schemas.Asset, error) {
	var podcast schemas.Podcast
	if err := pr.db.Preload("Thumbnail").Where("id = ?", podcastID).First(&podcast).Error; err != nil {
		return schemas.Asset{}, err
	}
	return podcast.Thumbnail, nil
}

func (pr *podcastRepository) UpdatePodcastThumbnail(podcastID uint, assetID uint) error {
	return pr.db.Model(&schemas.Podcast{}).Where("id = ?", podcastID).Update("thumbnail_id", assetID).Error
}

func (pr *podcastRepository) UpdatePodcastURL(podcastID uint, url string) error {
	return pr.db.Model(&schemas.Podcast{}).Where("id = ?", podcastID).Update("media_url", url).Error
}

func (pr *podcastRepository) UpdatePodcastDescription(podcastID uint, description string) error {
	return pr.db.Model(&schemas.Podcast{}).Where("id = ?", podcastID).Update("description", description).Error
}

func (pr *podcastRepository) DeletePodcast(podcastID uint) error {
	return pr.db.Unscoped().Delete(&schemas.Podcast{}, podcastID).Error
}

func (pr *podcastRepository) GetAllPodcasts() ([]models.PodcastRequest, error) {
	var podcasts []models.PodcastRequest

	if err := pr.db.Table("podcasts").Select(
		`podcasts.name,podcasts.description,podcasts.episode_no,podcasts.media_url,podcast_types.name as type,
		CONCAT(?::text,'/static/images','/',assets.name) as image_url`, config.Origin,
	).Joins(
		"JOIN assets on assets.id = podcasts.thumbnail_id").Joins(
		"JOIN podcast_types on podcast_types.id = podcasts.type_id").Scan(
		&podcasts).Error; err != nil {
		return nil, err
	}

	return podcasts, nil
}

func (pr *podcastRepository) GetPodcastByType(podcastType string) ([]models.PodcastRequest, error) {
	var podcasts []models.PodcastRequest
	if err := pr.db.Table("podcasts").Select(
		`podcasts.name,podcasts.description,podcasts.episode_no,podcasts.media_url,podcast_types.name as type,
		CONCAT(?::text,'/static/images','/',assets.name) as image_url`, config.Origin,
	).Where("podcast_types.name = ?", podcastType).Joins(
		"JOIN assets on assets.id = podcasts.thumbnail_id").Joins(
		"JOIN podcast_types on podcast_types.id = podcasts.type_id").Scan(
		&podcasts).Error; err != nil {
		return nil, err
	}

	return podcasts, nil
}

func (pr *podcastRepository) GetPodcastByEpisodeNoAndType(episodeNo uint, podcastType string) (models.PodcastRequest, error) {
	var podcasts models.PodcastRequest
	if err := pr.db.Table("podcasts").Select(
		`podcasts.name,podcasts.description,podcasts.episode_no,podcasts.media_url,podcast_types.name as type,
		CONCAT(?::text,'/static/images','/',assets.name) as image_url`, config.Origin,
	).Where("podcast_types.name = ? AND episode_no = ?", podcastType, episodeNo).Joins(
		"JOIN assets on assets.id = podcasts.thumbnail_id").Joins(
		"JOIN podcast_types on podcast_types.id = podcasts.type_id").Scan(
		&podcasts).Error; err != nil {
		return models.PodcastRequest{}, err
	}

	return podcasts, nil
	// var podcastTypeSchema schemas.PodcastType
	// if err := pr.db.Where("name = ?", podcastType).First(&podcastTypeSchema).Error; err != nil {
	// 	return schemas.Podcast{}, err
	// }
	// var podcast schemas.Podcast
	// if err := pr.db.Where("episode_no = ? AND type_id = ?", episodeNo, podcastTypeSchema.ID).First(&podcast).Error; err != nil {
	// 	return schemas.Podcast{}, err
	// }
	// return podcast, nil
}