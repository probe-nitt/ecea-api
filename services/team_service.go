package services

import (
	"errors"
	"mime/multipart"

	"github.com/ecea-nitt/ecea-server/helpers"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/schemas"
)

type teamService struct {
	repo repositories.TeamRepository
}

type TeamService interface {
	CreateTeamMember(
		memberDetails models.MemberRequest,
		memberImage *multipart.FileHeader) error
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

	go helpers.FetchTeamID(teamChannel, memberDetails.Team, ts.repo)

	go helpers.FetchRoleID(roleChannel, memberDetails.Role, ts.repo)

	go helpers.UploadAndFetchAssetID(assetChannel, memberImage, ts.repo)

	teamID := <-teamChannel
	roleID := <-roleChannel
	assetID := <-assetChannel

	if teamID == -1 || roleID == -1 || assetID == -1 {
		return errors.New("Error Occurred")
	}

	member := schemas.Member{
		Name:    memberDetails.Name,
		RollNo:  memberDetails.RollNo,
		TeamID:  uint(teamID),
		RoleID:  uint(roleID),
		AssetID: uint(assetID),
	}

	return ts.repo.InsertMember(&member)

}
