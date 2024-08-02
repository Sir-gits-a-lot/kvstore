package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Store map[string]string

var store Store = make(Store)

func getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, ok := store[key]
	if !ok {
		http.Error(w, "Key not found, please recheck your key", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	store[data.Key] = data.Value
	json.NewEncoder(w).Encode(map[string]string{"message": "Key-value pair inserted successfully"})
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	prefix := query.Get("prefix")
	suffix := query.Get("suffix")
	results := []string{}
	for key, _ := range store {
		if prefix != "" && !strings.HasPrefix(key, prefix) {
			continue
		}
		if suffix != "" && !strings.HasSuffix(key, suffix) {
			continue
		}
		results = append(results, key)
	}
	json.NewEncoder(w).Encode(map[string][]string{"Results": results})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get/{key}", getHandler).Methods("GET")
	r.HandleFunc("/set", setHandler).Methods("POST")
	r.HandleFunc("/search", searchHandler).Methods("GET")

	http.ListenAndServe(":8010", r)
}
