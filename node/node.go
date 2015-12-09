package main

import (
	"github.com/Arukim/overmind2/models"
	"github.com/bndr/gopencils"
	"time"
)

var _address string

func Init(address string) {
	_address = address
	go ticker()
}

func ticker() {
	for {
		<-time.After(1 * time.Second)
		go reportToServer()
	}
}

func reportToServer() {
	api := gopencils.Api("http://localhost:8000")
	node := models.Node{Address: _address}
	api.Res("nodes").Put(node)
}
