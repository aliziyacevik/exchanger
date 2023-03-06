package service 

type Transaction struct {
	Query		Query	`json:"query"	bson:"query"`
	Result		float64	`json:"result"  bson:"result"`

}

type Query struct {
	From		string 	`json:"from"	bson:"from"`
	To		string  `json:"to"	bson:"to"`
	Amount		float64	`json:"amount"  bson:"amount"`
}

type Symbol struct {
	Value		string `json:"value"	bson:"value"`
	Description	string `json:"desc"	bson:"desc"`
}


type Currency struct {
	Base		string			`bson:"base"` 
	Rates		map[string]float64	`bson:"rates"`
}

