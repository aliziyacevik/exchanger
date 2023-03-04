package service 

type ConverterService interface {
	Convert (q Query) (error)
	//Store(transaction *Transaction) (error)
}
