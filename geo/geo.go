package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	r, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.New("not 200")
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil

}