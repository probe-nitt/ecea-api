package helpers

import (
	"log"
	"mime/multipart"

	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/utils"
)

func FetchTeamID(
	channel chan int,
	teamName string,
	repo repositories.TeamRepository) {

	id, err := repo.FindTeamByName(teamName)
	if err != nil {
		log.Println(err)
		channel <- -1
		return
	}
	channel <- int(id)

}

func FetchRoleID(
	channel chan int,
	roleName string,
	repo repositories.TeamRepository) {

	id, err := repo.FindRoleByName(roleName)
	if err != nil {
		log.Println(err)
		channel <- -1
		return
	}
	channel <- int(id)

}

func UploadAndFetchAssetID(
	channel chan int,
	image *multipart.FileHeader,
	repo repositories.TeamRepository) {

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
