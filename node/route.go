package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func MakeHandler() http.Handler {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, _ := rest.MakeRouter(
		rest.Get("/status", GetStatus),
	)

	api.SetApp(router)

	return api.MakeHandler()
}
