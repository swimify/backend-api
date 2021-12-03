package model

type Response struct {
	Persons []Person `json:"persons"`
}

type Organization struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Season struct {
	Id        int    `json:"id"`
	OrgId     int    `json:"org_id"`
	Label     string `json:"label"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Team struct {
	Id    int `json:"id"`
	OrgId int `json:"org_id"`
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
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Sex         Sex    `json:"gender"`
	IsCoach     bool   `json:"is_coach"`
	IsSwimmer   bool   `json:"is_swimmer"`
	IsParent    bool   `json:"is_swimmer"`
}

type Address struct {
	Street1    string `json:"street1"`
	Street2    string `json:"street2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}

type Pool struct {
	Id      int     `json:"id"`
	OrgId   int     `json:"org_id"`
	Label   string  `json:"label"`
	Address Address `json:"address"`
}

type Meet struct {
	Id        int    `json:"id"`
	OrgId     int    `json:"org_id"`
	SeasonId  int    `json:"season_id"`
	PoolId    int    `json:"pool_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	NumLanes  int    `json:"num_lanes"`
}

type Event struct {
	Id       int    `json:"id"`
	OrgId    int    `json:"org_id"`
	MeetId   int    `json:"meet_id"`
	Sex      Sex    `json:"sex"`
	Stroke   Stroke `json:"stroke"`
	Distance int    `json:"distance"`
	Metric   Metric `json:"metric"`
}

type Entry struct {
	Id         int `json:"id"`
	OrgId      int `json:"org_id"`
	EventId    int `json:"meet_id"`
	Heat       int `json:"heat"`
	Lane       int `json:"lane"`
	EntryTime  int `json:"entry_time"`
	FinishTime int `json:"finish_time"`
}
