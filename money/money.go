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

// Dollar currency
type dollar struct {
	Money
}

func Dollar(amount int) Money {
	return Money{"USD", amount}
}

// for currency interface
func (d dollar) quantity() int {
	return d.amount
}
func (d dollar) name() string {
	return d.unit
}

// Franc currency
type franc struct {
	Money
}

// for currency interface
func (f franc) quantity() int {
	return f.amount
}
func (f franc) name() string {
	return f.unit
}

func Franc(amount int) Money {
	return Money{"CHF", amount}
}
