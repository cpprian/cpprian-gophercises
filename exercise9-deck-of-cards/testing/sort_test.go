package card_test

import (
	"testing"

	"github.com/cpprian/cpprian-gophercises/exercise9-deck-of-cards/pkg/card"
)

func TestSort(t *testing.T) {

	t.Run("test sorting with simple cards array", func(t *testing.T) {
		cards := &card.Deck{
			*card.CreateCard(card.DIAMOND, card.FIVE),
			*card.AddJoker(),
			*card.CreateCard(card.DIAMOND, card.THREE),
			*card.CreateCard(card.CLUB, card.TWO),
			*card.CreateCard(card.CLUB, card.A),
			*card.CreateCard(card.HEART, card.K),
		}

		want := &card.Deck{
			*card.CreateCard(card.HEART, card.K),
			*card.CreateCard(card.DIAMOND, card.THREE),
			*card.CreateCard(card.DIAMOND, card.FIVE),
			*card.CreateCard(card.CLUB, card.TWO),
			*card.CreateCard(card.CLUB, card.A),
			*card.AddJoker(),
		}

		cards.Sort(card.DefaultSort)

		AssertSortedDeck(t, cards, want)
	})
}