package helpers

import (
	"log"
	"mime/multipart"

	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/utils"
)

func UploadAndFetchAssetID(
	channel chan int,
	image *multipart.FileHeader,
	repo interface{},
	repoType string) {

	switch repoType {
	case "podcast":
		repo := repo.(repositories.PodcastRepository)
		fileName, err := utils.UploadImage(image)
		if err != nil {
			log.Println(err)
			channel <- -1
			return
		}

		id, err := repo.InsertAsset(fileName)
		if err != nil {
			log.Println(err)
			channel <- -1
			return
		}

		channel <- int(id)
	case "team":
		repo := repo.(repositories.TeamRepository)
		fileName, err := utils.UploadImage(image)
		if err != nil {
			log.Println(err)
			channel <- -1
			return
		}

		id, err := repo.InsertAsset(fileName)
		if err != nil {
			log.Println(err)
			channel <- -1
			return
		}

		channel <- int(id)
	}
}
