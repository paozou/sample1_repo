package card

import (
	"fmt"
	"sort"
)

// Rank はカードの順位です。
type Rank int

func (rank Rank) String() (msg string) {
	switch rank {
	case 1:
		return "Ace"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return "Unknown"
	}
}

// Suit はカードのマークです。
type Suit int

func (suit Suit) String() (msg string) {
	switch suit {
	case 1:
		return "Spades"
	case 2:
		return "Hearts"
	case 3:
		return "Diamonds"
	case 4:
		return "Clubs"
	default:
		return "Unknown"
	}
}

// Card は順位とマークを持っています。
type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) String() (msg string) {
	msg = fmt.Sprintf("%s of %s", card.Rank, card.Suit)
	return msg
}

// Cards はカードのスライスです。
type Cards []Card

// Len はカードの長さです。
func (cards Cards) Len() (length int) {
	return len(cards)
}

// Swap は2マイのカードを交換する。
func (cards Cards) Swap(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}

// BySuit はマークでそーとしたカードのラッパーです。
type BySuit struct {
	Cards
}

// Less はマークでソートした時に最初のカードが2枚目のカードより小さいかどうか判定する。
func (b BySuit) Less(i, j int) (less bool) {
	return CompareByRank(b.Cards[i], b.Cards[j]) < 0
}

// ByRank は順位でソートしたカードのラッパーです。
type ByRank struct {
	Cards
}

// Less は数字でソートした時に最初のカードが2枚目のカードより小さいかどうか判定する。
func (b ByRank) Less(i, j int) bool {
	return CompareByRank(b.Cards[i], b.Cards[j]) < 0
}

// SortBySuit は絵柄でソートする。
func (cards Cards) SortBySuit() {
	sort.Sort(BySuit{cards})
}

// SortByRank は数字でソートする。
func (cards Cards) SortByRank() {
	sort.Sort(ByRank{cards})
}

// これらの定数はカードの絵柄
const (
	SPADES Suit = iota + 1
	HEARTS
	DIAMONDS
	CLUBS
)

// これらの定数はカードの順位
const (
	ACE Rank = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var rankingOfSuits = map[Suit]int{
	CLUBS:    1,
	DIAMONDS: 2,
	HEARTS:   3,
	SPADES:   4,
}

var rankingOfRanks = map[Rank]int{
	TWO:   1,
	THREE: 2,
	FOUR:  3,
	FIVE:  4,
	SIX:   5,
	SEVEN: 6,
	EIGHT: 7,
	NINE:  8,
	TEN:   9,
	JACK:  10,
	QUEEN: 11,
	KING:  12,
	ACE:   13,
}

// SetRankingOfSuits はカードのマークの順位をセットする。
func SetRankingOfSuits(suits []Suit) {
	for i, suit := range suits {
		rankingOfSuits[suit] = i
	}
}

// SetRankingOfRanks はカードの数字の順位をセットする。
func SetRankingOfRanks(ranks []Rank) {
	for i, rank := range ranks {
		rankingOfRanks[rank] = i
	}
}

// CompareBySuit は絵柄を比較し、差分を返却する。
func CompareBySuit(card1 Card, card2 Card) (diff int) {
	diff = rankingOfSuits[card1.Suit] - rankingOfSuits[card2.Suit]
	if diff == 0 {
		diff = rankingOfRanks[card1.Rank] - rankingOfRanks[card2.Rank]
	}
	return diff
}

// CompareByRank はカードの数字を比較し、差分を返却する。
func CompareByRank(card1 Card, card2 Card) (diff int) {
	diff = rankingOfRanks[card1.Rank] - rankingOfRanks[card2.Rank]
	if diff == 0 {
		diff = rankingOfSuits[card1.Suit] - rankingOfSuits[card2.Suit]
	}
	return diff
}
