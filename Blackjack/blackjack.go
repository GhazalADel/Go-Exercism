package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardsMapping := make(map[string]int)
	cardsMapping["ace"] = 11
	cardsMapping["two"] = 2
	cardsMapping["three"] = 3
	cardsMapping["four"] = 4
	cardsMapping["five"] = 5
	cardsMapping["six"] = 6
	cardsMapping["seven"] = 7
	cardsMapping["eight"] = 8
	cardsMapping["nine"] = 9
	cardsMapping["ten"] = 10
	cardsMapping["jack"] = 10
	cardsMapping["queen"] = 10
	cardsMapping["king"] = 10
	if pop, ok := cardsMapping[card]; !ok {
		return 0
	} else {
		return pop
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == card2 && card1 == "ace" {
		return "P"
	}
	sum := ParseCard(card2) + ParseCard(card1)
	if sum >= 17 && sum <= 20 {
		return "S"
	}
	if sum == 21 {
		if ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11 {
			return "S"
		} else {
			return "W"
		}
	}
	if sum <= 11 {
		return "H"
	}
	if sum >= 12 && sum <= 16 && ParseCard(dealerCard) >= 7 {
		return "H"
	}
	return "S"
}
