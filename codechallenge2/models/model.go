package models

var Cards = map[string]string{
	"1": "A",
	"2": "K",
	"3": "Q",
	"4": "J",
}

type ThreeCards struct {
	Cards map[string]string
}

type Player struct {
	PlayerName string
	Cards      ThreeCards
}
