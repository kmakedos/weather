package svc

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Service interface {
	GetTemp(ctx context.Context, city string) (*WeatherApi, error)
}

type Weather struct {
	apiKey string
	url    string
}

func GetNewWeatherService(url, apiKey string) *Weather {
	return &Weather{
		url:    url,
		apiKey: apiKey,
	}
}

func (w *Weather) GetTemp(ctx context.Context, city string) (*WeatherApi, error) {
	fullUrl := w.url + "key=" + w.apiKey + "&q=" + city
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Printf("ERROR: Invalid URL %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	WeatherApi := &WeatherApi{}
	err = json.NewDecoder(resp.Body).Decode(WeatherApi)
	if err != nil {
		log.Println("ERROR: Decoding response failed")
		return WeatherApi, err
	}
	return WeatherApi, nil
}
