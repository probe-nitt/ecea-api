package models

// HTTPError example
type Error struct {
	Code    int    `json:"code" example:"000"`
	Message string `json:"message" example:"status"`
}
