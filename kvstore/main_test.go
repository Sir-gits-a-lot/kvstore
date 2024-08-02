package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetHandler(t *testing.T) {

	store["testkey"] = "testvalue"
	r := mux.NewRouter()
	r.HandleFunc("/get/{key}", getHandler).Methods("GET")

	// 1st Test Case: Successful retrieval
	req, err := http.NewRequest("GET", "/get/testkey", nil)
	if err != nil {
			t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusOK)
	}

	var result map[string]string
	json.Unmarshal(rr.Body.Bytes(), &result)
	if result["value"] != "testvalue" {
			t.Errorf("handler returned unexpected body: got %v want %v", result, map[string]string{"value": "testvalue"})
	}

	// 2nd Test Case: Key not found
	req, err = http.NewRequest("GET", "/get/nonexistent", nil)
	if err != nil {
			t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusNotFound)
	}
}

func TestSetHandler(t *testing.T) {
	// Setup
	r := mux.NewRouter()
	r.HandleFunc("/set", setHandler).Methods("POST")

	// Test case 1: Successful setting
	data := map[string]string{"key": "newkey", "value": "newvalue"}
	jsonData, err := json.Marshal(data)
	if err != nil {
			t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonData))
	if err != nil {
			t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusOK)
	}

	if _, ok := store["newkey"]; !ok {
			t.Error("key not stored")
	}
}
