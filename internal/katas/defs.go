package katas

import (
	"net/http"
)

type Kata struct {
	Name string //This corresponds with the endpoint
	Verb string //GET POST DELETE etc...
	Description string //Description of the task
	Answer string //Example curl that should solve the kata
	Handler func(http.ResponseWriter, *http.Request)
}


var Katas = []Kata{}

func getKatas() []Kata {
	return Katas
}
