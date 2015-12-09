package main

import (
	"net/http"
	"os"
)

func main() {
	adr := ":" + os.Args[1]
	Init(adr)
	http.ListenAndServe(adr, MakeHandler())
}
