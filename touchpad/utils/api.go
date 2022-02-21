package utils

import (
	"encoding/json"
	"fmt"
	touchpad "github.com/ansg191/northstars-backend/touchpad/proto"
	"github.com/micro/micro/v3/service/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strconv"
	"time"
)

const TouchpadDateFmt = "2006-01-02"

type GetMeetInfoResponse struct {
	CourseName         string `json:"courseName"`
	CourseOrder        string `json:"courseOrder"`
	EndDate            string `json:"endDate"`
	Events             int    `json:"events"`
	Females            int    `json:"females"`
	HostTeamName       string `json:"hostTeamName"`
	Id                 int    `json:"id"`
	IsEnded            bool   `json:"isEnded"`
	IsEnded24Hrs       bool   `json:"isEnded24Hrs"`
	IsFinalsMeet       bool   `json:"isFinalsMeet"`
	IsPrelimMeet       bool   `json:"isPrelimMeet"`
	IsStarted          bool   `json:"isStarted"`
	LocalizedEndDate   string `json:"localizedEndDate"`
	LocalizedStartDate string `json:"localizedStartDate"`
	Males              int    `json:"males"`
	MeetName           string `json:"meetName"`
	StartDate          string `json:"startDate"`
	TeamCount          int    `json:"teamCount"`
}

func (r *GetMeetInfoResponse) ToProto() (touchpad.GetMeetInfoResponse, error) {
	startDate, err := time.Parse(TouchpadDateFmt, r.StartDate)
	if err != nil {
		return touchpad.GetMeetInfoResponse{}, nil
	}
	endDate, err := time.Parse(TouchpadDateFmt, r.EndDate)
	if err != nil {
		return touchpad.GetMeetInfoResponse{}, nil
	}

	return touchpad.GetMeetInfoResponse{
		Id:           int32(r.Id),
		CourseName:   r.CourseName,
		CourseOrder:  r.CourseOrder,
		StartDate:    timestamppb.New(startDate),
		EndDate:      timestamppb.New(endDate),
		Events:       int32(r.Events),
		Females:      int32(r.Females),
		Males:        int32(r.Males),
		HostTeamName: r.HostTeamName,
		IsEnded:      r.IsEnded,
		IsEnded24Hrs: r.IsEnded24Hrs,
		IsFinalsMeet: r.IsFinalsMeet,
		IsPrelimMeet: r.IsPrelimMeet,
		IsStarted:    r.IsStarted,
		MeetName:     r.MeetName,
		TeamCount:    int32(r.TeamCount),
	}, nil
}

type GetMeetEventsResponse struct {
	AgeHi       int    `json:"age_hi"`
	AgeLow      int    `json:"age_low"`
	AgeGroup    string `json:"ageGroup"`
	Day         string `json:"day"`
	Distance    int    `json:"distance"`
	EventNumber string `json:"eventNumber"`
	Gender      string `json:"gender"`
	Id          int    `json:"id"`
	Relay       bool   `json:"relay"`
	Rounds      int    `json:"rounds"`
	Session     int    `json:"session"`
	Sponsor     string `json:"sponsor"`
	Status      int    `json:"status"`
	Stroke      string `json:"stroke"`
}

func (r *GetMeetEventsResponse) ToProto() (touchpad.Event, error) {
	eventNumber, err := strconv.ParseInt(r.EventNumber, 10, 32)
	if err != nil {
		return touchpad.Event{}, err
	}

	var gender touchpad.Event_Gender
	switch r.Gender {
	case "Mixed":
		gender = touchpad.Event_MIXED
	case "Male":
		gender = touchpad.Event_MALE
	case "Female":
		gender = touchpad.Event_FEMALE
	}

	var stroke touchpad.Event_Stroke
	switch r.Stroke {
	case "Free":
		stroke = touchpad.Event_FREE
	case "Breast":
		stroke = touchpad.Event_BREAST
	case "Back":
		stroke = touchpad.Event_BACK
	case "Fly":
		stroke = touchpad.Event_FLY
	case "Medley":
		stroke = touchpad.Event_MEDLEY
	}

	var status touchpad.Event_Status
	switch r.Status {
	case 10:
		status = touchpad.Event_UPCOMING
	case 50:
		status = touchpad.Event_SEEDED
	case 60:
		status = touchpad.Event_IN_PROGRESS
	case 80:
		status = touchpad.Event_COMPLETED
	}

	return touchpad.Event{
		Id:          int32(r.Id),
		AgeHi:       int32(r.AgeHi),
		AgeLow:      int32(r.AgeLow),
		AgeGroup:    r.AgeGroup,
		Day:         r.Day,
		Distance:    int32(r.Distance),
		EventNumber: int32(eventNumber),
		Gender:      gender,
		Relay:       r.Relay,
		Rounds:      int32(r.Rounds),
		Session:     int32(r.Session),
		Status:      status,
		Stroke:      stroke,
	}, nil
}

type IndividualEvent struct {
	Age                 int           `json:"age"`
	EventID             int           `json:"eventID"`
	FinalFormattedTime  string        `json:"finalFormattedTime"`
	FinalsHeat          int           `json:"finalsHeat"`
	FinalsLane          int           `json:"finalsLane"`
	FinalsPoints        float64       `json:"finalsPoints"`
	FinalsRank          int           `json:"finalsRank"`
	FinalsSplits        []interface{} `json:"finalsSplits"`
	FinalsSplitsLap     []interface{} `json:"finalsSplitsLap"`
	FinalsTime          int           `json:"finalsTime"`
	FirstName           string        `json:"firstName"`
	FormattedSeedTime   string        `json:"formattedSeedTime"`
	Heat                int           `json:"heat"`
	Lane                int           `json:"lane"`
	LastName            string        `json:"lastName"`
	PrelimFormattedTime string        `json:"prelimFormattedTime"`
	PrelimPoints        float64       `json:"prelimPoints"`
	PrelimRank          int           `json:"prelimRank"`
	PrelimSplits        []interface{} `json:"prelimSplits"`
	PrelimSplitsLap     []interface{} `json:"prelimSplitsLap"`
	PrelimTime          int           `json:"prelimTime"`
	SeedCourse          string        `json:"seedCourse"`
	SeedTime            int           `json:"seedTime"`
	SeedTime2Compare    int           `json:"seedTime2Compare"`
	SwimmerID           int           `json:"swimmerID"`
	TeamID              int           `json:"teamID"`
	TeamName            string        `json:"teamName"`
}

func (e *IndividualEvent) ToProto() (touchpad.ITiming, error) {
	return touchpad.ITiming{
		EventID:           int32(e.EventID),
		SwimmerID:         int32(e.SwimmerID),
		TeamID:            int32(e.TeamID),
		FirstName:         e.FirstName,
		LastName:          e.LastName,
		TeamName:          e.TeamName,
		Heat:              int32(e.FinalsHeat),
		Lane:              int32(e.FinalsLane),
		SeedTime:          int32(e.SeedTime),
		FormattedSeedTime: e.FormattedSeedTime,
		Points:            float32(e.FinalsPoints),
		Rank:              int32(e.FinalsRank),
		Time:              int32(e.FinalsTime),
		FormattedTime:     e.FinalFormattedTime,
	}, nil
}

func GetMeetInfo(meetID int) (*GetMeetInfoResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.touchpadlive.com/rest/touchpadlive/meets/%d", meetID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.NotFound("touchpad.GetMeetInfo.MeetNotFound", "Meet not found for id %d", meetID)
	}

	var info GetMeetInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func GetMeetEvents(meetID int) ([]GetMeetEventsResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.touchpadlive.com/rest/touchpadlive/meets/%d/events", meetID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.NotFound("touchpad.GetMeetEvents.MeetNotFound", "Meet not found for id %d", meetID)
	}

	var info []GetMeetEventsResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func GetIndividualEvent(meetID int, eventID int) ([]IndividualEvent, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.touchpadlive.com/rest/touchpadlive/meets/%d/individual/%d", meetID, eventID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.NotFound("touchpad.GetIEvent.MeetNotFound", "Meet not found for id %d", meetID)
	}

	var info []IndividualEvent
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	return info, nil
}
