package services

import (
	"errors"
	"mime/multipart"

	"github.com/ecea-nitt/ecea-server/helpers"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/schemas"
	"github.com/ecea-nitt/ecea-server/utils"
)

type podcastService struct {
	repo repositories.PodcastRepository
}

type PodcastService interface {
	CreatePodcast(
		podcastDetails models.PodcastRequest,
		podcastImage *multipart.FileHeader) error
}

func NewPodcastService(repo repositories.PodcastRepository) PodcastService {
	return &podcastService{repo}
}

func (ps *podcastService) CreatePodcast(
	podcast models.PodcastRequest,
	podcastImage *multipart.FileHeader) error {

	name, err := utils.NameValidator(podcast.Name)
	if err != nil {
		return err
	}

	mediaURL, err := utils.URLValidator(podcast.MediaURL)
	if err != nil {
		return err
	}

	assetChannel := make(chan int)
	typeChannel := make(chan int)

	go helpers.FetchPodcastTypeID(typeChannel, string(podcast.Type), ps.repo)

	go helpers.UploadAndFetchAssetID(assetChannel, podcastImage, ps.repo, "podcast")

	assetID := <-assetChannel
	podcastTypeID := <-typeChannel

	if assetID == -1 || podcastTypeID == -1 {
		return errors.New("error uploading asset")
	}

	newPodcast := schemas.Podcast{
		Name:        name,
		Description: podcast.Description,
		MediaURL:    mediaURL,
		ThumbnailID: uint(assetID),
		TypeID:      uint(podcastTypeID),
	}

	return ps.repo.InsertPodcast(newPodcast)
}
