package api

import (
	"testing"
	"weather/svc"
)

func TestServer_HandleTemp(t *testing.T) {
	w := svc.GetNewWeatherService("https://api.weatherapi.com/v1/current.json?", "dab46697453448b0ab6162147231304")
	a := GetNewServer(w)
	a.Start(":4000")
}
