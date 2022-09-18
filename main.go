package main

import (
	"github.com/GabrielAchumba/CodeChallenges/codechallenge1"
	//"github.com/GabrielAchumba/CodeChallenges/codechallenge2"
)

func main() {

	terminalsFilePath := "./codechallenge1/files/terminals.json"
	transFareRateFilePath := "./codechallenge1/files/transfer_rates.json"
	startBusStop := 0
	stopBusStop := 3

	estimatedResult, _, _ := codechallenge1.Run(terminalsFilePath,
		transFareRateFilePath, startBusStop, stopBusStop)
	println("estimatedResult: ", estimatedResult)

	//codechallenge2.Main()
}
