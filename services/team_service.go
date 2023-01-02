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

type teamService struct {
	repo repositories.TeamRepository
}

type TeamService interface {
	CreateTeamMember(
		memberDetails models.MemberRequest,
		memberImage *multipart.FileHeader) error
	EditTeamMemberImage(
		memberDetails models.MemberRequest,
		memberImage *multipart.FileHeader) error
	EditTeamMemberName(
		memberDetails models.MemberRequest,
	) error
	EditTeamMemberRole(
		reqMember models.MemberRequest) error
	EditTeamMemberTeam(
		reqMember models.MemberRequest) error
	RemoveTeamMember(
		rollNumber string) error
	GetTeamMember(
		rollNumber string) (models.Members, error)
	GetAllTeamMember() ([]models.Members, error)
}

func NewTeamService(repo repositories.TeamRepository) TeamService {
	return &teamService{repo}
}

func (ts *teamService) CreateTeamMember(
	memberDetails models.MemberRequest,
	memberImage *multipart.FileHeader) error {

	teamChannel := make(chan int)
	roleChannel := make(chan int)
	assetChannel := make(chan int)

	go helpers.FetchTeamID(teamChannel, string(memberDetails.Team), ts.repo)

	go helpers.FetchRoleID(roleChannel, string(memberDetails.Role), ts.repo)

	go helpers.UploadAndFetchAssetID(assetChannel, memberImage, ts.repo)

	teamID := <-teamChannel
	roleID := <-roleChannel
	assetID := <-assetChannel

	if teamID == -1 || roleID == -1 || assetID == -1 {
		return errors.New("Error Occurred")
	}

	name, err := utils.NameValidator(memberDetails.Name)
	if err != nil {
		return err
	}
	rollNo, err := utils.NumericValidator(memberDetails.RollNo)
	if err != nil {
		return err
	}

	member := schemas.Member{
		Name:    name,
		RollNo:  rollNo,
		TeamID:  uint(teamID),
		RoleID:  uint(roleID),
		AssetID: uint(assetID),
	}

	return ts.repo.InsertMember(&member)
}

func (ts *teamService) EditTeamMemberName(
	reqMember models.MemberRequest) error {

	rollNo, err := utils.NumericValidator(reqMember.RollNo)
	if err != nil {
		return err
	}

	dbMember, err := ts.repo.FindMemberByRollNo(rollNo)
	if err != nil {
		return err
	}

	name, err := utils.NameValidator(reqMember.Name)
	if err != nil {
		return err
	}

	member := schemas.Member{
		Name: name,
	}

	member.ID = dbMember.ID

	return ts.repo.UpdateMember(&member)
}

func (ts *teamService) EditTeamMemberTeam(
	reqMember models.MemberRequest) error {

	rollNo, err := utils.NumericValidator(reqMember.RollNo)
	if err != nil {
		return err
	}

	dbMember, err := ts.repo.FindMemberByRollNo(rollNo)
	if err != nil {
		return err
	}

	teamID := helpers.UpdateAndFetchTeamID(
		dbMember.TeamID,
		dbMember.Team.Name,
		string(reqMember.Team),
		ts.repo)

	if teamID == -1 {
		return errors.New("Unable to Edit")
	}

	member := schemas.Member{
		TeamID: uint(teamID),
	}

	member.ID = dbMember.ID

	return ts.repo.UpdateMember(&member)
}

func (ts *teamService) EditTeamMemberRole(
	reqMember models.MemberRequest) error {

	rollNo, err := utils.NumericValidator(reqMember.RollNo)
	if err != nil {
		return err
	}

	dbMember, err := ts.repo.FindMemberByRollNo(rollNo)
	if err != nil {
		return err
	}

	roleID := helpers.UpdateAndFetchRoleID(
		dbMember.RoleID,
		dbMember.Role.Name,
		string(reqMember.Team),
		ts.repo)

	if roleID == -1 {
		return errors.New("Unable to Edit")
	}

	member := schemas.Member{
		RoleID: uint(roleID),
	}

	member.ID = dbMember.ID

	return ts.repo.UpdateMember(&member)
}

func (ts *teamService) EditTeamMemberImage(
	reqMember models.MemberRequest,
	memberImage *multipart.FileHeader) error {

	rollNo, err := utils.NumericValidator(reqMember.RollNo)
	if err != nil {
		return err
	}

	dbMember, err := ts.repo.FindMemberByRollNo(rollNo)
	if err != nil {
		return err
	}

	assetID := helpers.UpdateAndFetchAssetID(dbMember.AssetID, dbMember.Asset.Name, memberImage, ts.repo)

	if assetID == -1 {
		return errors.New("Error Occurred")
	}

	member := schemas.Member{
		AssetID: uint(assetID),
	}
	member.ID = dbMember.ID

	return ts.repo.UpdateMember(&member)
}

func (ts *teamService) RemoveTeamMember(rollNumber string) error {

	rollNo, err := utils.NumericValidator(rollNumber)
	if err != nil {
		return err
	}

	dbMember, err := ts.repo.FindMemberByRollNo(rollNo)
	if err != nil {
		return err
	}
	err = utils.DeleteImage(dbMember.Asset.Name)
	if err != nil {
		return err
	}
	return ts.repo.DeleteMember(rollNumber)
}

func (ts *teamService) GetTeamMember(rollNumber string) (models.Members, error) {
	return ts.repo.FindMember(rollNumber)

}

func (ts *teamService) GetAllTeamMember() ([]models.Members, error) {
	return ts.repo.FindAllMembers()
}
