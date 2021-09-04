package main

import (
	"fmt"
	"github.com/bomoko/curlkata/internal/katas"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Println(writer, "<h1>Hey biatch</h1>")
	})

	http.Handle("/", r)
	//r :=
	//r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)
	katas.Hello()
}
