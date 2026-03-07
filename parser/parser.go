package parser

import (
	"firefly-iii-transaction-relay/initialize"
	"strings"
	"time"
)

type Transaction struct {
	Type         string  `json:"type"`
	Date         string  `json:"date"`
	Amount       string  `json:"amount"`
	Currency     string  `json:"currency_code"`
	SourceVendor *string `json:"source_name,omitempty"`
	DestVendor   *string `json:"destination_name,omitempty"`
}

func ParseMessage(msg string) Transaction {
	var transaction Transaction

	if initialize.WithdrawRegex.MatchString(msg) {
		transaction.Type = "withdrawal"
	} else if initialize.DepositRegex.MatchString(msg) {
		transaction.Type = "deposit"
	}

	if matches := initialize.DateRegex.FindStringSubmatch(msg); len(matches) > 2 {
		for _, fmt := range initialize.DateFormats {
			t, err := time.Parse(fmt, matches[2])
			if err == nil {
				transaction.Date = t.Format(time.RFC3339)
				break
			}
		}
	}

	if matches := initialize.AmountRegex.FindStringSubmatch(msg); len(matches) > 1 {
		transaction.Amount = strings.TrimSpace(matches[2])
	}

	if matches := initialize.CurrencyRegex.FindStringSubmatch(msg); len(matches) > 1 {
		key := strings.TrimSpace(matches[1])
		if code, ok := initialize.CurrencyMap[key]; ok {
			transaction.Currency = code
		}
	}

	if matches := initialize.VendorRegex.FindStringSubmatch(msg); len(matches) > 1 {
		vendor := strings.TrimSpace(matches[1])
		if transaction.Type == "withdrawal" {
			transaction.DestVendor = &vendor
		} else if transaction.Type == "deposit" {
			transaction.SourceVendor = &vendor
		}
	}

	return transaction
}
