package service 

import (
//	"fmt"
//	"time"
)

type Transaction struct {
	query		Query	`json:"query"	bson:"query"`
	Result		int64	`json:"result"  bson:"result"`

}

type Query struct {
	From		Symbol 	`json:"from"	bson:"from"`
	To		Symbol  `json:"from"	bson:"from"`
	Amount		int64	`json:"amount"  bson:"value"`
}

type Symbol struct {
	Value		string `json:"value"	bson:"value"`
	Description	string `json:"desc"	bson:"desc"`
}

type Currency struct {
	Base		string			`bson:"base"` 
	Rates		map[string]float64	`bson:"rates"`
}


func NewTransaction(from string, to string, amount int64) (*Transaction) {
	transaction := &Transaction{}
	return transaction
}



