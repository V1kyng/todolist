package main

import (
	"log"
	"net/http"
	"todolist/src/routes"
)

func main() {
	r := routes.SetupRouter()
	http.ListenAndServe(":8000", r)
	log.Println("Server begined working...")
}
