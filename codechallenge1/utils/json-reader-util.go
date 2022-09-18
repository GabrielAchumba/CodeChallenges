package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/GabrielAchumba/CodeChallenges/codechallenge1/models"
)

type IJsonReaderUtil interface {
	ReadBustopDistances(filePath string) []models.OriginToBusStop
	ReadTransportFareRates(filePath string) []models.TransportFareRate
}

type JsonReaderUtil struct {
}

func NewJsonReaderUtil() IJsonReaderUtil {

	return &JsonReaderUtil{}
}

func (impl JsonReaderUtil) ReadBustopDistances(filePath string) []models.OriginToBusStop {

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error in loading bus stops", err)
	}

	result := make([]models.OriginToBusStop, 0)
	err = json.Unmarshal(fileContent, &result)

	return result
}

func (impl JsonReaderUtil) ReadTransportFareRates(filePath string) []models.TransportFareRate {

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error in loading transport fare rates", err)
	}

	result := make([]models.TransportFareRate, 0)
	err = json.Unmarshal(fileContent, &result)

	return result
}
