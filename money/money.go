package money

type Money struct {
	unit   string
	amount int
}

func (m Money) Equals(c money) bool {
	return m.amount == c.quantity() && m.unit == c.currency()
}

func (m Money) Plus(c money) Money {
	return Money{unit: m.unit, amount: m.amount + c.quantity()}
}

func (m Money) Times(multiplier int) Money {
	return Money{unit: m.unit, amount: m.amount * multiplier}
}

func (m Money) currency() string {
	return m.unit
}

func (m Money) quantity() int {
	return m.amount
}

type money interface {
	quantity() int
	currency() string
}

func Dollar(amount int) Money {
	return Money{"USD", amount}
}

func Franc(amount int) Money {
	return Money{"CHF", amount}
}
