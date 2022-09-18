package utils

import (
	"math/rand"
	"strconv"

	"github.com/GabrielAchumba/CodeChallenges/codechallenge2/models"
)

const NumberOfCardsPerPlayer int = 3

type ICardUtil interface {
	GenerateThreeCardRandomly() models.ThreeCards
}

type CardUtil struct {
}

func NewCardUtil() ICardUtil {

	return &CardUtil{}
}

func (impl CardUtil) GenerateThreeCardRandomly() models.ThreeCards {

	min := 1
	max := 4
	var threeCards models.ThreeCards
	threeCards.Cards = map[string]string{}
	counter := 0
	for {
		counter++
		if len(threeCards.Cards) == NumberOfCardsPerPlayer {
			break
		}

		newKey := strconv.Itoa(rand.Intn(max-min) + min)
		threeCards.Cards[newKey] = models.Cards[newKey]

	}

	return threeCards
}
