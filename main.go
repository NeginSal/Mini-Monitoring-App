package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	healthRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "health_requests_total",
			Help: "Total number of health check requests",
		},
	)

	responseLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "response_latency_seconds",
			Help:    "Response latency in seconds",
			Buckets: prometheus.LinearBuckets(0.01, 0.02, 10), // 10 buckets, each 0.02s wide, starting from 0.01s
		},
	)
)

func init() {
	prometheus.MustRegister(healthRequests)
	prometheus.MustRegister(responseLatency)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	healthRequests.Inc()

	delay := time.Duration(rand.Intn(200)) * time.Millisecond
	time.Sleep(delay)

	fmt.Fprint(w, `{"status":"ok"}`)

	duration := time.Since(start).Seconds()
	responseLatency.Observe(duration)
}

func main() {

	http.HandleFunc("/health", healthHandler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
