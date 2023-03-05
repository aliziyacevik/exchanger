package service 

type Service interface {
	Convert (q Query) (*Transaction, error)
	//Store(transaction *Transaction) (error)
}
