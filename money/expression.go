package money

type Expression interface {
	reduce(bank, string) Money
}

type sum struct {
	augend Money
	addend Money
}

func Sum(augend, addend Money) sum {
	return sum{augend, addend}
}

func (s sum) reduce(b bank, to string) Money {
	amount := s.augend.amount + s.addend.amount
	return Money{unit: to, amount: amount}
}
