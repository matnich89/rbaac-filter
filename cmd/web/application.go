package main

import (
	"github.com/gorilla/mux"
	"github.com/matnich89/rbacfilter/cmd/internal/logger"
	"github.com/matnich89/rbacfilter/cmd/web/handlers"
	"net/http"
	"sync"
)

type app struct {
	wg       sync.WaitGroup
	handlers *handlers.Handlers
	mux      *mux.Router
	logger   *logger.Logger
}

func newApp(handlers *handlers.Handlers, mux *mux.Router, logger *logger.Logger) *app {
	return &app{
		wg:       sync.WaitGroup{},
		handlers: handlers,
		mux:      mux,
		logger:   logger,
	}
}

func (a *app) routes() {
	a.logger.OutputInfo("setting up routes....")
	a.mux.HandleFunc("/health", a.handlers.HealthCheck).Methods(http.MethodGet)
	/*
		As the requirement specified  using a payload rather than query params no choice but to make this a POST or PUT,
		Though I feel the operation is really a GET,
	*/
	a.mux.HandleFunc("/", a.handlers.Search).Methods(http.MethodGet)
}
