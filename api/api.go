package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"weather/svc"
)

type Server struct {
	svc svc.Service
}

func GetNewServer(svc svc.Service) *Server {
	return &Server{
		svc: svc,
	}
}

func (a *Server) Start(listenAddress string) {
	server := &http.Server{
		Addr: listenAddress,
	}

	http.HandleFunc("/temp", a.HandleTemp)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP Server error: %v", err)
		}
		log.Println("Stopped serving connections")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP Shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete")
}

func (a *Server) HandleTemp(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	weatherApi, err := a.svc.GetTemp(context.Background(), city)
	if err != nil {
		log.Printf("ERROR: Could not get Temp, %v\n", err)
		fmt.Fprintf(w, "ERROR: Could not get temp, probanly malformed url\n")
		return
	}
	fmt.Fprintf(w, "City: %s, Temp: %0.2f Celsius", city, weatherApi.Current.TempCelsius)
}
