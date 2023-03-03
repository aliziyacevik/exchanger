package main

import (
//	"fmt"
//	"time"
)

type Transaction struct {
	query		Query	`json:"query"	bson:"query"`
	Result		int64	`json:"result"  bson:"result"`

}

type Query struct {
	From		string 	`json:"from" 	bson:"from"`
	To		string  `json:"to"  	bson:"to"`
	Amount		int64	`json:"amount"  bson:"value"`
}

func (t *Transaction) Read(p []byte) (n int, err error) {
		
}


func NewTransaction() (*Transaction) {
	transaction := &Transaction {
	}

	return transaction
}

func NewTransaction(from string, to string, amount int64) (*Transcation) {
	transaction := &Transaction{
		from:	from,
		to:	to,
		amount	amount,
	}

	return transaction
}



