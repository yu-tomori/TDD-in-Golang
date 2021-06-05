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
	reduced := Bank().Reduce(sum, "USD")
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
	var result Money = Bank().Reduce(s, "USD")

	if Dollar(7) != result {
		t.Errorf("3$ + 4$ should be 7$")
	}
}

func TestMoneyReduceMoney(t *testing.T) {
	result := Bank().Reduce(Dollar(1), "USD")
	if Dollar(1) != result {
		t.Errorf("Reduced Money should be same Money")
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

func TestMoneyReduceMoneyDifferentCurrency(t *testing.T) {
	bank := Bank()
	rate := 2
	bank.AddRate("CHF", "USD", rate)
	if bank.rates[CurrencyMap{"CHF", "USD"}] != rate {
		t.Errorf("Rate setting doesn't work correctly. %d should be %d",
			bank.rates[CurrencyMap{"CHF", "USD"}], rate)
		t.Logf("rates: %v, len(rates): %d", bank.rates, len(bank.rates))
	}

	result := bank.Reduce(Franc(2), "USD")
	if Dollar(1) != result {
		t.Errorf("exchange between currencies failed: result %v", result)
	}
}

func TestMoneyMixedAddition(t *testing.T) {
	var fiveBucks Expression = Dollar(5)
	var tenFrans Expression = Franc(10)
	bank := Bank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrans), "USD")
	if Dollar(10) != result {
		t.Errorf("Mixed addition failed: result %v", result)
	}
}

func TestArrayEquality(t *testing.T) {
	a1 := [2]string{"ABC", "EFG"}
	a2 := [2]string{"ABC", "EFG"}
	if a1 != a2 {
		t.Errorf("Array equality is misconcieved")
	}
}

func TestIdentifyRate(t *testing.T) {
	if 1 != Bank().Rate("USD", "USD") {
		t.Errorf("Rate of USD to USD should be 1")
	}
}
