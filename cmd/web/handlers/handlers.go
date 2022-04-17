package handlers

import (
	"encoding/json"
	"github.com/matnich89/rbacfilter/cmd/internal/logger"
	"net/http"
	"time"
)

type Handlers struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Handlers {
	return &Handlers{logger: logger}
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	response := &Response{
		Message: "Hello Fro Mat",
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

type Response struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
