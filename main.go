package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}


func getsavedmodels(w http.ResponseWriter, r *http.Request) {
 
    m := map[string]map[string][]string{
		"stocks": {
			"prediction_type":  {"trend_prediction_lstm", "sentiments_nlp"},
			"available_values": {"Yahoo", "Google"},
		},
		"recipe": {
			"prediction_type":  {"ingredents_list_texgenrnn", "process_texgenrnn"},
			"available_values": {"idli", "pasta"},
		},
	}
   	w.Header().Set("Content-Type", "application/json")
   	w.WriteHeader(http.StatusOK)
	jsonbyte, _ := json.Marshal(m)
	jsonString := string(jsonbyte)

	fmt.Fprint(w, jsonString)

}



func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/getsavedmodels", getsavedmodels)\
	log.Fatal(http.ListenAndServe(":8080", router))
}


