package service 

type ConverterService interface {
	Convert (q Query) (*Transaction, error)
	//Store(transaction *Transaction) (error)
}
