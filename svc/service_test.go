package svc

import (
	"context"
	"log"
	"reflect"
	"testing"
)

func TestWeather_GetTemp(t *testing.T) {
	w := GetNewWeatherService("https://api.weatherapi.com/v1/current.json?", "dab46697453448b0ab6162147231304")
	cities := []string{"London", "Madrid", "Athens", "Nea Makri", "Marathonas"}
	for _, city := range cities {
		a, err := w.GetTemp(context.Background(), city)
		if err != nil {
			log.Printf("Error reading city: %s %v\n", city, err)
			t.Error(err)
		}
		if reflect.TypeOf(a.Current.TempCelsius).Kind() != reflect.Float32 {
			t.Error("Invalid result")
		}
	}

}
