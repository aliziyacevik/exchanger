package main

import (
//	"os"
//	"net/http"
)


type converterService struct {
	
}
	

func NewConverterService() ConverterService{
	return &converterService{
	}
}

func (c *converterService) Convert(from string, to string, amount int64) (error){
	return nil

}

