package main

import (
	a "restep/app"
)

// Entry point for the rest ep
func main() {
	app := a.App{}
	app.Init()
	runHttpEndpoint(&app)
}
