package money

type Expression interface {
	Plus(Expression) Expression
	Times(int) Expression
	reduce(bank, string) Money
}

type sum struct {
	augend Expression
	addend Expression
}

func Sum(augend, addend Expression) sum {
	return sum{augend, addend}
}

func (s sum) Plus(exp Expression) Expression {
	return Sum(s, exp)
}

func (s sum) Times(multiplier int) Expression {
	return Sum(s.augend.Times(multiplier),
		s.addend.Times(multiplier))
}

func (s sum) reduce(b bank, to string) Money {
	amount := s.augend.reduce(b, to).amount +
		s.addend.reduce(b, to).amount
	return Money{unit: to, amount: amount}
}
