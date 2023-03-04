package service

type converterService struct {
	repository	Repository
}

func NewConverterService(repo Repository) ConverterService {
	return &converterService{
		repository:	repo,
	}
}

func (c *converterService) Convert(from string, to string, amount int64) error {
	return nil
}

