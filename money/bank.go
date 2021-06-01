package money

type Bank struct {
}

func (b Bank) Reduce(source Expression, to string) Money {
	s, _ := source.(sum)
	amount := s.augend.amount + s.addend.amount
	return Money{unit: to, amount: amount}
}
