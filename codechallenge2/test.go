package codechallenge2

import (
	"fmt"

	"github.com/GabrielAchumba/CodeChallenges/codechallenge2/utils"
)

func Main() {

	numberOfPlayers := 4
	cardUtil := utils.NewCardUtil()
	playerUtil := utils.NewPlayerUtil(cardUtil)
	players := playerUtil.CreatePlayers(numberOfPlayers)
	winner := playerUtil.GetAWinner(players)
	fmt.Println(winner)
}
