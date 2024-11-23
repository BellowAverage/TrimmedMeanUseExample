package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	trimmedmean "github.com/BellowAverage/TrimmedMeanPackage"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a sample of 100 integers
	intData := make([]int, 100)
	for i := range intData {
		intData[i] = rand.Intn(1000)
	}

	// Compute trimmed mean for integers
	meanInt, err := trimmedmean.TrimmedMeanInt(intData, 0.05)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trimmed mean of integers: %f\n", meanInt)

	// Generate a sample of 100 floats
	floatData := make([]float64, 100)
	for i := range floatData {
		floatData[i] = rand.Float64() * 1000
	}

	// Compute trimmed mean for floats
	meanFloat, err := trimmedmean.TrimmedMeanFloat(floatData, 0.05)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trimmed mean of floats: %f\n", meanFloat)
}
