package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
    case card == "ace":
        return 11
    case card == "two":
        return 2
    case card == "three":
        return 3
    case card == "four":
        return 4
    case card == "five":
        return 5
    case card == "six":
        return 6
    case card == "seven":
        return 7
    case card == "eight":
        return 8
    case card == "nine":
        return 9
    case card == "ten" || card == "jack" || card == "queen" || card == "king":
        return 10
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    card1Val := ParseCard(card1)
    card2Val := ParseCard(card2)
    dealerCardVal := ParseCard(dealerCard)
    playerSum := card1Val + card2Val

    if playerSum == 22 {
        return "P"
    } else if playerSum == 21 {
        if dealerCardVal == 10 {
            return "S"
        } else if dealerCardVal == 11 {
            return "S"
        }
        return "W"
    } else if 17 <= playerSum && playerSum <= 20 {
        return "S"
    } else if 12 <= playerSum && playerSum <= 16 {
        if dealerCardVal >= 7 {
            return "H"
        }
        return "S"
    } else if playerSum <= 11 {
        return "H"
    }

    return "S"
}
