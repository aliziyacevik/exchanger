package main

import (
	"net/http"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/aliziyacevik/exchanger/internal/repository/mongo"
	s"github.com/aliziyacevik/exchanger/internal/service"
	"github.com/aliziyacevik/exchanger/internal/api"
)


func main() {
	repo := chooseRepo()
	repo.InsertInitialDataToMongo()
	

	service := s.NewService(repo)
	handler := api.NewHandler(service)
	
	http.HandleFunc("/convert", api.AllowMethods(handler.Post, "POST"))
	http.HandleFunc("/", api.AllowMethods(handler.Get, "GET"))

	server := http.Server{
		Addr:		":3000",
	}
	
	log.Println("listening..")	
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

