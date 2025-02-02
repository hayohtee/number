package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/classify-number", numberHandler)
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
	}

	var wg sync.WaitGroup

	
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
