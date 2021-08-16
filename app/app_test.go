package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var result = []Info{{Item: "item1", Quantity: "3"}, {Item: "item2", Quantity: "7"}}

func Test_handleRest(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handleRest(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var resp *[]Info
	json.Unmarshal(rr.Body.Bytes(), resp)
	// Check response
	if resp == &result {
		t.Errorf("handler returned wrong json: got %v want %v",
			resp, result)
	}
}

func Test_example(t *testing.T) {
	resp := example()
	// Check response
	if &resp == &result {
		t.Errorf("handler returned wrong json: got %v want %v",
			resp, result)
	}
}

func Test_Init(t *testing.T) {
	app := App{}
	app.Init()

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
