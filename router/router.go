package router

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	hr "github.com/julienschmidt/httprouter"
)

// Router struct
type Router struct {
	Router *hr.Router
	Port   string
	logger log.Logger
}

// Initialize the router to return sample JSON
func (router *Router) Init() {
	var ok bool
	router.Router = hr.New()
	ExampleRet = example()
	router.Router.HandlerFunc("GET", "/info", router.HandleRest)
	if router.Port, ok = os.LookupEnv("ROUTER_PORT"); !ok {
		router.Port = "8000"
	}
	router.logger = *log.Default()
	router.logger.Println("Application initialized with port: ", router.Port)
}

// Run the router at the provided address
func (router *Router) Run(addr string) {}

// Router handler for incoming rest requests
func (router *Router) HandleRest(w http.ResponseWriter, r *http.Request) {
	router.logger.Println("Handling request: ", r)
	err := json.NewEncoder(w).Encode(ExampleRet)
	if err != nil {
		panic(err)
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
