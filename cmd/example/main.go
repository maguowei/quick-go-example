package main

import "github.com/maguowei/example/internal/example"

func main() {
	app := example.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
