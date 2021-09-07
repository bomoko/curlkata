package katas

import (
	"fmt"
	"net/http"
)

var kata1 = Kata{
	Name:    "/kata1",
	Verb:    "GET",
	Answer: "curl http://localhost:8080/kata1",
	Description: "A simple GET",
	Handler: func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "If you can see this, good job")
	},
}

func init() {
	Katas = append(Katas, kata1)
}