package katas

import (
	"fmt"
	"net/http"
)

type Kata struct {
	Name string //This corresponds with the endpoint
	Verb string //GET POST DELETE etc...
	Handler func(http.ResponseWriter, *http.Request)
}


var Katas = []Kata{}

func Hello() {
	fmt.Println("Hello world")
}