package codechallenge1

import "github.com/GabrielAchumba/CodeChallenges/codechallenge1/utils"

func Run(terminalsFilePath string, transFareRateFilePath string,
	startBusStop int, stopBusStop int) (string, float64, error) {

	jsonReaderUtil := utils.NewJsonReaderUtil()
	busStops := jsonReaderUtil.ReadBustopDistances(terminalsFilePath)
	transportFareRates := jsonReaderUtil.ReadTransportFareRates(transFareRateFilePath)
	transportFareCalc := utils.NewTransportFareCalc(transportFareRates)
	estimatedResult, cumulativeCost, _ := transportFareCalc.TransportFare(busStops[1], busStops[3])

	return estimatedResult, cumulativeCost, nil
}
