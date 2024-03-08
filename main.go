package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CityInfo struct {
	Name    string   `json:"name"`
	Weather string   `json:"weather"`
	Status  []string `json:"status"`
}
type WeatherInfo struct {
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []CityInfo `json:"data"`
}

func fetchData(apiURL string, name string) ([]CityInfo, error) {
	var cities []CityInfo
	apiURL = apiURL + name

	//fetching data from each page using for loop
	pageNo := 1
	for {

		url := fmt.Sprintf("%s&page=%d", apiURL, pageNo)
		//url:="https://jsonmock.hackerrank.com/api/weather/search?name=B&page=2"
		// send an HTTP GET request to the specified URL (url)

		response, err := http.Get(url)

		if err != nil {
			fmt.Println("Error while fetching data:", err)
			return nil, err
		}
		body, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}
		//json.Unmarshal function is used to decode JSON data
		var weather_Info WeatherInfo
		err = json.Unmarshal(body, &weather_Info)
		if err != nil {
			return nil, err
		}

		cities = append(cities, weather_Info.Data...)
		if pageNo >= weather_Info.TotalPages {
			break
		}
		pageNo++

	}
	return cities, nil

}
func main() {
	apiURL := "https://jsonmock.hackerrank.com/api/weather/search?name="
	fmt.Println("Enter the name whose pages you want to access:")
	var name string
	fmt.Scan(&name)
	allCity, err := fetchData(apiURL, name)
	if err != nil {
		fmt.Println("Error while fetching data:", err)
		return
	}
	for _, city := range allCity {

		fmt.Println("Name:", city.Name)
		fmt.Println("Weather:", city.Weather)
		fmt.Println("Status:")
		for _, status := range city.Status {
			// fmt.Printf("%d: ",_)
			fmt.Println(status)
		}
		fmt.Println(" ")

	}
}
