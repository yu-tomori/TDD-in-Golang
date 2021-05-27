package money

type dollar struct {
	amount int
}

func Dollar(amount int) dollar {
	return dollar{amount}
}

func (d dollar) times(multiplier int) dollar {
	return Dollar(d.amount * multiplier)
}

func (d1 dollar) equals(d2 dollar) bool {
	return d1.amount == d2.amount
}

type franc struct {
	amount int
}

func Franc(amount int) franc {
	return franc{amount}
}

func (f franc) times(multiplier int) franc {
	return Franc(f.amount * multiplier)
}

func (f1 franc) equals(f2 franc) bool {
	return f1.amount == f2.amount
}
