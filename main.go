package main

import (
	"net/http"
	"log"
	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/aliziyacevik/exchanger/internal/repository/mongo"
	s"github.com/aliziyacevik/exchanger/internal/service"
)


func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	uri := "mongodb+srv://alizcev:lalalandAa.1@cluster0.qhgc1iz.mongodb.net/?retryWrites=true&w=majority"
	
	mr, err := mongo.NewMongoRepository(uri, "exchanger", 10)	
	if err != nil {
		log.Fatal(err)
	}
	err = mr.ImportInitialData()
	if err != nil {
		log.Fatal(err)
	}

	service := s.NewConverterService(mr)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to the bank"))
	})
	
	r.Post("/convert", func(w http.ResponseWriter, r* http.Request) {
			var q s.Query

			err := json.NewDecoder(r.Body).Decode(&q)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			log.Println("/convert", q.Amount, q.To)
			service.Convert(q)	
			w.Write([]byte(q.From + q.To))
	})

	http.ListenAndServe(":3000", r)


}
