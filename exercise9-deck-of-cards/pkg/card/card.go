package card

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	isJoker bool
	Suit
	Rank
}

type Deck []Card
type Game []Deck

func (c *Card) IsJoker() bool {
	return c.isJoker
}

func AddJoker() *Card {
	return &Card{
		isJoker: true,
	}
}

func CreateCard(suit Suit, rank Rank) *Card {
	return &Card{
		isJoker: false,
		Suit:    suit,
		Rank:    rank,
	}
}

func CreateDeck(num_joker int) *Deck {
	newDeck := &Deck{}

	for rank := MINRANK; rank < MAXRANK; rank++ {
		for suit := MINSUIT; suit < MINSUIT; suit++ {
			*newDeck = append(*newDeck, *CreateCard(suit, rank))
		}
	}

	for num_joker > 0 {
		*newDeck = append(*newDeck, *AddJoker())
		num_joker--
	}

	return newDeck
}

func (d *Deck) Len() int {
	var i int
	for i = range *d {
	}
	return i + 1
}

func New(num_decks int, num_joker int) *Game {
	newGame := &Game{}

	for num_decks > 0 {
		*newGame = append(*newGame, *CreateDeck(num_joker))
		num_decks--
	}

	return newGame
}

func (c *Card) String() string {
	if c.IsJoker() {
		return "Joker"
	}
	return fmt.Sprintf("%s of %s\n", CardName[c.Rank], CardType[c.Suit])
}

func (d *Deck) Sort(f func(first Card, second Card) bool) {
	is_swapped := true

	for is_swapped {
		is_swapped = false
		for j := 0; j < d.Len()-1; j++ {
			if f((*d)[j], (*d)[j+1]) {
				is_swapped = true
				d.Swap(j, j+1)
			}
		}
	}
}

func (d *Deck) Swap(i, j int) {
	temp := (*d)[i]
	(*d)[i] = (*d)[j]
	(*d)[j] = temp
}

func DefaultSort(first Card, second Card) bool {
	if first.IsJoker() {
		return true
	}
	if second.IsJoker() {
		return false
	}
	if first.Suit > second.Suit {
		return true
	}
	return first.Suit == second.Suit && first.Rank > second.Rank
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffled_deck := &Deck{}
	for _, m := range r.Perm(d.Len()){
		*shuffled_deck = append(*shuffled_deck, (*d)[m])
	}

	*d = *shuffled_deck
}

func (g *Game) ShuffleAllDecks() {
	for _, d := range *g {
		d.Shuffle()
	}
}

func (d *Deck) FilterDeck(rank ...Rank) {
	var filteredDeck Deck

	for _, r := range rank {
		for _, s := range *d {
			if s.Rank != r {
				filteredDeck = append(filteredDeck, s)
			}
		}
	}

	d = &filteredDeck
}
