package codechallenge2

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	winner := Run()
	fmt.Println(winner)
	if winner.PlayerName == "No Winner" {
		t.Errorf("Did not get expected result. Wanted winner's player name, got No Winner")
	}
}
