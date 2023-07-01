package model

import gomoney "github.com/Rhymond/go-money"

type Money struct {
	amount *gomoney.Money
}

var ZeroUSD = NewMoney(0)

func NewMoney(amount int64) Money {
	return Money{
		amount: gomoney.New(amount, gomoney.USD),
	}
}

func (m Money) IsPositiveOrZero() bool {
	return m.amount.IsPositive() || m.amount.IsZero()
}

func (m Money) IsNegative() bool {
	return m.amount.IsNegative()
}

func (m Money) IsPositive() bool {
	return m.amount.IsPositive()
}

func (m Money) isGreaterThanOrEqualTo(money Money) (bool, error) {
	return m.amount.GreaterThanOrEqual(money.amount)
}

func (m Money) isGreaterThan(money Money) (bool, error) {
	return m.amount.GreaterThan(money.amount)
}

func (m Money) Add(otherMoney Money) (Money, error) {
	result, err := m.amount.Add(otherMoney.amount)
	if err != nil {
		return Money{}, err
	}

	return Money{
		amount: result,
	}, nil
}

func (m Money) Subtract(otherMoney Money) (Money, error) {
	result, err := m.amount.Subtract(otherMoney.amount)
	if err != nil {
		return Money{}, err
	}

	return Money{
		amount: result,
	}, nil
}

func (m Money) Positive() Money {
	return Money{
		amount: m.amount.Absolute(),
	}
}

func (m Money) Negative() Money {
	return Money{
		amount: m.amount.Negative(),
	}
}
