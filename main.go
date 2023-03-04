package main

import (
	"net/http"
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/aliziyacevik/exchanger/internal/repository/mongo"

)


func main() {
	r := chi.NewRouter()
	uri := "mongodb+srv://alizcev:lalalandAa.1@cluster0.qhgc1iz.mongodb.net/?retryWrites=true&w=majority"
	mr, err := mongo.NewMongoRepository(uri, "exchanger", 10)	
	if err != nil {
		log.Fatal(err)
	}
	err = mr.ImportInitialData()
	if err != nil {
		log.Fatal(err)
	}

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
