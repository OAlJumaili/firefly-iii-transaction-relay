package parser

import (
	"firefly-iii-transaction-relay/initialize"
	"strings"
	"time"
)

type Transaction struct {
	Type     string `json:"type"`
	Date     string `json:"date"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Vendor   string `json:"vendor"`
}

func ParseMessage(msg string) Transaction {
	var transaction Transaction

	if initialize.WithdrawRegex.MatchString(msg) {
		transaction.Type = "withdrawal"
	} else if initialize.DepositRegex.MatchString(msg) {
		transaction.Type = "deposit"
	}

	if matches := initialize.DateRegex.FindStringSubmatch(msg); len(matches) > 1 {
		t, err := time.Parse(matches[1], t)
		if err != nil {
			transaction.Date = t.Format("")
		}
	}

	if matches := initialize.AmountRegex.FindStringSubmatch(msg); len(matches) > 1 {
		transaction.Amount = strings.TrimSpace(matches[1])
	}

	if matches := initialize.CurrencyRegex.FindStringSubmatch(msg); len(matches) > 1 {
		transaction.Currency = strings.TrimSpace(matches[1])
	}

	if matches := initialize.VendorRegex.FindStringSubmatch(msg); len(matches) > 1 {
		transaction.Vendor = strings.TrimSpace(matches[1])
	}

	return transaction
}
