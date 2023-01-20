package repositories

import (
	"github.com/ecea-nitt/ecea-server/config"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/schemas"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type studyMaterialRepository struct {
	db *gorm.DB
}

type StudyMaterialRepository interface {
	GetCategoryStudyMaterials(name string) ([]models.StudyMaterials, error)
	FindStudyMaterialByName(name string) (models.StudyMaterials, error)
	FindStudyMaterialByNameReturnSchema(name string) (schemas.StudyMaterial, error)
	GetSubjectID(name string) (uint, error)
	GetSubjectCategoryID(name string) (uint, error)
	GetSubjectCategory(subject string) (string, error)
	GetSubjectIDByName(name string) (uint, error)
	GetAssetIDByName(name string) (uint, error)
	InsertAsset(name string) (uint, error)
	UpdateAsset(id uint, name string) error
	CreateSubjectAndReturnID(studyMaterialDetails models.StudyMaterialRequest, id uint) (uint, error)
	CreateStudyMaterial(studyMaterial *schemas.StudyMaterial) error
	UpdateStudyMaterialSubject(editrequest *schemas.Subject, name string, code string) error
	UpdateStudyMaterial(studyMaterial *schemas.StudyMaterial) error
	DeleteStudyMaterial(name string, subjectID uint, assetID uint) error
}

func NewStudyMaterialRepository(db *gorm.DB) StudyMaterialRepository {
	return &studyMaterialRepository{db}
}

func (rr *studyMaterialRepository) UpdateStudyMaterialSubject(editrequest *schemas.Subject, name string, code string) error {
	/*
		return rr.db.Table("subjects").Where("subjects.subject_code = ?", code).Update("subjects.name", name).Error
			return rr.db.Model(&schemas.Subject{}).Where("subjects.subject_code = ?", code).Update("name", name).Error
	*/
	return rr.db.Model(&editrequest).Where("subject_code = ?", code).Updates(&editrequest).Error
}

func (rr *studyMaterialRepository) CreateStudyMaterial(studyMaterial *schemas.StudyMaterial) error {
	return rr.db.Create(&studyMaterial).Error
}

func (rr *studyMaterialRepository) GetSubjectIDByName(name string) (uint, error) {
	var studyMaterial schemas.StudyMaterial
	res := rr.db.Where("name = ?", name).First(&studyMaterial)
	if res.Error != nil {
		return 0, res.Error
	}
	return studyMaterial.SubjectID, nil
}
func (rr *studyMaterialRepository) DeleteStudyMaterial(name string, subjectID uint, assetID uint) error {
	err := rr.db.Unscoped().Where("id = ?", subjectID).Delete(&schemas.Subject{}).Error
	if err != nil {
		return err
	}
	err = rr.db.Unscoped().Where("id = ?", assetID).Delete(&schemas.Asset{}).Error
	if err != nil {
		return err
	}
	return rr.db.Unscoped().Where("name = ?", name).Delete(&schemas.StudyMaterial{}).Error
}

func (rr *studyMaterialRepository) GetSubjectID(name string) (uint, error) {
	var subject schemas.Subject
	res := rr.db.Where("subject_code = ?", name).First(&subject)
	if res.Error != nil {
		return 0, res.Error
	}
	return subject.ID, nil
}

func (rr *studyMaterialRepository) GetSubjectCategoryID(name string) (uint, error) {
	var subjectCategory schemas.SubjectCategory
	res := rr.db.Where("name = ?", name).First(&subjectCategory)
	if res.Error != nil {
		return 0, res.Error
	}
	return subjectCategory.ID, nil
}

func (rr *studyMaterialRepository) GetAssetIDByName(name string) (uint, error) {
	var studyMaterial schemas.StudyMaterial
	res := rr.db.Where("name = ?", name).First(&studyMaterial)
	if res.Error != nil {
		return 0, res.Error
	}
	return studyMaterial.AssetID, nil
}

func (rr *studyMaterialRepository) InsertAsset(name string) (uint, error) {
	var assetType schemas.AssetType
	if err := rr.db.Where("name = ?", schemas.Document).First(&assetType).Error; err != nil {
		return 0, err
	}
	asset := schemas.Asset{
		Name:        name,
		AssetTypeID: assetType.ID,
	}
	if err := rr.db.Create(&asset).Error; err != nil {
		return 0, err
	}
	return asset.ID, nil
}

func (rr *studyMaterialRepository) FindStudyMaterialByName(name string) (models.StudyMaterials, error) {
	/*
		var studyMaterial schemas.StudyMaterial

		err := rr.db.Preload(clause.Associations).Where("name = ?", name).First(&studyMaterial).Error
		return studyMaterial, err
	*/
	var studyMaterials models.StudyMaterials
	err := rr.db.Table("study_materials").Select(
		`study_materials.name as name,subjects.name as subject,subjects.subject_code as subject_code,CONCAT(?::text,'/static/documents','/',assets.name) as document_url`, config.Origin,
	).Joins(
		"JOIN assets on assets.id = study_materials.asset_id").Joins(
		"JOIN subjects on subjects.id = study_materials.subject_id").Where(
		"study_materials.name = ?", name).First(
		&studyMaterials).Error
	if err != nil {
		return models.StudyMaterials{}, err
	}
	return studyMaterials, nil
}

func (rr *studyMaterialRepository) UpdateStudyMaterial(studyMaterial *schemas.StudyMaterial) error {
	return rr.db.Model(&studyMaterial).Where("id = ?", studyMaterial.ID).Updates(&studyMaterial).Error
}

func (rr *studyMaterialRepository) GetCategoryStudyMaterials(name string) ([]models.StudyMaterials, error) {
	var studyMaterials []models.StudyMaterials
	if err := rr.db.Table("study_materials").Select(
		`study_materials.name as name,subjects.name as subject,subject_categories.name as subject_category,subjects.subject_code as subject_code, CONCAT(?::text,'/static/documents','/',assets.name) as document_url`, config.Origin,
	).Joins(
		"JOIN assets on assets.id = study_materials.asset_id").Joins(
		"JOIN subjects on subjects.id = study_materials.subject_id").Joins("JOIN subject_categories on subject_categories.id = subjects.subject_category_id").Where("subject_categories.name = ?", name).Scan(
		&studyMaterials).Error; err != nil {
		return nil, err
	}

	return studyMaterials, nil
}

func (rr *studyMaterialRepository) UpdateAsset(id uint, name string) error {
	return rr.db.Model(&schemas.Asset{}).Where("id = ?", id).Update("name", name).Error
}

func (rr *studyMaterialRepository) CreateSubjectAndReturnID(studyMaterialDetails models.StudyMaterialRequest, id uint) (uint, error) {
	subject := schemas.Subject{
		Name:              studyMaterialDetails.Subject,
		SubjectCategoryID: id,
		SubjectCode:       studyMaterialDetails.SubjectCode,
	}
	res := rr.db.Create(&subject)
	if res.Error != nil {
		return 0, res.Error
	}
	return subject.ID, nil
}

func (rr *studyMaterialRepository) GetSubjectCategory(subject string) (string, error) {
	var category string
	if err := rr.db.Table("subject_categories").Select(
		`subject_categories.name`,
	).Joins(
		"JOIN subjects on subjects.subject_category_id = subject_categories.id").Where("subjects.name = ?", subject).Scan(
		&category).Error; err != nil {
		return "", err
	}
	return category, nil
}

func (rr *studyMaterialRepository) FindStudyMaterialByNameReturnSchema(name string) (schemas.StudyMaterial, error) {
	var studyMaterial schemas.StudyMaterial

	err := rr.db.Preload(clause.Associations).Where("name = ?", name).First(&studyMaterial).Error
	return studyMaterial, err
}
