package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	a.Router = mux.NewRouter()
	exampleRet = example()
	a.Router.HandleFunc("/info", handleRest).Methods("GET")
}

func (a *App) Run(addr string) {}

// Handle incoming rest requests
func handleRest(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(exampleRet)
}

type Info struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}

var exampleRet []Info

func example() []Info {
	Info := []Info{
		{Item: "item1", Quantity: "3"},
		{Item: "item2", Quantity: "7"},
	}
	return Info
}
