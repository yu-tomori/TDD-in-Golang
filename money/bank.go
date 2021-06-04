package money

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.Ldate)
}

type bank struct {
	rates map[CurrencyMap]int
}

func Bank() (b bank) {
	b.rates = make(map[CurrencyMap]int)
	return b
}

type CurrencyMap [2]string

func (b bank) Reduce(source Expression, to string) Money {
	return source.reduce(b, to)
}

func (b *bank) AddRate(from, to string, rate int) {
	if b.rates == nil {

		b.rates = make(map[CurrencyMap]int)
	}
	b.rates[CurrencyMap{from, to}] = rate
	return
}

func (b bank) Rate(from, to string) (rate int) {
	// if USD to USD
	if from == to {
		return 1
	}

	if r, ok := b.rates[CurrencyMap{from, to}]; ok {
		rate = r
	}
	return rate
}
