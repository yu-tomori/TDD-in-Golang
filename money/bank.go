package money

type Bank struct {
}

func (b Bank) Reduce(source Expression, to string) Money {
	return source.reduce(to)
}

func (b Bank) AddRate(from, to string, rate int) {
	return
}
