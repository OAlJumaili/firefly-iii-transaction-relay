package parser

import (
	"firefly-iii-transaction-relay/core"
	"fmt"
	"strings"
	"time"
)

type Transaction struct {
	Type         string  `json:"type"`
	Date         string  `json:"date"`
	Description  string  `json:"description"`
	Amount       string  `json:"amount"`
	Currency     string  `json:"currency_code"`
	Budget       string  `json:"budget_name"`
	SourceVendor *string `json:"source_name,omitempty"`
	DestVendor   *string `json:"destination_name,omitempty"`
}

func ParseMessage(msg string) (Transaction, bool) {
	var transaction Transaction

	if core.BlacklistRegex.MatchString(msg) {
		return transaction, true
	}

	transaction.Budget = core.FireflyBudget

	if core.WithdrawRegex.MatchString(msg) {
		transaction.Type = "withdrawal"
	} else if core.DepositRegex.MatchString(msg) {
		transaction.Type = "deposit"
	}

	if matches := core.DateRegex.FindStringSubmatch(msg); len(matches) > 2 {
		for _, fmt := range core.DateFormats {
			t, err := time.Parse(fmt, matches[2])
			if err == nil {
				transaction.Date = t.Format(time.RFC3339)
				break
			}
		}
	}

	if matches := core.AmountRegex.FindStringSubmatch(msg); len(matches) > 1 {
		transaction.Amount = strings.TrimSpace(matches[2])
	}

	if matches := core.CurrencyRegex.FindStringSubmatch(msg); len(matches) > 1 {
		key := strings.TrimSpace(matches[1])
		if code, ok := core.CurrencyMap[key]; ok {
			transaction.Currency = code
		}
	}

	if transaction.Type == "withdrawal" {
		if matches := core.WithdrawalVendorRegex.FindStringSubmatch(msg); len(matches) > 1 {
			vendor := strings.TrimSpace(matches[1])
			transaction.SourceVendor = &core.FireflyAccount
			transaction.DestVendor = &vendor
			transaction.Description = fmt.Sprintf("Automated Imported Transaction For %s, %s", *transaction.SourceVendor, transaction.Date)
		}
	} else if transaction.Type == "deposit" {
		if matches := core.DepositVendorRegex.FindStringSubmatch(msg); len(matches) > 1 {
			vendor := strings.TrimSpace(matches[1])
			transaction.DestVendor = &core.FireflyAccount
			transaction.SourceVendor = &vendor
			transaction.Description = fmt.Sprintf("Automated Imported Transaction For %s, %s", *transaction.DestVendor, transaction.Date)
		}
	}
	return transaction, false
}
