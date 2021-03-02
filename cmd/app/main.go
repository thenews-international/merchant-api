package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/spf13/viper"
	"go.opencensus.io/trace"

	"merchant/api/handler"
	"merchant/api/router"
	c "merchant/config"
	"merchant/model"
	"merchant/mysql"
	"merchant/server"
	"merchant/server/health"
	"merchant/server/requestlog"
	"merchant/util/logutil"
	"merchant/util/validator"
)

type customHealthCheck struct {
	mu      sync.RWMutex
	healthy bool
}

func (h *customHealthCheck) CheckHealth() error {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if !h.healthy {
		return errors.New("not ready yet!")
	}
	return nil
}

// @title Merchant service API
// @version 1.0
// @description This is the API documentation of the Merchant service
// @license.name MIT

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	logger, cleanup, err := logutil.NewLogger("*", "light-console", "stdout")
	if err != nil {
		panic(err)
	}
	defer cleanup()

	viper.SetConfigType("yaml")

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	var cfg c.Config

	if err := viper.ReadInConfig(); err != nil {
		logger.Warn(fmt.Sprintf("Error reading config file, %s", err))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		logger.Warn(fmt.Sprintf("Unable to decode into struct, %v", err))
	}

	addr := cfg.Server.Port

	db, err := mysql.New(&cfg.Database, cfg.Debug)
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	_ = db.AutoMigrate(&model.Merchant{}, &model.TeamMember{})

	var exporter trace.Exporter

	// Get validator
	appValidator := validator.New()

	srv := handler.New(db, appValidator, logger)

	mux := router.New(srv)

	// healthCheck will report the server is unhealthy for 10 seconds after
	// startup, and as healthy henceforth. Check the /healthz/readiness
	// HTTP path to see readiness.
	healthCheck := new(customHealthCheck)
	time.AfterFunc(10*time.Second, func() {
		healthCheck.mu.Lock()
		defer healthCheck.mu.Unlock()
		healthCheck.healthy = true
	})

	options := &server.Options{
		RequestLogger:         requestlog.NewNCSALogger(os.Stdout, func(error) {}),
		HealthChecks:          []health.Checker{healthCheck},
		TraceExporter:         exporter,
		DefaultSamplingPolicy: trace.AlwaysSample(),
		Driver:                &server.DefaultDriver{},
	}

	s := server.New(mux, options)
	fmt.Printf("Listening on %s\n", addr)

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		for {
			<-interrupt
			_ = s.Shutdown(context.Background())
			if sqlDB, err := db.DB(); err != nil {
				if err = sqlDB.Close(); err != nil {
					logger.Warn(err.Error())
				}
			}
		}
	}()
	err = s.ListenAndServe(fmt.Sprintf(":%s", addr))
	if err != nil {
		log.Fatal(err)
	}
}
