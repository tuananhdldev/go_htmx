package main

import (
	"gohtmx/model"
	"gohtmx/routes"
)

func main() {
	model.Setup()
	routes.SetupAndRun()
}
