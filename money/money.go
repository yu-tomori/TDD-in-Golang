package money

type money struct {
	unit   string
	amount int
}

func (m money) equals(c currency) bool {
	return m.amount == c.quantity() && m.unit == c.name()
}

type currency interface {
	quantity() int
	name() string
}

// Dollar currency
type dollar struct {
	money
}

func Dollar(amount int) dollar {
	return dollar{money{"Dollar", amount}}
}

// for currency interface
func (d dollar) quantity() int {
	return d.amount
}
func (d dollar) name() string {
	return d.unit
}

func (d dollar) times(multiplier int) dollar {
	return Dollar(d.amount * multiplier)
}

// Franc currency
type franc struct {
	money
}

// for currency interface
func (f franc) quantity() int {
	return f.amount
}
func (f franc) name() string {
	return f.unit
}

func Franc(amount int) franc {
	return franc{money{"Franc", amount}}
}

func (f franc) times(multiplier int) franc {
	return Franc(f.amount * multiplier)
}
