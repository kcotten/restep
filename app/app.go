package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Mux router struct
type App struct {
	Router *mux.Router
}

// Initialize the router to sample JSON
func (a *App) Init() {
	a.Router = mux.NewRouter()
	ExampleRet = example()
	a.Router.HandleFunc("/info", handleRest).Methods("GET")
}

// Run the router at the provided address
func (a *App) Run(addr string) {}

// Router handler for incoming rest requests
func handleRest(w http.ResponseWriter, r *http.Request) {
	if ExampleRet == nil {
		err := json.NewEncoder(w).Encode(example())
		if err != nil {
			panic(err)
		}
	} else {
		err := json.NewEncoder(w).Encode(ExampleRet)
		if err != nil {
			panic(err)
		}
	}
}

// Example go<=>json for endpoint to return
type Info struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}

// Example return value
var ExampleRet []Info

// Fill in the example struct with some arbitrary values
func example() []Info {
	Info := []Info{
		{Item: "item1", Quantity: "3"},
		{Item: "item2", Quantity: "7"},
	}
	return Info
}
