package main

import (
	"blogs/routes"
	"blogs/utils"
	"net/http"
)

func main() {
	utils.ConnectDatabase()
	routes.InitializeRoutes()

	http.ListenAndServe(":8080", nil)
}
