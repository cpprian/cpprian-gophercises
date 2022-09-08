package card_test

import (
	"testing"

	"github.com/cpprian/cpprian-gophercises/exercise9-deck-of-cards/pkg/card"
)

func TestShuffle(t *testing.T) {

	t.Run("shuffle deck from a sorted deck", func(t *testing.T) {
		cards := &card.Deck{
			*card.CreateCard(card.HEART, card.K),
			*card.CreateCard(card.DIAMOND, card.THREE),
			*card.CreateCard(card.DIAMOND, card.FIVE),
			*card.CreateCard(card.CLUB, card.TWO),
			*card.CreateCard(card.CLUB, card.A),
			*card.AddJoker(),
		}

		want := &card.Deck{
			*card.CreateCard(card.HEART, card.K),
			*card.CreateCard(card.DIAMOND, card.THREE),
			*card.CreateCard(card.DIAMOND, card.FIVE),
			*card.CreateCard(card.CLUB, card.TWO),
			*card.CreateCard(card.CLUB, card.A),
			*card.AddJoker(),
		}

		cards.Shuffle()
		AssertIsShuffled(t, cards, want)
	})
}
