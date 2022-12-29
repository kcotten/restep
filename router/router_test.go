package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var result = `[{"item":"item1","quantity":"3"},{"item":"item2","quantity":"7"}]`

func Test_HandleRest(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	w := httptest.NewRecorder()
	r := Router{}
	r.Init()
	r.Router.ServeHTTP(w, req)

	// Check the status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check response
	if strings.Compare(w.Body.String(), result) == 0 {
		t.Errorf("handler returned incorrect json: got %v want %v",
			w.Body.String(), result)
	}
}

func Test_example(t *testing.T) {
	resp := example()
	val, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	// Check response
	if string(val[:]) != result {
		t.Errorf("handler returned incorrect json: got %v want %v",
			resp, result)
	}
}

func Test_Init(t *testing.T) {
	app := Router{}
	app.Init()

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	w := httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)

	// Check the status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code: got %v want %v",
			status, http.StatusOK)
	}
}
