package initialize

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var ListenAddress string
var FireflyAddress string
var AuthKey string
var WithdrawalKeys []string
var DepositKeys []string
var VendorKeys []string
var DateKeys []string
var AmountKeys []string
var CurrencyKeys []string

var WithdrawRegex *regexp.Regexp
var DepositRegex *regexp.Regexp
var CurrencyRegex *regexp.Regexp
var AmountRegex *regexp.Regexp
var DateRegex *regexp.Regexp
var VendorRegex *regexp.Regexp

func InitEnv() {
	ListenAddress = os.Getenv("LISTEN_ADDRESS")
	FireflyAddress = os.Getenv("FIREFLY_ADDRESS")
	AuthKey = os.Getenv("AUTH_KEY")
	WithdrawalKeys = strings.Split(os.Getenv("WITHDRAWAL_KEYS"), ",")
	DepositKeys = strings.Split(os.Getenv("DEPOSIT_KEYS"), ",")
	CurrencyKeys = strings.Split(os.Getenv("CURRENCY_KEYS"), ",")
	VendorKeys = strings.Split(os.Getenv("VENDOR_KEYS"), ",")
	DateKeys = strings.Split(os.Getenv("DATE_KEYS"), ",")
	AmountKeys = strings.Split(os.Getenv("AMOUNT_KEYS"), ",")
}

func InitRegex() {
	withdrawPattern := strings.Join(WithdrawalKeys, "|")
	depositPattern := strings.Join(DepositKeys, "|")
	currencyPattern := strings.Join(CurrencyKeys, "|")
	amountPattern := strings.Join(AmountKeys, "|")
	datePattern := strings.Join(DateKeys, "|")
	vendorPattern := strings.Join(VendorKeys, "|")

	WithdrawRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, withdrawPattern))
	DepositRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, depositPattern))
	CurrencyRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, currencyPattern))
	AmountRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)[:\s]*(\d+(\.\d{2})?)`, amountPattern))
	DateRegex = regexp.MustCompile(fmt.Sprintf(`(?i)(%s)[:\s]*([\d/\-:\s]+(?:AM|PM)?)`, datePattern))
	VendorRegex = regexp.MustCompile(
		fmt.Sprintf(`(?i)%s[:\s]*([^\n]+?)(?:\s+(?:%s|%s|%s|%s|%s)|$)`,
			vendorPattern,
			amountPattern,
			datePattern,
			currencyPattern,
			withdrawPattern,
			depositPattern,
		),
	)
}
