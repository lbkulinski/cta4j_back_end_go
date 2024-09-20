package models

import (
	"encoding/json"
	"time"
)

type Train struct {
	Run            string    `json:"run"`
	Line           Line      `json:"line"`
	Destination    string    `json:"destination"`
	Station        string    `json:"station"`
	PredictionTime time.Time `json:"predictionTime"`
	ArrivalTime    time.Time `json:"arrivalTime"`
	Due            bool      `json:"due"`
	Scheduled      bool      `json:"scheduled"`
	Delayed        bool      `json:"delayed"`
}

func (receiver *Train) UnmarshalJSON(data []byte) error {
	type apiTrain struct {
		Run            string `json:"rn"`
		Line           Line   `json:"rt"`
		Destination    string `json:"destNm"`
		Station        string `json:"staNm"`
		PredictionTime string `json:"prdt"`
		ArrivalTime    string `json:"arrT"`
		Due            string `json:"isApp"`
		Scheduled      string `json:"isSch"`
		Delayed        string `json:"isDly"`
	}

	var train apiTrain

	jsonError := json.Unmarshal(data, &train)

	if jsonError != nil {
		return jsonError
	}

	receiver.Run = train.Run

	receiver.Line = train.Line

	receiver.Destination = train.Destination

	receiver.Station = train.Station

	customFormat := "2006-01-02T03:04:05"

	location, locationError := time.LoadLocation("America/Chicago")

	if locationError != nil {
		return locationError
	}

	var predictionError error

	receiver.PredictionTime, predictionError = time.ParseInLocation(customFormat, train.PredictionTime, location)

	if predictionError != nil {
		return predictionError
	}

	var arrivalError error

	receiver.ArrivalTime, arrivalError = time.ParseInLocation(customFormat, train.ArrivalTime, location)

	if arrivalError != nil {
		return arrivalError
	}

	receiver.Due = train.Due == "1"

	receiver.Scheduled = train.Scheduled == "1"

	receiver.Delayed = train.Delayed == "1"

	return nil
}
