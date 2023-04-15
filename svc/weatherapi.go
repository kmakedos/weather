package svc

type WeatherApi struct {
	Location struct {
		Name       string  `json:"name"`
		Region     string  `json:"region"`
		Country    string  `json:"country"`
		Latitude   float32 `json:"lat"`
		Longitude  float32 `json:"lon"`
		TimezoneId string  `json:"tz_id"`
		LocalEpoch int32   `json:"localtime_epoch"`
		Localtime  string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdateEpoch int32   `json:"last_updated_epoch"`
		LastUpdated     string  `json:"last_updated"`
		TempCelsius     float32 `json:"temp_c"`
		TempFahrenheit  float32 `json:"temp_f"`
		IsDay           int32   `json:"is_day"`
		Condition       struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph             float32 `json:"wind_mph"`
		WindKph             float32 `json:"wind_kph"`
		WindDegree          float32 `json:"wind_degree"`
		WindDir             string  `json:"wind_dir"`
		PressureMb          float32 `json:"pressure_mb"`
		PrecipMM            float32 `json:"precip_mm"`
		PrecipIn            float32 `json:"precip_in"`
		Humidity            float32 `json:"humidity"`
		Cloud               float32 `json:"cloud"`
		FeelsLikeCelsius    float32 `json:"feelslike_c"`
		FeelsLikeFahrenheit float32 `json:"feelslike_f"`
		VisibilityKm        float32 `json:"vis_km"`
		VisibilityMiles     float32 `json:"vis_miles"`
		UltraViolent        float32 `json:"uv"`
		GustMph             float32 `json:"gust_mph"`
		GustKph             float32 `json:"gust_kph"`
	} `json:"current"`
}
