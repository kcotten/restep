package main

import (
	a "restep/router"
)

// Entry point for the rest ep
func main() {
	router := a.Router{}
	router.Init()
	runHttpEndpoint(&router)
}
