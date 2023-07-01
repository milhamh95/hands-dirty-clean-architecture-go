package model

import (
	"errors"
	"time"
)

type ActivityID string

type Activity struct {
	ID              ActivityID
	ownerAccountID  AccountID
	sourceAccountID AccountID
	targetAccountID AccountID
	timestamp       time.Time
	amount          Money
}

func NewActivity(
	ownerAccountID AccountID,
	sourceAccountID AccountID,
	targetAccountID AccountID,
	timestamp time.Time,
	amountMoney Money) (Activity, error) {
	if ownerAccountID.IsEmpty() {
		return Activity{}, errors.New("owner account id is empty")
	}

	if sourceAccountID.IsEmpty() {
		return Activity{}, errors.New("source account id is empty")
	}

	if targetAccountID.IsEmpty() {
		return Activity{}, errors.New("target account id is empty")
	}

	if timestamp.IsZero() {
		return Activity{}, errors.New("timestamp is empty")
	}

	if amountMoney.amount == nil {
		return Activity{}, errors.New("amount money is empty")
	}

	return Activity{
		ownerAccountID:  ownerAccountID,
		sourceAccountID: sourceAccountID,
		targetAccountID: targetAccountID,
		timestamp:       timestamp,
		amount:          amountMoney,
	}, nil
}
