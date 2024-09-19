package models

import (
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
