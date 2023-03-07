package utils 

import (
	"log"
	"os"
	"github.com/pkg/errors"
	"encoding/json"
	"encoding/csv"
	"io/ioutil"

	"github.com/aliziyacevik/exchanger/domain"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	
)

func ExistInEnv(checks []string) bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("There is no log file!!")
	}
	for _, check := range checks {
		if os.Getenv(check) == "" {
			log.Fatal("Couldn't find ", check, "in .env")
		}
	}
	return true
}

func InsertSymbolsToMemory() []interface{} {
	count := 0
	f, err := os.Open("symbols.csv")

	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertSymbolsToMemory"))
	}
	defer f.Close()
	
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertSymbolsToMemory"))
	}

	lines = lines[1:]
	var symbols []interface{}
	for _, line := range lines {
		value := line[0]
		desc := line[1]
		
		symbol := bson.D{{"value", value}, {"description", desc}}
		symbols = append(symbols, symbol)
		
		count ++
	}

	log.Printf("%d symbols has been loaded into memory", count)
	return symbols
}

func InsertCurrenciesToMemory() []interface{} {
	count := 0
	f, err := os.Open("currencies.json")

	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertCurrenciesToMemory error occured while opening json file"))
	}
	defer f.Close()
		
	jsonByte, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertCurrenciesToMemory error occured while reading file"))
	}
	
	var jsonData []domain.Currency
	err = json.Unmarshal(jsonByte, &jsonData)

	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertCurrenciesToMemory error occured while unmarshaling"))
	}
	
	var currencies []interface{}
	for _, data := range jsonData {
		currency := bson.D {
			{"base",	data.Base},
			{"rates",	data.Rates},
		}
		currencies = append(currencies, currency)
		count += 1
	}
	log.Printf("%d currencies has been loaded into memory.", count)
	return currencies

}
