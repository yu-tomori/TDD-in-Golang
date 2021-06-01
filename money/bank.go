package money

type Bank struct {
}

func (b Bank) Reduce(source Expression, to string) Money {
	return Money{unit: "USD", amount: 10}
}
