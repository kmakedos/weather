package main

import (
	"fmt"
	"weather/api"
	"weather/logging"
	"weather/svc"
)

func main() {
	fmt.Println("Weather server starting")
	w := svc.GetNewWeatherService("https://api.weatherapi.com/v1/current.json?", "dab46697453448b0ab6162147231304")
	l := logging.GetNewLogger(w)
	a := api.GetNewServer(l)
	a.Start(":3000")
}
