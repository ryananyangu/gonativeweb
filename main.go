package main

import (
	"fmt"
	"net/http"
)

func SetupRouter() {

	// default_handler

	for path, handlers := range Routes {
		fmt.Printf("Registering paths : %s \n", path)
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			handler, ok := handlers[r.Method]
			if !ok {
				http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
				return
			}

			handler(w, r)
		})
	}

}

func main() {
	SetupRouter()
	http.ListenAndServe(":8080", nil)
}
