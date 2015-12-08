package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func GetStatus(w rest.ResponseWriter, r *rest.Request) {
	status := "online"
	w.WriteJson(status)
}
