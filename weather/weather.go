package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Babahasko/go-weather/geo"
)
func GetWeather(geo *geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/"+ geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	r, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)
}