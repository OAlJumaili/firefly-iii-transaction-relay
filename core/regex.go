package core

import (
	"fmt"
	"regexp"
	"strings"
)

var WithdrawRegex *regexp.Regexp
var DepositRegex *regexp.Regexp
var CurrencyRegex *regexp.Regexp
var AmountRegex *regexp.Regexp
var DateRegex *regexp.Regexp
var WithdrawalVendorRegex *regexp.Regexp
var DepositVendorRegex *regexp.Regexp
var BlacklistRegex *regexp.Regexp

func InitRegex() {
	withdrawPattern := strings.Join(WithdrawalKeys, "|")
	depositPattern := strings.Join(DepositKeys, "|")
	amountPattern := strings.Join(AmountKeys, "|")
	datePattern := strings.Join(DateKeys, "|")
	withdrawalVendorPattern := strings.Join(WithdrawalVendorKeys, "|")
	depositVendorPattern := strings.Join(DepositVendorKeys, "|")
	blacklistPattern := strings.Join(BlacklistKeys, "|")

	keys := make([]string, 0, len(CurrencyMap))
	for k := range CurrencyMap {
		keys = append(keys, regexp.QuoteMeta(k))
	}
	currencyPattern := strings.Join(keys, "|")

	WithdrawRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, withdrawPattern))
	DepositRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, depositPattern))
	CurrencyRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, currencyPattern))
	AmountRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)[:\s]*(\d+(?:\.\d+)?)`, amountPattern))
	DateRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)[:\s]*([\d/\-:\s]+(?:AM|PM)?)`, datePattern))
	BlacklistRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, blacklistPattern))

	WithdrawalVendorRegex = regexp.MustCompile(
		fmt.Sprintf(`(?i)%s[:\s]*([^\n]+?)(?:\s+(?:%s|%s|%s|%s|%s|%s)|$)`,
			withdrawalVendorPattern,
			amountPattern,
			datePattern,
			currencyPattern,
			withdrawPattern,
			depositPattern,
		),
	)

	DepositVendorRegex = regexp.MustCompile(
		fmt.Sprintf(`(?i)%s[:\s]*([^\n]+?)(?:\s+(?:%s|%s|%s|%s|%s|%s)|$)`,
			depositVendorPattern,
			amountPattern,
			datePattern,
			currencyPattern,
			withdrawPattern,
			depositPattern,
		),
	)
}
