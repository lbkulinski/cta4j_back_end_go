package services

import (
	"cta4j_back_end_go/models"
	"github.com/go-resty/resty/v2"
	"os"
)

func GetTrains(stationId string) []models.Train {
	client := resty.New()

	apiKey := os.Getenv("TRAIN_API_KEY")

	var trainResponse models.TrainResponse

	response, err := client.R().
		SetQueryParams(map[string]string{
			"key":        apiKey,
			"outputType": "json",
			"mapid":      stationId,
		}).
		SetResult(&trainResponse).
		Get("https://lapi.transitchicago.com/api/1.0/ttarrivals.aspx")

	if err != nil || response.IsError() {
		return nil
	}

	return trainResponse.Trains
}
