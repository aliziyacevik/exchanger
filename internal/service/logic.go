package service

import (
	"log"
)

type converterService struct {
	repository	Repository
}

func NewConverterService(repo Repository) ConverterService {
	return &converterService{
		repository:	repo,
	}
}

func (c *converterService) Convert(q Query) error {
	log.Println("here", q.From, q.To, q.Amount)

	return nil
}

