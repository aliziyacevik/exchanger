package main

import (
	"context"
	"time"
	"github.com/pkg/errors"
	"os"
	"log"
	"encoding/csv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Replace the placeholders with your credentials
const uri = "mongodb+srv://alizcev:lalalandAa.@cluster0.sample.mongodb.net/?retryWrites=true&w=majority"

var symbols	[] interface{}
var currencies 	[] interface{}

type mongoRepository struct {
	client		*mongo.Client
	database	string
	timeout		time.Duration
}

/*
func (mr *MongoRepository) Add(from string, to string, amount int64) (*currency){
	key := os.Getenv("FIXER_IO_KEY")
	uri := fmt.Sprintf("https://data.fixer.io/api/convert?access_key=%d&from=%s&to=%s&amount=%s") 
	
}
*/

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
	ApplyURI(mongoURL).
	SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "mongo.newMongoClient")
	}
	
	return client, nil
}	


func NewMongoRepository(mongoURL string, mongoDB string, mongoTimeout int) (*mongoRepository, error) {
	client, err := newMongoClient(mongoURL, mongoTimeout)
	
	if err != nil {
		return nil, err
	}
	repo := &mongoRepository{
		client:		client,
		timeout:	time.Duration(mongoTimeout) * time.Second,
		database:	mongoDB,
	}

	return repo, nil
}
/*
func (mr *mongoRepository)createIndex() {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	model := mongo.IndexModel{
		Keys:	bson.D{
				{"value"}
			}
	}
}
*/
func (mr *mongoRepository) ImportInitialData() error {
	insertSymbols()
	insertCurrencies()

	coll := mr.client.Database(mr.database).Collection("symbols")
	
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	_, err := coll.InsertMany(ctx, symbols)
	if err != nil {
		return errors.Wrap(err, "mongo.ImportInitialData")
	}

	log.Println("Initial data imported successfully..")
	return nil
}


func insertSymbols() error {
	count := 0
	f, err := os.Open("symbols.csv")

	if err != nil {
		return errors.Wrap(err,"mongo.insertSymbols") 
	}
	
	defer f.Close()
	
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return errors.Wrap(err, "mongo.insertSymbols")
	}

	lines = lines[1:]
	for _, line := range lines {
		value := line[0]
		desc := line[1]
		print(value, desc)	
		symbol := bson.D{{"value", value}, {"description", desc}}
		symbols = append(symbols, symbol)
		count ++
	}

	log.Printf("%d symbols has been loaded into memory", count)
	return nil
}

func insertCurrencies() {
	count := 0
	f, err := os.Open("currencies.json")

	if err != nil {
		return errors.Wrap(err,"mongo.insertCurrenciess") 
	}
	
	defer f.Close()
	
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return errors.Wrap(err, "mongo.insertCurrencies")
	}
	
	for _, line := range lines {
		fmt.Println(line)
		//var symbol Symbol
		//bsonBytes,_ := bson.Marshal(value)
		bson.Unmarshal(bsonBytes, &symbol)
		
		//val := symbol.Value
	}

}

func fetchCurrency() {
  url := "https://api.apilayer.com/fixer/convert?to={to}&from={from}&amount={amount}"

  client := &http.Client {}
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("apikey", "B8Pl9MAMErjNKChI8h0BEzMyb6Y986nS")

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
	if res.Body != nil {
    defer res.Body.Close()
  }
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}



