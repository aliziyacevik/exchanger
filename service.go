package main

type ConverterService interface {
	Convert (from string, to string, amount int64) (error)
	//Store(transaction *Transaction) (error)
}
