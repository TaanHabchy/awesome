package main

import (
	"awesomeProject/adventure_time"
	"net/http"
)

func main() {
	http.HandleFunc("/", adventure_time.StoryHandler)
	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
