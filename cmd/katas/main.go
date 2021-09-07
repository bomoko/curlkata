package main

import (
	"fmt"
	"github.com/bomoko/curlkata/internal/katas"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {

	r := mux.NewRouter()

	for _, k := range katas.Katas {
		r.HandleFunc(k.Name, k.Handler)
	}

	//set up help/description
	cwd, _ := os.Getwd()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(cwd)
		tmpl, err := template.ParseFiles(cwd + "/assets/templates/help.html")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, err.Error())
			return
		}
		writer.WriteHeader(http.StatusOK)
		tmpl.Execute(writer, katas.Katas)
		return
	})

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./assets/static"))))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
