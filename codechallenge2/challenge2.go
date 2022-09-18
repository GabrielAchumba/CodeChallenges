package codechallenge2

import (
	"github.com/GabrielAchumba/CodeChallenges/codechallenge2/models"
	"github.com/GabrielAchumba/CodeChallenges/codechallenge2/utils"
)

func Run() models.Player {

	numberOfPlayers := 4
	cardUtil := utils.NewCardUtil()
	playerUtil := utils.NewPlayerUtil(cardUtil)
	players := playerUtil.CreatePlayers(numberOfPlayers)
	winner := playerUtil.GetAWinner(players)

	return winner
}
