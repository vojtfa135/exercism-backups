package interest

const pctDivisor = 0.01
const pctProduct = 1.0

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interestRate float32

	switch {
	case balance < 0:
		interestRate = 3.213
	case balance >= 0 && balance < 1000.0:
		interestRate = 0.5
	case balance >= 1000.0 && balance < 5000.0:
		interestRate = 1.621
	case balance >= 5000.0:
		interestRate = 2.475
	}

	return interestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) * balance * pctDivisor
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return (float64(InterestRate(balance))*pctDivisor + pctProduct) * balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	if balance >= targetBalance {
		return 0
	}

	newBalance := balance + Interest(balance)
	return 1 + YearsBeforeDesiredBalance(newBalance, targetBalance)
}
