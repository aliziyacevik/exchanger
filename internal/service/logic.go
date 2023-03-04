package service

type converterService struct {
	repository	Repository
}

func NewConverterService(repo Repository) ConverterService {
	return &converterService{
		repository:	repo,
	}
}

func (c *converterService) Convert(q Query) (*Transaction,error) {
	currency := &Currency{}	
	currency, err := c.repository.Find(q.From)
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


