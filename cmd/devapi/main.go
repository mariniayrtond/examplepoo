package main

import (
	"log"
	"os"
	"users_example/cmd/devapi/furyapp"
	"users_example/internal/platform/environment"
)

func main() {
	env := environment.GetFromString(os.Getenv("GO_ENVIRONMENT"))

	dependencies, err := furyapp.BuildDependencies(env)
	if err != nil {
		log.Fatal("error at dependencies building", err)
	}

	app := furyapp.Build(dependencies)
	if err := app.Run(); err != nil {
		log.Fatal("error at furyapp startup", err)
	}
}
