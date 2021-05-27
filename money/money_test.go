package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := Dollar(5)
	product := five.times(2)
	if 10 != product.amount {
		t.Errorf("product.amount should be %d", 10)
	}

	product = five.times(3)
	if 15 != product.amount {
		t.Errorf("product.amount should be %d", 15)
	}
}

func TestMoneyEquality(t *testing.T) {
	if !Dollar(5).equals(Dollar(5)) {
		t.Errorf("dollar equivalence is lost")
	}
}
