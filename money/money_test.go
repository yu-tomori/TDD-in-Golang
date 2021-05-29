package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := Dollar(5)
	if Dollar(10) != five.times(2) {
		t.Errorf("Dollar(5).times(2) should be Dollar(10)")
	}

	if Dollar(15) != five.times(3) {
		t.Errorf("Dollar(5).times(3) should be Dollar(15)")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := Franc(5)
	if Franc(10) != five.times(2) {
		t.Errorf("Franc(5).times(2) should be Franc(10)")
	}

	if Franc(15) != five.times(3) {
		t.Errorf("Franc(5).times(3) should be Franc(15)")
	}
}

func TestMoneyEquality(t *testing.T) {
	if !Dollar(5).equals(Dollar(5)) {
		t.Errorf("dollar equivalence is lost")
	}

	if Dollar(5).equals(Dollar(6)) {
		t.Errorf("dollar equivalence is lost")
	}

	if !Franc(5).equals(Franc(5)) {
		t.Errorf("franc equivalence is lost")
	}

	if Franc(5).equals(Franc(6)) {
		t.Errorf("franc equivalence is lost")
	}

}
