package config 

import (
	"os"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type RepositoryConfiguration interface {
	Insert(string, value string)		
	Fetch(string)	string	
}

type ServerConfiguration struct {
	Addr		string
	ReadTimeout	int64
	WriteTimeout	int64
}

type repositoryConfiguration struct {
	Data		map[string]string
}

func LoadServerConfig() *ServerConfiguration {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file is not found.")
	}
	readTimeout, err := strconv.ParseInt(os.Getenv("READ_TIMEOUT"),10,64)
	writeTimeout, err2 := strconv.ParseInt(os.Getenv("WRITE_TIMEOUT"),10,64)
	if err != nil || err2 != nil{
		log.Fatal("Timeout values must be given as integers.")
	}

	return &ServerConfiguration{
		Addr:		os.Getenv("ADDR"),
		ReadTimeout:	readTimeout,
		WriteTimeout:	writeTimeout,
	}
}


func (c *repositoryConfiguration) Insert(name string, value string) {
	c.Data[name] = value
}
func (c *repositoryConfiguration) Fetch(name string) string {
	return c.Data[name]
}


func LoadRepositoryConfig() RepositoryConfiguration {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("There is no .env file!!")
	}
	
	switch os.Getenv("DB") {
         case "mongo":
		 c := &repositoryConfiguration{}
		 c.Insert("MONGO_URL", os.Getenv("MONGO_URL"))
		 c.Insert("MONGO_DB", os.Getenv("MONGO_DB"))
		 c.Insert("MONGO_TIMEOUT", "15")
		 c.Insert("Database", "mongo")
		 return c	
         
         default:
                  log.Fatal("Mongo db is only dbms supported.")
        }

	return nil
}
