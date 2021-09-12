package katas

import (
	"fmt"
	"net/http"
)

const headerName = "x-curlkata-kata3"
const headerValue = "kata3"

var kata3 = Kata{
	Name:        "/kata3",
	Verb:        "GET",
	Description: "Tests sending arbitrary headers as part of the request",
	Answer:      "curl localhost:8080/kata3 --header \"x-curlkata-kata3: kata3\"",
	Handler: func(writer http.ResponseWriter, request *http.Request) {
		sentHeader := false

		if request.Header.Get(headerName) == headerValue {
			sentHeader = true
		}

		if sentHeader {
			writer.WriteHeader(http.StatusOK)
			fmt.Fprint(writer, "Congrats, you sent the correct header")
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "No luck, try sending the header %v with value %v", headerName, headerValue)
		}
	},
}


func init() {
	Katas = append(Katas, kata3)
}