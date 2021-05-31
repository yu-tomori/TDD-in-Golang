package money

type Money struct {
	unit   string
	amount int
}

func (m Money) Equals(c currency) bool {
	return m.amount == c.quantity() && m.unit == c.name()
}

func (m Money) Times(multiplier int) Money {
	return Money{unit: m.unit, amount: m.amount * multiplier}
}

func (m Money) name() string {
	return m.unit
}

func (m Money) quantity() int {
	return m.amount
}

type currency interface {
	quantity() int
	name() string
}

func Dollar(amount int) Money {
	return Money{"USD", amount}
}

func Franc(amount int) Money {
	return Money{"CHF", amount}
}
