package schemas

import (
	"gorm.io/gorm"
)

type Podcast struct {
	gorm.Model
	Name        string `gorm:"varchar(50);not null;unique"`
	Description string `gorm:"varchar(800);not null"`
	EpisodeNo   uint   `gorm:"not null;uniqueIndex:epno_type"`
	MediaURL    string `gorm:"not null"`

	// Relations
	TypeID      uint        `gorm:"not null;uniqueIndex:epno_type"`
	Type        PodcastType `gorm:"foreignKey:TypeID;"`
	ThumbnailID uint        `gorm:"not null;"`
	Thumbnail   Asset       `gorm:"foreignKey:ThumbnailID;"`
}

type PodcastType struct {
	ID uint `gorm:"primarykey"`
	//gorm.Model
	Name string `gorm:"not null;unique"`
}

type PodcastTypes string

const (
	GuestLecture PodcastTypes = "GUEST_LECTURE"
	CareerPath   PodcastTypes = "CAREER_PATH"
)
