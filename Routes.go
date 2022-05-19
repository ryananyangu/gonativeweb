package main

import (
	"net/http"

	"github.com/ryananyangu/gonativeweb/controllers"
)

var base_path = "/api/v1"

type HandlerFunction func(w http.ResponseWriter, r *http.Request)

var Routes = map[string]map[string]HandlerFunction{
	base_path + "/": {
		http.MethodGet: controllers.Index,
	},
	base_path + "/error": {
		http.MethodGet: controllers.SampleErr,
	},
}
