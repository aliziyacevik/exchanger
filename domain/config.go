package domain 

type RepositoryConfiguration interface {
	Insert(string, value string)		
	Fetch(string)			string	
	Database()			string
}

type ServerConfiguration struct {
	Addr		string
	ReadTimeout	int64
	WriteTimeout	int64
}


