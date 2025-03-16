package station

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Service interface {
	GetAllStations() ([]StationResponse, error)
	CheckSchedules(id string) (response []ScheduleResponse, err error)
}

type service struct {
	client *http.Client
}

// NewService creates a new instance of station service
func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func DoRequest(client *http.Client, url string)([]byte, error){
	resp, err := client.Get(url)
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected Status Code: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	return body, nil
}

func (s *service) GetAllStations() ([]StationResponse, error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return nil, err
	}

	// Initialize slice before appending
	response := make([]StationResponse, 0, len(stations))

	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return response, nil
}

func (s *service) CheckSchedules(id string)(response []ScheduleResponse, err error){
	url := "https://jakartamrt.co.id/id/val/stasiuns"
	byteResponse, err := DoRequest(s.client, url)
	if err != nil{
		return nil, err
	}
	var schedules []Schedule
	err = json.Unmarshal(byteResponse, &schedules)
	if err != nil {
		return 
	}
	var scheduleSelected *Schedule
	for _, item:=range schedules{
		if item.StationId == id{
			scheduleSelected = &item
			break
		}
	}

	if scheduleSelected == nil{
		return nil, errors.New("Station Not Found")
	}

	response, err = ConvertDataToResponses(*scheduleSelected)
	if err != nil{
		return
	}

	return 
}


func ConvertDataToResponses(schedule Schedule)(response []ScheduleResponse, err error){
	var (
		LebakBulusTripName = "Stasiun Lebak Bulus Grab"
		BundaranHITripName = "Stasiun Bundaran HI Bank DKI"
	)

	scheduleLebakBulus := schedule.ScheduleLebakBulus
	scheduleBundaranHI := schedule.ScheduleBundaranHI

	scheduleLebakBulusParsed, err := ConvertScheduleToTimeFormat(scheduleLebakBulus)
	if err != nil{
		return
	}
	scheduleBundaranHIParsed, err := ConvertScheduleToTimeFormat(scheduleBundaranHI)
	if err != nil{
		return
	}
	for _, item := range scheduleLebakBulusParsed{
		if item.Format("15:04") > time.Now().Format("15:04"){
			response = append(response, ScheduleResponse{
				StationName: LebakBulusTripName,
				Time: item.Format("15:04"),
			})
		}
	}
	for _, item := range scheduleBundaranHIParsed{
		if item.Format("15:04") > time.Now().Format("15:04"){
			response = append(response, ScheduleResponse{
				StationName: BundaranHITripName,
				Time: item.Format("15:04"),
			})
		}
	}
	return
}

func ConvertScheduleToTimeFormat(schedule string) ([]time.Time, error) {
	var response []time.Time

	schedules := strings.Split(schedule, ",")
	for _, item := range schedules {
		trimmedTime := strings.TrimSpace(item)
		if trimmedTime == "" {
			continue
		}

		parsedTime, err := time.Parse("15:04", trimmedTime)
		if err != nil {
			return nil, errors.New("invalid time format: " + trimmedTime)
		}

		response = append(response, parsedTime)
	}

	return response, nil
}