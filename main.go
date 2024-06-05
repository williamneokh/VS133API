package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type payload struct {
}

func main() {
	http.HandleFunc("/countapi", receiveJSONHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func receiveJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var result map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	fmt.Println(result)
	fmt.Fprintf(w, "%v", result)
}
