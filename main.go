package main

import (
	"cta4j_back_end_go/services"
	"fmt"
)

func main() {
	trains := services.GetTrains("41320")

	fmt.Println(len(trains))

	for _, value := range trains {
		fmt.Printf("Train %+v\n", value)
	}
}
