package main

import (
	"github.com/sergeyglazyrindev/es_gomonolith_example/blueprint/estweet"
	"github.com/sergeyglazyrindev/go-monolith"
	"os"
)

func main() {
	environment := os.Getenv("environment")
	if environment == "" {
		environment = "dev"
	}
	app1 := gomonolith.NewApp(environment, true)
	// please configure everything you need for your project here.
	app1.BlueprintRegistry.Register(estweet.ConcreteBlueprint)
	// next two lines are mandatory for uadmin to determine your blueprints and everything else.
	app1.Initialize()
	app1.InitializeRouter()
	// and here is the command handler that will be used for any functionality you may need.
	app1.ExecuteCommand()
}
