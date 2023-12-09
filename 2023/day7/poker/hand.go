package poker

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
	Value int
}

const (
	HighCard     = 10000000000
	OnePair      = 20000000000
	TwoPair      = 30000000000
	ThreeOfAKind = 40000000000
	FullHouse    = 50000000000
	FourOfAKind  = 60000000000
	FiveOfAKind  = 70000000000
)

var listOfCards = []struct {
	Value  int
	Symbol rune
}{
	{2, '2'},
	{3, '3'},
	{4, '4'},
	{5, '5'},
	{6, '6'},
	{7, '7'},
	{8, '8'},
	{9, '9'},
	{10, 'T'},
	{11, 'J'},
	{12, 'Q'},
	{13, 'K'},
	{14, 'A'},
}

func ValueOfCard(sym rune, withJoker bool) int {
	if withJoker && sym == 'J' {
		return 1
	}
	for _, c := range listOfCards {
		if c.Symbol == sym {
			return c.Value
		}
	}
	panic("Invalid card")
}

type CardCount struct {
	Card  rune
	Count int
}

func ValueOfHand(hand string, withJoker bool) int {

	// hash count of each card
	cardsValue := 0
	countJoker := 0
	countMap := make(map[rune]int)
	for _, sym := range hand {
		if withJoker && sym == 'J' {
			countJoker++
		} else {
			countMap[sym]++
		}
		cardsValue = 100*cardsValue + ValueOfCard(sym, withJoker)
	}
	// make list of counts
	cards := make([]CardCount, 0)
	for sym, cnt := range countMap {
		cards = append(cards, CardCount{sym, cnt})
	}
	slices.SortFunc(cards, func(a, b CardCount) int {
		return b.Count - a.Count
	})

	// fmt.Println("Cards value for", hand, ":", cardsValue, countJoker)
	// for _, card := range cards {
	// 	fmt.Println(string(card.Card), ":", card.Count)
	// }

	// Sonderfall: nur Joker
	if withJoker && countJoker == 5 {
		return FiveOfAKind + cardsValue
	}

	if cards[0].Count+countJoker == 5 { // 11111, 1111J, 111JJ, 11JJJ, 1JJJJ
		return FiveOfAKind + cardsValue
	}

	if cards[0].Count+countJoker == 4 {
		return FourOfAKind + cardsValue
	}

	if cards[0].Count+countJoker == 3 && cards[1].Count == 2 {
		return FullHouse + cardsValue
	}

	if cards[0].Count+countJoker == 3 {
		return ThreeOfAKind + cardsValue
	}

	if cards[0].Count == 2 && cards[1].Count+countJoker == 2 {
		return TwoPair + cardsValue
	}

	if cards[0].Count+countJoker == 2 {
		return OnePair + cardsValue
	}

	return HighCard + cardsValue
}

func NewHand() *Hand {
	return &Hand{}
}

func MakeHand(line string, withJoker bool) *Hand {
	if len(line) < 7 || line[5] != ' ' {
		panic("Invalid line: " + line)
	}
	bid, err := strconv.Atoi(line[6:])
	if err != nil {
		panic(err)
	}

	hand := NewHand()
	hand.Cards = line[:5]
	hand.Value = ValueOfHand(hand.Cards, withJoker)
	hand.Bid = bid
	return hand
}

func Task1(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	hands := make([]*Hand, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		hands = append(hands, MakeHand(line, false))
	}

	slices.SortFunc(hands, func(a, b *Hand) int {
		return a.Value - b.Value
	})

	total := 0
	for i, hand := range hands {
		// fmt.Println(i+1, hand.Cards, hand.Value, hand.Bid)
		total += hand.Bid * (i + 1)
	}
	fmt.Println(filename, "Task 1 total:", total)
}

func Task2(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	hands := make([]*Hand, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		hands = append(hands, MakeHand(line, true))
	}

	slices.SortFunc(hands, func(a, b *Hand) int {
		return a.Value - b.Value
	})

	total := 0
	for i, hand := range hands {
		// fmt.Println(i+1, hand.Cards, hand.Value, hand.Bid)
		total += hand.Bid * (i + 1)
	}
	fmt.Println(filename, "Task 2 total:", total)
}
