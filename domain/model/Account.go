package model

import "strings"

type AccountID string

func (ai AccountID) ToString() string {
	return string(ai)
}

func (ai AccountID) IsEmpty() bool {
	return strings.TrimSpace(ai.ToString()) == ""
}

type Account struct {
	AccountID       AccountID
	baselineBalance Money
}
