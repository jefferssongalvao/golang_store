package main

import (
	"net/http"
	"store/app/api/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
