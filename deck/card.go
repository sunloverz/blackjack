package deck

import (
	"time"
	"math/rand"
	"sort"
	"fmt"
)

type Suit uint8

const (
	Spade Suit = iota 
	Diamond
	Club 
	Heart
	Joker
)

var suits = [...]Suit {Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two 
	Three
	Four
	Five
	Six
	Sevegn
	Eiht
	Nine
	Ten 
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
       return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func([]Card) []Card) []Card {
   var cards []Card
   for _, suit := range suits {
     for rank := Ace; rank <= King; rank++ {
		 cards = append(cards, Card{Suit: suit, Rank: rank})
	 }
   }

   for _, opt := range opts {
		cards = opt(cards)
	}

   return cards
}

func Sort(cards []Card) []Card {
  sort.Slice(cards, func(i, j int) bool {
	  return absRank(cards[i]) < absRank(cards[j])
  })
  return cards
}

func Shuffle(cards []Card) []Card {
    rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}

func Jokers(n int) func(cards []Card) []Card {
  return func(cards []Card) []Card {
	  for i := 0; i<n; i++ {
		  cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
	  }
	  return cards
	}
}

func FilterCards(f func(card Card) bool) func(cards []Card) []Card {
  return func(cards []Card) []Card {
	var ret []Card
	for _, c := range cards {
		if !f(c) {
			ret = append(ret, c)	  
		}
	}
	return ret
  }
}

func Deck(n int) func(cards []Card) []Card {
  return func(cards []Card) []Card {
	  var ret []Card
	  for i:=0; i<n; i++ {
		  ret = append(ret, cards...)
	  }
	  return ret
  }
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}