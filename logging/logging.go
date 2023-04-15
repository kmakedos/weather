package logging

import (
	"context"
	"fmt"
	"time"
	"weather/svc"
)

type Logger struct {
	svc svc.Service
}

func GetNewLogger(svc svc.Service) *Logger {
	return &Logger{
		svc: svc,
	}
}

func (l *Logger) GetTemp(ctx context.Context, city string) (svc *svc.WeatherApi, err error) {
	defer func(start time.Time) {
		fmt.Printf("city=%s, temp=%f err=%v, took=%v\n", city, svc.Current.TempCelsius, err, time.Since(start))
	}(time.Now())

	return l.svc.GetTemp(ctx, city)
}
