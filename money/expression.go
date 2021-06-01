package money

type Expression interface {
}

type sum struct {
	augend Money
	addend Money
}

func Sum(augend, addend Money) sum {
	return sum{augend, addend}
}
