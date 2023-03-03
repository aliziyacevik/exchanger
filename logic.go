package main

import (
	"os"
	"net/http"
)


type converterService struct {
	
}
	

func NewConverterService() ConverterService{
	return &converterService{
	}
}

func (c *converterService) Convert(from string, to string, amount int64) {
	key := os.Getenv("FIXER_IO_KEY")
        url := fmt.Sprintf("https://data.fixer.io/api/convert?access_key=%d&from=%s&to=%s&amount=%s", from, to, amount)

	resp, err := http.Post(url, "application/json")
	

}

