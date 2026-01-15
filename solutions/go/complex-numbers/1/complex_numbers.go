package complexnumbers

import (
	"math"
)

// Define the Number type here.

type Number struct {
	a, b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	complexSum := Number{}
	complexSum.a = n1.a + n2.a
	complexSum.b = n1.b + n2.b
	return complexSum
}

func (n1 Number) Subtract(n2 Number) Number {
	complexRes := Number{}
	complexRes.a = n1.a - n2.a
	complexRes.b = n1.b - n2.b
	return complexRes
}

func (n1 Number) Multiply(n2 Number) Number {
	complexRes := Number{}
	complexRes.a = (n1.a * n2.a) - (n1.b * n2.b)
	complexRes.b = (n1.b * n2.a) + (n1.a * n2.b)
	return complexRes
}

func (n Number) Times(factor float64) Number {
	complexFactor := Number{}
	complexFactor.a = n.a * factor
	complexFactor.b = n.b * factor
	return complexFactor
}

func (n1 Number) Divide(n2 Number) Number {
	complexRes := Number{}
	complexRes.a = (n1.a*n2.a + n1.b*n2.b) / (n2.a*n2.a + n2.b*n2.b)
	complexRes.b = (n1.b*n2.a - n1.a*n2.b) / (n2.a*n2.a + n2.b*n2.b)
	return complexRes
}

func (n Number) Conjugate() Number {
	n.b = -n.b
	return n
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	var complexExp Number

	complexExp.a = math.Exp(n.a) * math.Cos(n.b)
	complexExp.b = math.Exp(n.a) * math.Sin(n.b)

	return complexExp
}
