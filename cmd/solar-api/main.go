package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chiaf1/solar-api/internal/api"
	"github.com/chiaf1/solar-api/internal/config"
	"github.com/chiaf1/solar-api/internal/influx"
	"github.com/chiaf1/solar-api/internal/metrics"
)

const CONFIG_PATH = `./config.yaml`

func main() {
	// Load configs from file
	var conf config.Config
	err := conf.Load(CONFIG_PATH)
	if err != nil {
		log.Fatal(err)
	}
	err = conf.Validate()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config Loaded")

	// Creating an influx db client
	infClient := influx.New(conf.InfluxDB.Url, conf.InfluxDB.Token, conf.InfluxDB.Org, conf.InfluxDB.Bucket)

	// Creating the repo, service and handler layers to serve the API router
	repo := metrics.NewRepository(infClient)
	service := metrics.NewService(repo)
	handler := api.NewHandler(service)

	// Creating the router for the API endpoints
	router := api.NewRouter(handler)

	// Creating an HTTP server to handle graceful shutdowns
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Starting the server in a go routine
	go func() {
		log.Println("Api listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v", err)
		}
	}()

	//Waiting for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown signal received")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited cleanly")
}
