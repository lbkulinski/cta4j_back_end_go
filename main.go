package main

import (
	"cta4j_back_end_go/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/trains/stations/:stationId", services.GetTrains)

	err := router.Run("localhost:8080")

	if err != nil {
		fmt.Println(err)

		return
	}
}
