package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Great first step")
	// io.WriteString(w,"Great first step")
}

func SampleErr(w http.ResponseWriter, r *http.Request) {
	// r.
	http.Error(w, "Method black listed", http.StatusBadRequest)
}
