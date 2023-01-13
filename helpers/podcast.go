package helpers

import (
	"log"

	"github.com/ecea-nitt/ecea-server/repositories"
)

func FetchPodcastTypeID(
	channel chan int,
	podcastTypeName string,
	repo repositories.PodcastRepository) {

	id, err := repo.FindPodcastTypeByName(podcastTypeName)
	if err != nil {
		log.Println(err)
		channel <- -1
		return
	}
	channel <- int(id)

}
