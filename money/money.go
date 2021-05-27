package money

type dollar struct {
	amount int
}

func Dollar(amount int) dollar {
	return dollar{amount}
}
