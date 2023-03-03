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
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to the bank"))
	})
	
	r.Get("/room/{room-id}", func(w http.ResponseWriter, r *http.Request) {
			roomId := chi.URLParam(r, "room-id")
			w.Write([]byte(roomId))
	})
	
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			username := chi.URLParam(r, "username")
			password := chi.URLParam(r, "password")
			
			lala := username + password
			w.Write([]byte(lala))
	})

	r.Post("/convert" func(w http.ResponseWriter, r* http.Request) {
			from := chi.URLParam("from")
			to := chi.URLParam("to")
			amount := chi.URLParam("amount")

			w.Write([]byte(from + to +amount))

	})

	r.Post("/create-user", func(w http.ResponseWriter, r *http.Request) {
			username := chi.URLParam(r, "username")
			password := chi.URLParam(r, "password")
			
			um := UserManager{}
			
			if (um.CheckIfUsernameExist(username)) {
				// username exist return a proper response

			}
			um.CreateUser(username, password)
			

			w.Write([]byte(username))
	})

		
	

	//http.ListenAndServe(":3000", r)

}
