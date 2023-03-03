package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {
	r := chi.NewRouter()
	uri := "mongodb+srv://alizcev:lalalandAa.1@cluster0.qhgc1iz.mongodb.net/?retryWrites=true&w=majority"
	mr, err := NewMongoRepository(uri, "lala", 10)	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mr.database)
	insertSymbols()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to the bank"))
	})
	
	r.Post("/convert", func(w http.ResponseWriter, r* http.Request) {
			from := chi.URLParam(r, "from")
			to := chi.URLParam(r, "to")
			amount := chi.URLParam(r, "amount")
			
			//service.Convert()
			w.Write([]byte(from + to +amount))
	})

	//http.ListenAndServe(":3000", r)

}
