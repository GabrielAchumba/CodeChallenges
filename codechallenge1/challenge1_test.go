package codechallenge1

import "testing"

func TestRun(t *testing.T) {

	terminalsFilePath := "./files/terminals.json"
	transFareRateFilePath := "./files/transfer_rates.json"
	startBusStop := 0
	stopBusStop := 3

	estimatedResult, tarnsportFare, _ := Run(terminalsFilePath,
		transFareRateFilePath, startBusStop, stopBusStop)

	expect := 2.300000
	println(estimatedResult)
	if tarnsportFare != expect {
		t.Errorf("Did not get expected result. Wanted %v, got: %v\n", expect, tarnsportFare)
	}
}
