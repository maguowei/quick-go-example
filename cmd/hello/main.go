package main

import "example/internal/app/hello"

func main() {
	app := hello.NewApp()
	app.Run()
}
