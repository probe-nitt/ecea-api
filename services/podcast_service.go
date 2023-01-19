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
	EditThumbnail(
		podcastDetails models.PodcastRequest,
		podcastImage *multipart.FileHeader) error

	EditURL(podcastDetails models.PodcastRequest) error
	EditDescription(podcastDetails models.PodcastRequest) error
	DeletePodcast(podcastDetails models.PodcastRequest) error
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
		EpisodeNo:   podcast.EpisodeNo,
		ThumbnailID: uint(assetID),
		TypeID:      uint(podcastTypeID),
	}

	return ps.repo.InsertPodcast(newPodcast)
}

func (ps *podcastService) EditThumbnail(
	podcast models.PodcastRequest,
	podcastImage *multipart.FileHeader) error {

	dbPodcast, err := ps.repo.FindPodcastByEpisodeNoAndType(podcast.EpisodeNo, string(podcast.Type))
	if err != nil {
		return err
	}

	thumbnail, err := ps.repo.FetchThumbnail(dbPodcast.ThumbnailID)
	if err != nil {
		return err
	}

	err = utils.DeleteImage(thumbnail.Name)
	if err != nil {
		return err
	}

	assetChannel := make(chan int)

	go helpers.UploadAndFetchAssetID(assetChannel, podcastImage, ps.repo, "podcast")

	assetID := <-assetChannel

	if assetID == -1 {
		return errors.New("error uploading asset")
	}

	return ps.repo.UpdatePodcastThumbnail(dbPodcast.ID, uint(assetID))
}

func (ps *podcastService) EditURL(podcast models.PodcastRequest) error {
	mediaURL, err := utils.URLValidator(podcast.MediaURL)
	if err != nil {
		return err
	}

	dbPodcast, err := ps.repo.FindPodcastByEpisodeNoAndType(podcast.EpisodeNo, string(podcast.Type))
	if err != nil {
		return err
	}

	return ps.repo.UpdatePodcastURL(dbPodcast.ID, mediaURL)
}

func (ps *podcastService) EditDescription(podcast models.PodcastRequest) error {
	dbPodcast, err := ps.repo.FindPodcastByEpisodeNoAndType(podcast.EpisodeNo, string(podcast.Type))
	if err != nil {
		return err
	}

	return ps.repo.UpdatePodcastDescription(dbPodcast.ID, podcast.Description)
}

func (ps *podcastService) DeletePodcast(podcast models.PodcastRequest) error {
	dbPodcast, err := ps.repo.FindPodcastByEpisodeNoAndType(podcast.EpisodeNo, string(podcast.Type))
	if err != nil {
		return err
	}

	thumbnail, err := ps.repo.FetchThumbnail(dbPodcast.ID)
	if err != nil {
		return err
	}

	err = utils.DeleteImage(thumbnail.Name)
	if err != nil {
		return err
	}

	return ps.repo.DeletePodcast(dbPodcast.ID)
}
