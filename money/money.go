package money

type dollar struct {
	amount int
}

func Dollar(amount int) dollar {
	return dollar{amount}
}

func (d *dollar) times(multiplier int) dollar {
	return Dollar(d.amount * multiplier)
}
