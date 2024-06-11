package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the libraGo management system")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	setupRoutes()
	fmt.Println("livrago server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("failed to start server", err)
	}
}
