package models

type MemberRequest struct {
	Name   string `json:"name" form:"name"`
	RollNo string `json:"roll_no" form:"roll_no"`
	Team   string `json:"team" form:"team"`
	Role   string `json:"role" form:"role"`
}
