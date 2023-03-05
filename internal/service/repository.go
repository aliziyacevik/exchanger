package service


type Repository interface {
	Find(base string) (*Currency, error)
	InsertInitialDataToMongo() 
}



