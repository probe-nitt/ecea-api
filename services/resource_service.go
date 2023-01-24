package services

import (
	"errors"
	"log"
	"mime/multipart"

	"github.com/ecea-nitt/ecea-server/helpers"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/schemas"
)

type studyMaterialService struct {
	repo repositories.StudyMaterialRepository
}

type StudyMaterialService interface {
	CreateNewStudyMaterial(studyMaterialDetails models.StudyMaterialRequest, studyMaterialFile *multipart.FileHeader) error
	//	EditStudyMaterialSubject(subject string, code string) error
	EditStudyMaterialURL(name string, file *multipart.FileHeader) error
	GetStudyMaterial(name string) (models.StudyMaterials, error)
	GetCategoryStudyMaterials(name string) ([]models.StudyMaterials, error)
	RemoveStudyMaterial(name string) error
	GetAllMaterials() ([]models.CategoryMaterials, error)
	AddSubject(studyMaterialDetails *models.StudyMaterialRequest) error
}

func NewStudyMaterialService(repo repositories.StudyMaterialRepository) StudyMaterialService {
	return &studyMaterialService{repo}
}

func (rs *studyMaterialService) CreateNewStudyMaterial(studyMaterialDetails models.StudyMaterialRequest, studyMaterialFile *multipart.FileHeader) error {
	assetChannel := make(chan int)
	subjectChannel := make(chan int)
	subjectCategoryChannel := make(chan int)
	/*
		subjectCategoryID, err := rs.repo.GetSubjectCategoryID(studyMaterialDetails.SubjectCategory)
		if err != nil {
			return err
		}
	*/
	go helpers.FetchSubjectCategoryID(subjectCategoryChannel, studyMaterialDetails.SubjectCode, rs.repo)
	subjectCategoryID := <-subjectCategoryChannel
	go helpers.FetchSubjectID(subjectChannel, studyMaterialDetails.SubjectCode, rs.repo)
	go helpers.UploadFileAndFetchAssetID(assetChannel, studyMaterialFile, rs.repo)
	subjectID := <-subjectChannel
	assetID := <-assetChannel
	if subjectID == -1 || assetID == -1 || subjectCategoryID == -1 {
		return errors.New("could not fetch data")
	}
	studyMaterial := schemas.StudyMaterial{
		Name:      studyMaterialDetails.Name,
		SubjectID: uint(subjectID),
		AssetID:   uint(assetID),
	}
	log.Println(subjectCategoryID, studyMaterialDetails.SubjectCategory)
	return rs.repo.CreateStudyMaterial(&studyMaterial)
}

/*
func (rs *studyMaterialService) EditStudyMaterialSubject(subject string, code string) error {
	editrequest := schemas.Subject{
		Name:        subject,
		SubjectCode: code,
	}
		subjectID, err := rs.repo.GetSubjectIDByCode(code)
		if err != nil {
			return err
		}
	return rs.repo.UpdateStudyMaterialSubject(&editrequest, subject, code)
}
*/

func (rs *studyMaterialService) EditStudyMaterialURL(name string, file *multipart.FileHeader) error {
	dbStudyMaterial, err := rs.repo.FindStudyMaterialByNameReturnSchema(name)
	if err != nil {
		return err
	}

	assetID := helpers.UpdateFileAndFetchAssetID(dbStudyMaterial.AssetID, dbStudyMaterial.Asset.Name, file, rs.repo)

	if assetID == -1 {
		return errors.New(" Error Occurred")
	}

	studyMaterial := schemas.StudyMaterial{
		AssetID: uint(assetID),
	}
	studyMaterial.ID = dbStudyMaterial.ID

	return rs.repo.UpdateStudyMaterial(&studyMaterial)
}

func (rs *studyMaterialService) RemoveStudyMaterial(name string) error {
	subjectID, err1 := rs.repo.GetSubjectIDByName(name)
	if err1 != nil {
		return err1
	}
	assetID, err2 := rs.repo.GetAssetIDByName(name)
	if err2 != nil {
		return err2
	}
	return rs.repo.DeleteStudyMaterial(name, subjectID, assetID)
}

func (rs *studyMaterialService) GetStudyMaterial(name string) (models.StudyMaterials, error) {
	studyMaterial, err := rs.repo.FindStudyMaterialByName(name)
	if err != nil {
		return models.StudyMaterials{}, err
	}
	subjectCategory, err1 := rs.repo.GetSubjectCategory(studyMaterial.Subject)
	if err1 != nil {
		return models.StudyMaterials{}, err1
	}
	result := models.StudyMaterials{
		Name:            studyMaterial.Name,
		Subject:         studyMaterial.Subject,
		SubjectCategory: subjectCategory,
		SubjectCode:     studyMaterial.SubjectCode,
		DocumentURL:     studyMaterial.DocumentURL,
	}
	log.Println(studyMaterial)
	log.Println(result)
	return result, err
}

func (rs *studyMaterialService) GetCategoryStudyMaterials(name string) ([]models.StudyMaterials, error) {
	return rs.repo.GetCategoryStudyMaterials(name)
}

func (rs *studyMaterialService) AddSubject(studyMaterialDetails *models.StudyMaterialRequest) error {
	subjectName := studyMaterialDetails.Subject
	subjectCode := studyMaterialDetails.SubjectCode
	subjectCategory := studyMaterialDetails.SubjectCategory
	id, err := rs.repo.GetSubjectCategoryID(subjectCategory)
	if err != nil {
		return err
	}
	return rs.repo.CreateSubject(subjectName, subjectCode, id)
}

func (rs *studyMaterialService) GetAllMaterials() ([]models.CategoryMaterials, error) {
	allCategories, err := rs.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	var allMaterials []models.CategoryMaterials
	var categoryMaterials models.CategoryMaterials
	for _, category := range allCategories {
		categoryMaterials.Category = category
		temp, err := rs.repo.GetCategoryStudyMaterials(category)
		if err != nil {
			return nil, err
		}
		categoryMaterials.Material = temp
		allMaterials = append(allMaterials, categoryMaterials)
	}
	return allMaterials, nil
}
