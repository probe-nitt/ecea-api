package helpers

import (
	"log"
	"mime/multipart"

	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/utils"
)

func FetchSubjectID(channel chan int, code string, repo repositories.StudyMaterialRepository) {
	subID, err := repo.GetSubjectID(code)
	if err != nil {
		return
	}
	channel <- int(subID)

}

func FetchSubjectCategoryID(channel chan int, subjectCode string, repo repositories.StudyMaterialRepository) {
	id, err := repo.GetSubjectCategoryIDByCode(subjectCode)
	if err != nil {
		log.Println("hehehaha")
		channel <- -1
		return
	}
	channel <- int(id)

}

func UploadFileAndFetchAssetID(channel chan int, file *multipart.FileHeader, repo repositories.StudyMaterialRepository) {

	fileName, err := utils.UploadDocument(file)
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

func UpdateFileAndFetchAssetID(dbID uint, dbName string, document *multipart.FileHeader, repo repositories.StudyMaterialRepository) int {

	if dbName == document.Filename {
		return int(dbID)
	}

	err := utils.DeleteDocument(dbName)

	if err != nil {
		log.Println(err)
		return -1
	}

	fileName, err := utils.UploadDocument(document)
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
