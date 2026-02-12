package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Uptime    string    `json:"uptime"`
	Timestamp time.Time `json:"timestamp"`
}

var startTime time.Time

func init() {
	startTime = time.Now()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	
	response := Response{
		Message:   "Welcome to Hello Web Server! Visit /api/hello for a greeting.",
		Status:    "success",
		Timestamp: time.Now(),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	
	response := Response{
		Message:   fmt.Sprintf("Hello, %s!", name),
		Status:    "success",
		Timestamp: time.Now(),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Health check from %s", r.RemoteAddr)
	
	uptime := time.Since(startTime)
	
	response := HealthResponse{
		Status:    "healthy",
		Uptime:    uptime.String(),
		Timestamp: time.Now(),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	
	port := ":8080"
	log.Printf("Starting Hello Web Server on port %s", port)
	log.Printf("Endpoints:")
	log.Printf("  - GET /           : Welcome message")
	log.Printf("  - GET /api/hello  : Hello endpoint (try ?name=YourName)")
	log.Printf("  - GET /health     : Health check")
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
