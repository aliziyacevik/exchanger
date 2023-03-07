package domain 


type Repository interface {
	Find(base string) (*Currency, error)
	InsertInitialDataToDatabase() 
}



