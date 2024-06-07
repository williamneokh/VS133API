package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LineTotals struct {
	Line1TotalIn  int `json:"line_1_total_in"`
	Line1TotalOut int `json:"line_1_total_out"`
}

var totals LineTotals

func main() {
	//Api that collects data
	http.HandleFunc("/countapi", receiveJSONHandler)
	//front end display number
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Create sample data
		totals := LineTotals{
			Line1TotalIn:  totals.Line1TotalIn,
			Line1TotalOut: totals.Line1TotalOut,
		}

		// Marshal the data to JSON
		jsonData, err := json.Marshal(totals)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set content type and write response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	fmt.Println("Server is running on port 3000...")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func dashboard(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello")
}

func receiveJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&totals)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	log.Println(totals.Line1TotalIn, totals.Line1TotalOut)
	fmt.Fprintf(w, "%v, %v", totals.Line1TotalIn, totals.Line1TotalOut)

	//var result map[string]interface{}
	//err := json.NewDecoder(r.Body).Decode(&result)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	fmt.Fprintf(w, "Error decoding JSON: %v", err)
	//	return
	//}

	//jsonString, _ := json.Marshal(result)

	//log.Println(string(jsonString))

	//fmt.Fprintf(w, "%v", string(jsonString))
}
