package main

import (
	"github.com/gorilla/mux"
	"github.com/matnich89/rbacfilter/cmd/internal/logger"
	"github.com/matnich89/rbacfilter/cmd/web/handlers"
)

func main() {
	logger := logger.New()

	requestHandlers := handlers.New(logger)
	application := newApp(requestHandlers, mux.NewRouter(), logger)
	application.routes()
	err := application.serve()
	if err != nil {
		application.logger.OutputFatal(err.Error())
	}
}
