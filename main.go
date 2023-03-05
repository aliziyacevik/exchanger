package main

import (
	"net/http"
	"log"
	"time"


	"github.com/aliziyacevik/exchanger/internal/repository/mongo"
	s"github.com/aliziyacevik/exchanger/internal/service"
	"github.com/aliziyacevik/exchanger/internal/api"
	"github.com/aliziyacevik/exchanger/internal/config"

)

func main() {
	serverConfig := config.LoadServerConfig()
	repoConfig := config.LoadRepositoryConfig()
	
	repo := chooseRepo(repoConfig)
	repo.InsertInitialDataToMongo()
	
	service := s.NewService(repo)
	handler := api.NewHandler(service)
	
	http.HandleFunc("/convert", api.AllowMethods(handler.Post, "POST"))
	http.HandleFunc("/", api.AllowMethods(handler.Get, "GET"))

	server := http.Server{
		Addr:		serverConfig.Addr,
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
	}
	
	log.Println("listening..")	
	log.Fatal(server.ListenAndServe())
}

func chooseRepo(cfg config.RepositoryConfiguration) s.Repository {
	switch cfg.Fetch("Database") {
		case "mongo":
			repo, err := mongo.NewMongoRepository(cfg.Fetch("MONGO_URL"), cfg.Fetch("MONGO_DB"), 15)
			if err != nil {
				log.Fatal(err)
			}
			return repo
		default:
			log.Fatal("Mongo db is only dbms supported.")
		}		
	return nil
}

