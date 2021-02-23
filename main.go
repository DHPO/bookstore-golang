package main

import (
	"bookstore/bookstore/app"
)

func main() {
	application, err := app.CreateApp()
	if err != nil {
		panic(err)
	}
	application.Serve()
}
