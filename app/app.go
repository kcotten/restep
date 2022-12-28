package router

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	hr "github.com/julienschmidt/httprouter"
)

// Router struct
type App struct {
	Router *hr.Router
	Port   string
	logger log.Logger
}

// Initialize the router to return sample JSON
func (a *App) Init() {
	var ok bool
	a.Router = hr.New()
	ExampleRet = example()
	a.Router.HandlerFunc("GET", "/info", a.HandleRest)
	if a.Port, ok = os.LookupEnv("ROUTER_PORT"); !ok {
		a.Port = "8000"
	}
	a.logger = *log.Default()
	a.logger.Println("Application initialized with port: ", a.Port)
}

// Run the router at the provided address
func (a *App) Run(addr string) {}

// Router handler for incoming rest requests
func (a *App) HandleRest(w http.ResponseWriter, r *http.Request) {
	a.logger.Println("Handling request: ", r)
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
