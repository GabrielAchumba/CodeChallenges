package utils

import (
	"strconv"

	"github.com/GabrielAchumba/CodeChallenges/codechallenge2/models"
)

type IPlayerUtil interface {
	CreatePlayers(numberOfPlayers int) []models.Player
	IsATrail(player models.Player) bool
	IsSequence(player models.Player) bool
	IsAPair(player models.Player) bool
	GetTrailWinners(players []models.Player) []models.Player
	GetSequenceWinners(players []models.Player) []models.Player
	GetPairsWinners(players []models.Player) []models.Player
	GetAWinner(players []models.Player) models.Player
}

type PlayerUtil struct {
	cardUtil ICardUtil
}

func NewPlayerUtil(cardUtil ICardUtil) IPlayerUtil {
	return &PlayerUtil{
		cardUtil: cardUtil,
	}
}

func (impl PlayerUtil) CreatePlayers(numberOfPlayers int) []models.Player {

	players := make([]models.Player, 0)
	for i := 0; i < numberOfPlayers; i++ {
		players = append(players, models.Player{
			PlayerName: "Player " + strconv.Itoa(i+1),
			Cards:      impl.cardUtil.GenerateThreeCardRandomly(),
		})
	}

	return players
}

func IsContain(input string, list []string) bool {
	check := false

	for _, v := range list {
		if v == input {
			check = true
			break
		}
	}

	return check
}

func (impl PlayerUtil) IsATrail(player models.Player) bool {

	values := make([]string, 0, len(player.Cards.Cards))

	for _, value := range player.Cards.Cards {

		if !IsContain(value, values) {
			values = append(values, value)
		}
	}

	if len(values) == 1 {
		return true
	} else {
		return false
	}
}

func (impl PlayerUtil) IsSequence(player models.Player) bool {
	keys := make([]string, 0, len(player.Cards.Cards))

	for key := range player.Cards.Cards {

		keys = append(keys, key)
	}

	isSequence := false

	for i := 1; i < len(keys); i++ {
		keyBefore, _ := strconv.Atoi(keys[i-1])
		currentKey, _ := strconv.Atoi(keys[i])

		if (currentKey - keyBefore) == 1 {
			isSequence = true
		} else {
			isSequence = false
			break
		}
	}

	return isSequence
}

func (impl PlayerUtil) IsAPair(player models.Player) bool {

	values := make([]string, 0, len(player.Cards.Cards))

	for _, value := range player.Cards.Cards {

		values = append(values, value)
	}

	for j := 0; j < len(values); j++ {
		for i := 0; i < len(values); i++ {
			if j != i && values[j] == values[i] {
				return true
			}
		}
	}

	return false
}

func (impl PlayerUtil) GetTrailWinners(players []models.Player) []models.Player {

	trailPlayers := make([]models.Player, 0)
	for _, v := range players {

		check := impl.IsATrail(v)
		if check {
			trailPlayers = append(trailPlayers, v)
		}
	}

	return trailPlayers
}

func (impl PlayerUtil) GetSequenceWinners(players []models.Player) []models.Player {

	trailPlayers := make([]models.Player, 0)
	for _, v := range players {

		check := impl.IsSequence(v)
		if check {
			trailPlayers = append(trailPlayers, v)
		}
	}

	return trailPlayers
}

func (impl PlayerUtil) GetPairsWinners(players []models.Player) []models.Player {

	trailPlayers := make([]models.Player, 0)
	for _, v := range players {

		check := impl.IsAPair(v)
		if check {
			trailPlayers = append(trailPlayers, v)
		}
	}

	return trailPlayers
}

func (impl PlayerUtil) GetAWinner(players []models.Player) models.Player {

	winners := impl.GetTrailWinners(players)

	if len(winners) == 1 {
		return winners[0]
	}

	winners = impl.GetSequenceWinners(players)
	if len(winners) == 1 {
		return winners[0]
	}

	winners = impl.GetPairsWinners(players)
	if len(winners) == 1 {
		return winners[0]
	}

	return models.Player{}
}
