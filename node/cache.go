package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"sync"
)

var cache = make(map[string]*string)
var lock = sync.RWMutex{}

type CacheWriteRequest struct {
	Value     string
	DurationS string
}

func PutCache(w rest.ResponseWriter, r *rest.Request) {
	key := r.PathParam("key")
	data := CacheWriteRequest{}
	err := r.DecodeJsonPayload(&data)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if key == "" {
		rest.Error(w, "key required", 400)
	}

	lock.Lock()
	cache[key] = &data.Value
	lock.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func GetCache(w rest.ResponseWriter, r *rest.Request) {
	key := r.PathParam("key")

	lock.RLock()
	res := cache[key]
	lock.RUnlock()

	w.WriteJson(&res)
}
