package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := Dollar(5)
	if Dollar(10) != five.Times(2) {
		t.Errorf("Dollar(5).times(2) should be Dollar(10)")
	}

	if Dollar(15) != five.Times(3) {
		t.Errorf("Dollar(5).times(3) should be Dollar(15)")
	}
}

func TestMoneySimpleAddition(t *testing.T) {
	five := Dollar(5)
	var sum Expression = five.Plus(five)
	reduced := Bank{}.Reduce(sum, "USD")
	if Dollar(10) != reduced {
		t.Errorf("10$ should be %v", reduced)
	}
}

func TestMoneyPlusReturnsSum(t *testing.T) {
	var five Money = Dollar(10)
	var result Expression = five.Plus(five)
	s, ok := result.(sum)

	if !ok {
		t.Errorf("underlying type of result should be %s", "sum")
	}

	if five != s.augend {
		t.Errorf("sum.%s doesn't work well", "augend")
	}

	if five != s.addend {
		t.Errorf("sum.%s doesn't work well", "addend")
	}
}

func TestMoneyReduceSum(t *testing.T) {
	var s Expression = Sum(Dollar(3), Dollar(4))
	var result Money = Bank{}.Reduce(s, "USD")

	if Dollar(7) != result {
		t.Errorf("3$ + 4$ should be 7$")
	}
}

func TestMoneyCurrency(t *testing.T) {
	if "USD" != Dollar(1).currency() {
		t.Errorf("Dollar's currency should be USD")
	}

	if "CHF" != Franc(1).currency() {
		t.Errorf("Franc's currency should be USD")
	}
}

func TestMoneyEquality(t *testing.T) {
	if !Dollar(5).Equals(Dollar(5)) {
		t.Errorf("dollar equivalence is lost")
	}

	if Dollar(5).Equals(Dollar(6)) {
		t.Errorf("dollar equivalence is lost")
	}

	if Dollar(5).Equals(Franc(5)) {
		t.Errorf("dollar(5) doesn't equal to franc(5)")
	}
}
