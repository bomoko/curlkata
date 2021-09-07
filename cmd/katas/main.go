package main

import (
	"fmt"
	"github.com/bomoko/curlkata/internal/katas"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()

	for _, k := range katas.Katas {
		r.HandleFunc(k.Name, k.Handler)
	}

	//set up help/description
	cwd, _ := os.Getwd()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		dat, err := os.ReadFile(cwd + "/templates/help.html")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, err.Error())
			return
		}
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, string(dat))
		return
	})

	http.Handle("/", r)
	http.Handle("/static/", http.FileServer(http.Dir(cwd + "/assets/static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
	//r :=
	//r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)
	//katas.Hello()

}
