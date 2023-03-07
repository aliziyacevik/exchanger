package mongo 

import (
	"context"
	"time"
	"github.com/pkg/errors"
	"log"

	"github.com/aliziyacevik/exchanger/domain"
	"github.com/aliziyacevik/exchanger/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client		*mongo.Client
	database	string
	timeout		time.Duration
}


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


func NewMongoRepository(mongoURL string, mongoDB string, mongoTimeout int) (domain.Repository, error) {
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

func (mr *mongoRepository) Find(base string) (*domain.Currency, error) {
	coll := mr.client.Database(mr.database).Collection("currencies")
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	
	var result domain.Currency 
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

func (mr *mongoRepository) InsertInitialDataToDatabase() {
	var (
		symbols	   []interface{}
		currencies []interface{}
	)
	symbols = utils.InsertSymbolsToMemory()
	currencies = utils.InsertCurrenciesToMemory()

	mr.insertSymbolsToMongo(symbols)
	mr.insertCurrenciesToMongo(currencies)

	log.Println("Initial data inserted successfully..")
}

func (mr *mongoRepository) insertSymbolsToMongo(symbols []interface{}) {
	coll := mr.client.Database(mr.database).Collection("symbols")
	
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	_, err := coll.InsertMany(ctx, symbols)
	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertSymbolsToMongo"))
	}
	
}

func (mr *mongoRepository) insertCurrenciesToMongo(currencies []interface{}) {
	coll := mr.client.Database(mr.database).Collection("currencies")
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	
	_, err := coll.InsertMany(ctx, currencies)
	if err != nil {
		log.Fatal(errors.Wrap(err, "mongo.insertCurrenciesToMongo"))
	}
}





