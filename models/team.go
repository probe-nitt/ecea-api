package models

type MemberRequest struct {
	Name   string      `json:"name" form:"name"`
	RollNo string      `json:"roll_no" form:"rollnumber"`
	Team   MemberTeams `json:"team" form:"team"`
	Role   MemberRoles `json:"role" form:"role"`
}

type Members struct {
	Name     string      `json:"name"`
	RollNo   string      `json:"rollnumber"`
	Team     MemberTeams `json:"team"`
	Role     MemberRoles `json:"role"`
	ImageURL string      `json:"image_url"`
}

type MemberTeams string

const (
	Webops MemberTeams = "Webops"
	Events MemberTeams = "Events"
)

type MemberRoles string

const (
	Chairperson MemberRoles = "Chairperson"
	OCBoy       MemberRoles = "Overall Coordinator(Boys)"
	OCGirl      MemberRoles = "Overall Coordinator(Girls)"
	Treasurer   MemberRoles = "Treasurer"
	Head        MemberRoles = "Head"
	DeputyHead  MemberRoles = "Deputy Head"
	Manager     MemberRoles = "Manager"
	Coordinator MemberRoles = "Coordinator"
)
