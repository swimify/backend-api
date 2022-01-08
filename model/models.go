package model

import (
	"time"

	"gorm.io/gorm"
)

type TimestampFields struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Response struct {
	Persons []Person `json:"persons"`
}

type Organization struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Label string `json:"label"`
	TimestampFields
}

type Season struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimestampFields
	OrganizationId uint
	Organization   Organization
	Label          string    `json:"label"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
}

type Team struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimestampFields
	OrganizationId uint
	Organization   Organization
	Label          string `json:"label"`
}

type Sex string
type Stroke string
type Metric string

const (
	SEX_MALE          = "M"
	SEX_FEMALE        = "F"
	STROKE_FLY        = "FLY"
	STROKE_FREE       = "FREE"
	STROKE_BACK       = "BACK"
	STROKE_BREAST     = "BREAST"
	STROKE_IM         = "IM"
	STROKE_FREE_RELAT = "FREE_RELAY"
	STROKE_IM_RELATY  = "IM_RELAY"
	METRIC_YARDS      = "Y"
	METRIC_METERS     = "M"
)

type Person struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimestampFields
	OrganizationId uint
	Organization   Organization
	Label          string `json:"label"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DateOfBirth    string `json:"date_of_birth"`
	Sex            Sex    `json:"gender"`
	IsCoach        bool   `json:"is_coach"`
	IsSwimmer      bool   `json:"is_swimmer"`
	IsGuardian     bool   `json:"is_guardian"`
	FamilyID       uint
	Family         Family
	Address        `gorm:"embedded"`
}

type Family struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimestampFields
	OrganizationId uint
	Organization   Organization
	Persons        []Person
}

type Address struct {
	Street1    string `json:"street1"`
	Street2    string `json:"street2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}

type Pool struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimestampFields
	OrganizationId uint
	Organization   Organization
	Label          string `json:"label"`
	Address        `gorm:"embedded"`
	Meets          []Meet
}

type Meet struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimestampFields
	OrganizationId uint
	Organization   Organization
	Label          string `json:"label"`
	SeasonId       uint
	PoolId         uint
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	NumLanes       int       `json:"num_lanes"`
}

type Event struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimestampFields
	OrganizationId uint
	Organization   Organization
	Label          string `json:"label"`
	MeetId         uint
	Sex            Sex    `json:"sex"`
	Stroke         Stroke `json:"stroke"`
	Distance       int    `json:"distance"`
	Metric         Metric `json:"metric"`
}

type Entry struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TimestampFields
	OrganizationId uint
	Organization   Organization
	Label          string `json:"label"`
	EventId        uint
	Heat           int `json:"heat"`
	Lane           int `json:"lane"`
	EntryTime      int `json:"entry_time"`
	FinishTime     int `json:"finish_time"`
}
