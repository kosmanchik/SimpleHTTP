package main

import (
	"http-server/api"
	"net/http"
)

func main() {
	server := api.InitServer()
	http.ListenAndServe(":8080", server)
}
