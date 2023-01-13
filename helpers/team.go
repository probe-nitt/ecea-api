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

func UpdateAndFetchTeamID(
	dbID uint,
	dbName string,
	reqName string,
	repo repositories.TeamRepository) int {

	if dbName == reqName {
		return int(dbID)
	}

	id, err := repo.FindTeamByName(reqName)
	if err != nil {
		log.Println(err)

		return -1
	}
	return int(id)
}

func UpdateAndFetchRoleID(
	dbID uint,
	dbName string,
	reqName string,
	repo repositories.TeamRepository) int {

	if dbName == reqName {
		return int(dbID)
	}
	id, err := repo.FindRoleByName(reqName)
	if err != nil {
		log.Println(err)
		return -1
	}
	return int(id)
}

func UpdateAndFetchAssetID(
	dbID uint,
	dbName string,
	image *multipart.FileHeader,
	repo repositories.TeamRepository) int {

	if dbName == image.Filename {
		return int(dbID)
	}

	err := utils.DeleteImage(dbName)

	if err != nil {
		log.Println(err)
		return -1
	}

	fileName, err := utils.UploadImage(image)
	if err != nil {
		log.Println(err)
		return -1
	}

	err = repo.UpdateAsset(dbID, fileName)

	if err != nil {
		log.Println(err)
		return -1
	}

	return int(dbID)
}
