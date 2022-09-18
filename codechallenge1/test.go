package codechallenge1

import "github.com/GabrielAchumba/CodeChallenges/codechallenge1/utils"

func Main() {

	/* Sample Input:
	1>10
	7>12

	Sample Output:
	H.S.R. Layout > Vasantanagara = Rs.47 Bowring Institute > Mekhri Circle = Rs.17 */

	jsonReaderUtil := utils.NewJsonReaderUtil()
	terminalsFilePath := "./codechallenge1/files/terminals.json"
	transFareRateFilePath := "./codechallenge1/files/transfer_rates.json"
	busStops := jsonReaderUtil.ReadBustopDistances(terminalsFilePath)
	transportFareRates := jsonReaderUtil.ReadTransportFareRates(transFareRateFilePath)
	transportFareCalc := utils.NewTransportFareCalc(transportFareRates)
	estimatedResult, _ := transportFareCalc.TransportFare(busStops[1], busStops[3])
	println("estimatedResult: ", estimatedResult)

}
