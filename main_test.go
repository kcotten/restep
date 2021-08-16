package main

import (
	"net/http"
	"testing"
	"time"
)

// Functional test
func Test_main(t *testing.T) {
	go main()

	// allow for server to start
	time.Sleep(1 * time.Second)
	// url := "http://localhost:8000" + "/info"
	urlIpv6 := "http://[::1]:8000" + "/info"
	rr, err := http.Get(urlIpv6)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code
	if rr != nil {
		if status := rr.Status; status != "200 OK" {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}
