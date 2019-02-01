package main

import (
	"fmt"

	"github.com/TheMickeyMike/lets-go/pkg/planets"
)

// App is backbone for application
type App struct {
	planet planets.Planet
}

// Initialize setup app
func (app *App) Initialize() {
	fmt.Printf("%-13s: %s\n", "App name", appName)
	fmt.Printf("%-13s: %s\n", "App version", appVersion)

	app.planet = planets.NewMoon("Moon 🌕")
}

// Run 3 2 1.. Let's go
func (app *App) Run() {
	fmt.Println("Let's Go! 🚀")
	fmt.Printf("Landing on the %s took: %v\n", app.planet, app.planet.Landing())
}
