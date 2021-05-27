package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := Dollar(5)
	if Dollar(10) != five.times(2) {
		t.Errorf("product.amount should be %d", 10)
	}

	if Dollar(15) != five.times(3) {
		t.Errorf("product.amount should be %d", 15)
	}
}

func TestMoneyEquality(t *testing.T) {
	if !Dollar(5).equals(Dollar(5)) {
		t.Errorf("dollar equivalence is lost")
	}

	if Dollar(5).equals(Dollar(6)) {
		t.Errorf("dollar equivalence is lost")
	}
}
