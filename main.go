package main

import (
	"net/http"
	"log"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/aliziyacevik/exchanger/internal/repository/mongo"
	s"github.com/aliziyacevik/exchanger/internal/service"
	"github.com/aliziyacevik/exchanger/internal/api"
)


func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	
	repo := chooseRepo()
	
	err :=repo.ImportInitialData()
	if err != nil {
		log.Fatal(err)
	}

	service := s.NewService(repo)
	handler := api.NewHandler(service, r)
	r.Post("/convert", handler.Post)	
	r.Get("/", handler.Get)
	
	server := http.Server{
		Addr:		":3000",
		Handler:	r,
	}
	
	log.Fatal(server.ListenAndServe())
}

func chooseRepo() s.Repository {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("must have a .env file")
	}

	switch os.Getenv("db") {
		case "mongo":
			mongoUrl := os.Getenv("MONGO_URL")
			mongoDb := os.Getenv("MONGO_DB")
			timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
			repo, err := mongo.NewMongoRepository(mongoUrl, mongoDb, timeout)
			if err != nil {
				log.Fatal(err)
			}
			return repo
		default:
			log.Fatal("Mongo db is only dbms supported.")
		}		
	return nil
}

