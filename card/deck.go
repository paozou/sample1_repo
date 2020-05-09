package card

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Deck はカードの集合
type Deck Cards

// 初期化
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Shuffle はデッキをシャッフルする。
func (deck Deck) Shuffle() {
	for i := len(deck); i > 0; i-- {
		randIndex := rand.Intn(i)
		deck[i-1], deck[randIndex] = deck[randIndex], deck[i-1]
	}
}

// Draw はカードを引く。
func (deck *Deck) Draw() (card Card, err error) {
	if len(*deck) == 0 {
		err = errors.New("couldn't draw, deck is empty")
		return card, err
	}
	card, *deck = (*deck)[0], (*deck)[1:]
	return card, err
}

// PutTop はカードを山札の1番上に置く。
func (deck *Deck) PutTop(card Card) {
	*deck = append(Deck{card}, *deck...)
}

// PutBottom はカードを山札の1番下に置く。
func (deck *Deck) PutBottom(card Card) {
	*deck = append(*deck, card)
}

// NewDeck は52マイの山札を返却する。
func NewDeck() (deck Deck) {
	for suitNumber := 1; suitNumber <= 4; suitNumber++ {
		for rankNumber := 1; rankNumber <= 13; rankNumber++ {
			deck = append(deck, Card{Rank: Rank(rankNumber), Suit: Suit(suitNumber)})
		}
	}
	fmt.Println(deck)
	return deck
}
