package katas

import (
	"fmt"
	"net/http"
)

var kata2 = Kata{
	Name:        "/kata2",
	Verb:        "GET",
	Description: "GET protected by basic auth - username/password = 'username/password'",
	Answer: "curl -u username:password http://localhost:8080/kata2",
	Handler: func(writer http.ResponseWriter, request *http.Request) {

		authorized := false

		username, password, _ := request.BasicAuth()
		fmt.Printf("Got username/password: %v/%v", username, password)
		if username == "username" && password == "password" {
			authorized = true
		}

		if !authorized {
			writer.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(writer, "FAIL: you haven't sent the basic auth header, try again")
			return
		}

		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "If you can see this, good job")
	},
}

func init() {
	Katas = append(Katas, kata2)
}
