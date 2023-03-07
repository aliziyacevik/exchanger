package mongo

import (
	"os"	

	"github.com/aliziyacevik/exchanger/domain"
	"github.com/aliziyacevik/exchanger/utils"
)

type mongoConfiguration struct {
	Data	map[string]string
}

func (c *mongoConfiguration) Insert(name string, value string) {
        c.Data[name] = value
}
func (c *mongoConfiguration) Fetch(name string) string {
        return c.Data[name]
}
func (c *mongoConfiguration) Database() string {
	return "mongo"
}

func InitializeMongoConfiguration() domain.RepositoryConfiguration{
	c := &mongoConfiguration{}
	c.Data = make(map[string]string)
	mongoConfigNames := []string{
		"MONGO_URL",
		"MONGO_DB",
		"MONGO_TIMEOUT",
	}
	
	if utils.ExistInEnv(mongoConfigNames) {
		c.Insert("MONGO_URL", os.Getenv("MONGO_URL"))
        	c.Insert("MONGO_DB", os.Getenv("MONGO_DB"))
		c.Insert("MONGO_TIMEOUT", "15")
	}

        return c
}




