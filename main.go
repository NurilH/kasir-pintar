package main

import (
	middlewares "test_kasir_pintar/middleware"
	"test_kasir_pintar/routes"
)

func main() {
	e := routes.New()
	// take notes http method activity
	middlewares.LogMiddlewares(e)
	// start on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
