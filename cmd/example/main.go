package main

import (
    "example/internal/example"
    _ "go.uber.org/automaxprocs"
)

func main() {
	app := example.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
