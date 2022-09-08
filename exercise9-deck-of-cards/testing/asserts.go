package card_test

import (
	"testing"

	"github.com/cpprian/cpprian-gophercises/exercise9-deck-of-cards/pkg/card"
)

func AssertSortedDeck(t testing.TB, cards, want *card.Deck) {
	cards_len := cards.Len()
	want_len := want.Len()
	
	AssertDeckSize(t, cards_len, want_len)

	for i := 0; i < cards_len; i++ {
		if (*cards)[i].Rank != (*want)[i].Rank || (*cards)[i].Suit != (*want)[i].Suit {
			t.Fatalf("incorrect position at %d: want %v, get %v\n", i, (*want)[i], (*cards)[i])
		}
	}
}

func AssertIsShuffled(t testing.TB, cards, want *card.Deck) {
	cards_len := cards.Len()
	want_len := want.Len()
	count := 0

	AssertDeckSize(t, cards_len, want_len)

	for i := 0; i < cards_len; i++ {
		if (*cards)[i].Rank == (*want)[i].Rank || (*cards)[i].Suit == (*want)[i].Suit {
			count++
		}
	}

	if count == cards_len {
		t.Fatalf("deck wasn't shuffled")
	}
}

func AssertDeckSize(t testing.TB, cards, want int) {
	if cards != want {
		t.Fatalf("incorrect size: want %d, get %d\n", want, cards)
	}
}