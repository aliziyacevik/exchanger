package main

import (
	"net/http"
	"log"
	"time"
	"strconv"
	"os"

	"github.com/joho/godotenv"

	"github.com/aliziyacevik/exchanger/infra/repository/mongo"
	"github.com/aliziyacevik/exchanger/domain"
	"github.com/aliziyacevik/exchanger/infra/api"

)

func main() {
	serverConfig := LoadServerConfig()
	repoConfig := LoadRepositoryConfig()
	
	repo := chooseRepo(repoConfig)
	repo.InsertInitialDataToDatabase()
	
	service := domain.NewService(repo)
	handler := api.NewHandler(service)
	
	http.HandleFunc("/convert", api.AllowMethods(handler.Post, "POST"))
	http.HandleFunc("/", api.AllowMethods(handler.Get, "GET"))

	server := http.Server{
		Addr:		serverConfig.Addr,
		ReadTimeout:	time.Duration(serverConfig.ReadTimeout)   * time.Second,
		WriteTimeout:	time.Duration(serverConfig.WriteTimeout)  *time.Second,
	}
	
	log.Println("listening..")	
	log.Fatal(server.ListenAndServe())
}

func LoadServerConfig() *domain.ServerConfiguration {
        err := godotenv.Load(".env")
        if err != nil {
                log.Fatal(".env file is not found.")
        }
        readTimeout, err := strconv.ParseInt(os.Getenv("READ_TIMEOUT"),10,64)
        writeTimeout, err2 := strconv.ParseInt(os.Getenv("WRITE_TIMEOUT"),10,64)
        if err != nil || err2 != nil{
                log.Fatal("Timeout values must be given as integers.")
        }

        return &domain.ServerConfiguration{
                Addr:           os.Getenv("ADDR"),
                ReadTimeout:    readTimeout,
                WriteTimeout:   writeTimeout,
        }
}


func LoadRepositoryConfig() domain.RepositoryConfiguration {
        err := godotenv.Load(".env")
        if err != nil {
                log.Fatal("There is no .env file!!")
        }

        switch os.Getenv("DB") {
         case "mongo":
                 c := mongo.InitializeMongoConfiguration()
                 return c

         default:
                  log.Fatal("Mongo db is only dbms supported.")
        }

        return nil
}

func chooseRepo(cfg domain.RepositoryConfiguration) domain.Repository {
	switch cfg.Database() {
		case "mongo":
			repo, err := mongo.NewMongoRepository(cfg.Fetch("MONGO_URL"), cfg.Fetch("MONGO_DB"), 15)
			if err != nil {
				log.Fatal(err)
			}
			return repo
		default:
			log.Fatal("Mongo db is the only database that is  supported for now.")
		}		
	return nil
}

