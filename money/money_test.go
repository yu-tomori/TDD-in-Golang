package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := Dollar(5)
	five.times(2)
	if 10 != five.amount {
		t.Errorf("five.amount should be %d", 10)
	}
}
