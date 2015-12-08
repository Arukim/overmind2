package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"sync"
)

type Node struct {
	Address string
	Status  string
}

var nodes = make(map[string]*Node)
var lock = sync.RWMutex{}

func GetNodes(w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()
	res := make([]Node, len(nodes))
	i := 0
	for _, node := range nodes {
		res[i] = *node
		i++
	}
	lock.RUnlock()
	w.WriteJson(&res)
}

func PutNode(w rest.ResponseWriter, r *rest.Request) {
	node := Node{}
	err := r.DecodeJsonPayload(&node)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if node.Address == "" {
		rest.Error(w, "node address required", 400)
	}

	lock.Lock()
	nodes[node.Address] = &node
	lock.Unlock()
	w.WriteHeader(http.StatusNoContent)
}
