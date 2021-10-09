package main

import "example/internal/app/example"

func main() {
	app := example.NewApp()
	app.Run()
}
