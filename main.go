package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/classify-number", numberHandler)

	srv := http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
		Addr:         ":4000",
	}

	slog.Info("starting server", "addr", "4000")

	if err := srv.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	numberParam := r.URL.Query().Get("number")
	if numberParam == "" {
		badRequestResponse(w, numberParam)
		return
	}

	number, err := strconv.Atoi(numberParam)
	if err != nil {
		badRequestResponse(w, numberParam)
		return
	}

	isPrimeChan := make(chan bool)
	isPerfectChan := make(chan bool)
	isArmstrongChan := make(chan bool)
	digitSumChan := make(chan int)

	go isArmstrong(number, isArmstrongChan)
	go isPerfect(number, isPerfectChan)
	go isPrime(number, isPrimeChan)
	go digitSum(number, digitSumChan)

	fact, err := getFunFact(number)
	if err != nil {
		http.Error(w, "the server encountered a problem when processing the request", http.StatusInternalServerError)
		return
	}

	var properties []string

	if <-isArmstrongChan {
		properties = append(properties, "armstrong")
	}

	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	resp := map[string]any{
		"number":     number,
		"is_prime":   <-isPrimeChan,
		"is_perfect": <-isPerfectChan,
		"properties": properties,
		"digit_sum":  <-digitSumChan,
		"fun_fact":   fact,
	}

	js, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "the server encountered a problem when processing the request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// badRequestResponse sends a JSON response with a bad request status code.
// The response includes the provided number and an error flag.
//
// Parameters:
//   - w: The http.ResponseWriter to write the response to.
//   - number: The number to include in the response.
//
// If there is an error marshalling the JSON response, it sends an internal server error response.
func badRequestResponse(w http.ResponseWriter, number string) {
	resp := map[string]any{
		"number": number,
		"error":  true,
	}

	js, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "the server encountered a problem when processing the request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(js)
}
