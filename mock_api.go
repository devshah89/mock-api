package main

import (
	"encoding/json" 
	"log"           
	"net/http"      
)


type AuthResponse struct {
	Valid bool `json:"valid"` 
}


func validateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	deviceID := r.URL.Query().Get("deviceID")
	if deviceID != "" {
		log.Printf("Extracted deviceID from query: %s", deviceID)
	} else {
		log.Println("No deviceID found in query parameters.")
	}

	w.Header().Set("Content-Type", "application/json")

	response := AuthResponse{Valid: true}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("Responded with: {\"valid\": %t}", response.Valid)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/validate-device", validateDeviceHandler)

	serverAddress := "localhost:9000" 

	log.Printf("Starting mock authentication API server on http://%s", serverAddress)

	err := http.ListenAndServe(serverAddress, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
