package utils

import (
	"fmt"
	"math"

	"github.com/GabrielAchumba/CodeChallenges/codechallenge1/errors"
	"github.com/GabrielAchumba/CodeChallenges/codechallenge1/models"
)

type ITransportFareCalc interface {
	distanceBetweenBustStops(from models.OriginToBusStop, to models.OriginToBusStop) float64
	TransportFare(from models.OriginToBusStop, to models.OriginToBusStop) (string, error)
}

type TransportFareCalc struct {
	transportFareRates []models.TransportFareRate
}

func NewTransportFareCalc(transportFareRates []models.TransportFareRate) ITransportFareCalc {
	return &TransportFareCalc{
		transportFareRates: transportFareRates,
	}
}

func (impl TransportFareCalc) distanceBetweenBustStops(from models.OriginToBusStop, to models.OriginToBusStop) float64 {
	distance := math.Abs(from.Distance - to.Distance)
	return distance
}

func (impl TransportFareCalc) TransportFare(from models.OriginToBusStop, to models.OriginToBusStop,
) (string, error) {

	distance := impl.distanceBetweenBustStops(from, to)
	var cumulativeCost float64 = 0

	if impl.transportFareRates[0].From > 0 {
		return "", errors.Error("The start distance must be zero")
	}
	for _, item := range impl.transportFareRates {

		if distance > float64(item.To) {
			dist := math.Abs(float64(item.To) - float64(item.From))
			cumulativeCost += dist * item.Rate
		} else if distance > float64(item.From) && distance <= float64(item.To) {
			dist := math.Abs(distance - float64(item.From))
			cumulativeCost += dist * item.Rate
		}
	}

	result := from.Bus_stop + " > " + to.Bus_stop + " = Rs" + fmt.Sprintf("%f", cumulativeCost)

	return result, nil
}
