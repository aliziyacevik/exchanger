package domain 

type service struct {
	repository	Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository:	repo,
	}
}

func (s *service) Convert(q Query) (*Transaction,error) {
	currency := &Currency{}	
	currency, err := s.repository.Find(q.From)
	if err != nil {
		return nil, err
	}
	result := calculateConversion(q.Amount, currency.Rates[q.To])
	
	transaction := &Transaction{
		Query:		q,
		Result:		result,
	}

	return transaction, nil
}

func calculateConversion(amount float64, rate float64) float64 {
	return amount * rate
}


