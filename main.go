package main

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
)

type HealthResponse struct {
	Status string `json:"status"`
}

var healthRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "health_requests_total",
		Help: "Total number of requests to the /health endpoint",
	},
)

func healthHandler(w http.ResponseWriter, r *http.Request) {

	healthRequests.Inc()

	response := HealthResponse{Status: "UP"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	prometheus.MustRegister(healthRequests)

	http.HandleFunc("/health", healthHandler)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
