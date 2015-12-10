package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"sync"
	"time"
)

var cache = make(map[string]*CacheItem)
var lock = sync.RWMutex{}

type CacheWriteRequest struct {
	Value     string
	DurationS int
}

type CacheItem struct {
	Value   *string
	Expires time.Time
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

	item := CacheItem{
		Value:   &data.Value,
		Expires: time.Now().Add(time.Duration(data.DurationS) * time.Second),
	}
	lock.Lock()
	cache[key] = &item
	lock.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func GetCache(w rest.ResponseWriter, r *rest.Request) {
	key := r.PathParam("key")

	lock.RLock()
	res := cache[key]
	lock.RUnlock()

	if res != nil && res.Expires.Before(time.Now()) {
		res = nil
	}
	w.WriteJson(&res)
}
