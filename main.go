package main

import (
    "cta4j_back_end_go/services"
    "fmt"
)

func main() {
    trains := services.GetTrains("41710")

    fmt.Printf("Trains: %v\n", trains)
}
