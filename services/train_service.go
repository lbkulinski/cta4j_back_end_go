package services

import (
	"cta4j_back_end_go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
)

func GetTrains(context *gin.Context) {
	client := resty.New()

	apiKey := os.Getenv("TRAIN_API_KEY")

	stationId := context.Param("stationId")

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
		context.Status(http.StatusInternalServerError)

		return
	}

	if trainResponse.Trains == nil {
		context.Status(http.StatusNotFound)

		return
	}

	context.JSON(http.StatusOK, trainResponse.TrainBody.Trains)
}
