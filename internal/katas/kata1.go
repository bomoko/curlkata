package katas

import (
	"fmt"
	"net/http"
)

var kata1 = Kata{
	Name:    "/kata1",
	Verb:    "GET",
	Handler: func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Got one")
	},
}


func init() {
	Katas = append(Katas, kata1)
}