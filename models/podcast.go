package models

type PodcastRequest struct {
	Name        string      `json:"name" form:"name"`
	Description string      `json:"description" form:"description"`
	MediaURL    string      `json:"mediaUrl" form:"mediaUrl"`
	Type        PodcastType `json:"type" form:"type"`
}

type PodcastType string

const (
	Guestlecture PodcastType = "GUEST_LECTURE"
	CareerPath   PodcastType = "CAREER_PATH"
)
