package main

import (
	"flag"
	"fmt"

	"github.com/Babahasko/go-weather/geo"
	"github.com/Babahasko/go-weather/weather"
)
func main() {
	fmt.Println("Cheking weather")
	city := flag.String("city", "Moscow", "City")
	format := flag.Int("format", 1, "Weather format")
	
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weatherData := weather.GetWeather(geoData, *format)
	fmt.Println(weatherData)
}
