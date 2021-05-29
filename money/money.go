package money

type money struct {
	amount int
}

func (m money) equals(c currency) bool {
	return m.amount == c.quantity()
}

type currency interface {
	quantity() int
}

// Dollar currency
type dollar struct {
	money
}

func Dollar(amount int) dollar {
	return dollar{money{amount}}
}

func (d dollar) quantity() int {
	return d.amount
}

func (d dollar) times(multiplier int) dollar {
	return Dollar(d.amount * multiplier)
}

// Franc currency
type franc struct {
	money
}

func (f franc) quantity() int {
	return f.amount
}

func Franc(amount int) franc {
	return franc{money{amount}}
}

func (f franc) times(multiplier int) franc {
	return Franc(f.amount * multiplier)
}
