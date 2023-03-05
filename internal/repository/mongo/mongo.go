package mongo 

import (
	"context"
	"time"
	"github.com/pkg/errors"
	"os"
	"log"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"

	s"github.com/aliziyacevik/exchanger/internal/service"


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
		return nil, errors.Wrap(err, "mongo.newMongoClient while connecting")
	}
	
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "mongo.newMongoClient while pinging")
	}
	
	return client, nil
}	


func NewMongoRepository(mongoURL string, mongoDB string, mongoTimeout int) (s.Repository, error) {
	client, err := newMongoClient(mongoURL, mongoTimeout)
	
	if err != nil {
		return nil, err
	}
	repo := &mongoRepository{
		client:		client,
		timeout:	time.Duration(mongoTimeout) * time.Second,
		database:	mongoDB,
	}
	log.Println("Mongo repository created and ready..")
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

func (mr *mongoRepository) Find(base string) (*s.Currency, error) {	
	coll := mr.client.Database(mr.database).Collection("currencies")
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	
	var result s.Currency 
	err := coll.FindOne(
		ctx,
		bson.D{
			{"base", base},
		},
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "repository.Find")
		}
	}
	return &result, nil 
}

func (mr *mongoRepository) InsertInitialDataToMongo() {
	insertSymbolsToMemory()
	insertCurrenciesToMemory()

	mr.insertSymbolsToMongo()
	mr.insertCurrenciesToMongo()

	log.Println("Initial data inserted successfully..")
}

func (mr *mongoRepository) insertSymbolsToMongo() {
	coll := mr.client.Database(mr.database).Collection("symbols")
	
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	_, err := coll.InsertMany(ctx, symbols)
	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertSymbolsToMongo"))
	}
	
}

func (mr *mongoRepository) insertCurrenciesToMongo() {
	coll := mr.client.Database(mr.database).Collection("currencies")
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	
	_, err := coll.InsertMany(ctx, currencies)
	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertCurrenciesToMongo"))
	}
}


func insertSymbolsToMemory(){
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
	for _, line := range lines {
		value := line[0]
		desc := line[1]
		//print(value, desc)	
		symbol := bson.D{{"value", value}, {"description", desc}}
		symbols = append(symbols, symbol)
		count ++
	}

	log.Printf("%d symbols has been loaded into memory", count)
}

func insertCurrenciesToMemory() {
	f, err := os.Open("currencies.json")
	defer f.Close()
	
	if err != nil {
		log.Fatal(errors.Wrap(err,"mongo.insertCurrenciess")) 
	}
	
	byteValue, _ := ioutil.ReadAll(f)
	//result := Currency{}

	err = json.Unmarshal([]byte(byteValue), &currencies)
	if err != nil {
		log.Fatal(errors.Wrap(err,"mongo.insertCurrenciesToMemory")) 
	}
	log.Println("currencies has been aded to the memory.")
}



