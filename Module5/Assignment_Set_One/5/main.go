package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CityClimate struct {
	Name        string
	Temperature float64
	Rainfall    float64
}

var climateData = []CityClimate{
	{Name: "A", Temperature: 30.5, Rainfall: 120},
	{Name: "B", Temperature: 25.0, Rainfall: 80},
	{Name: "C", Temperature: 35.2, Rainfall: 200},
	{Name: "D", Temperature: 20.1, Rainfall: 50},
	{Name: "E", Temperature: 28.4, Rainfall: 100},
}

func FindCityWithHighestTemperature() CityClimate {
	highest := climateData[0]
	for _, city := range climateData {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}

func FindCityWithLowestTemperature() CityClimate {
	lowest := climateData[0]
	for _, city := range climateData {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

func CalculateAverageRainfall() float64 {
	totalRainfall := 0.0
	for _, city := range climateData {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(climateData))
}

func FilterCitiesByRainfall(threshold float64) []CityClimate {
	var filteredCities []CityClimate
	for _, city := range climateData {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}
func SearchCityByName(name string) (CityClimate, error) {
	for _, city := range climateData {
		if strings.EqualFold(city.Name, name) {
			return city, nil
		}
	}
	return CityClimate{}, errors.New("City not found")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Climate Data Analysis System")

	fmt.Println("\nAnalyzing temperatures...")
	highest := FindCityWithHighestTemperature()
	lowest := FindCityWithLowestTemperature()
	fmt.Printf("City with the highest temperature: %s (%.2f°C)\n", highest.Name, highest.Temperature)
	fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowest.Name, lowest.Temperature)

	fmt.Println("\nCalculating average rainfall...")
	averageRainfall := CalculateAverageRainfall()
	fmt.Printf("Average rainfall across all cities: %.2f mm\n", averageRainfall)

	fmt.Println("\nFilter cities by rainfall...")
	fmt.Print("Enter rainfall threshold (mm): ")
	thresholdInput, _ := reader.ReadString('\n')
	thresholdInput = strings.TrimSpace(thresholdInput)
	threshold, err := strconv.ParseFloat(thresholdInput, 64)
	if err != nil {
		fmt.Println("Invalid input for rainfall threshold.")
		return
	}
	filteredCities := FilterCitiesByRainfall(threshold)
	if len(filteredCities) == 0 {
		fmt.Println("No cities found with rainfall above the threshold.")
	} else {
		fmt.Println("Cities with rainfall above the threshold:")
		for _, city := range filteredCities {
			fmt.Printf("- %s: %.2f mm\n", city.Name, city.Rainfall)
		}
	}
	fmt.Println("\nSearch for a city by name...")
	fmt.Print("Enter city name: ")
	cityName, _ := reader.ReadString('\n')
	cityName = strings.TrimSpace(cityName)
	city, err := SearchCityByName(cityName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("City: %s\nTemperature: %.2f°C\nRainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
	}
}
