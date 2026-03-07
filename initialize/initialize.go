package initialize

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var ListenAddress string
var FireflyAddress string
var FireflyPAT string
var AuthKey string
var WithdrawalKeys []string
var DepositKeys []string
var VendorKeys []string
var DateKeys []string
var AmountKeys []string
var CurrencyMap map[string]string
var DateFormats []string

var WithdrawRegex *regexp.Regexp
var DepositRegex *regexp.Regexp
var CurrencyRegex *regexp.Regexp
var AmountRegex *regexp.Regexp
var DateRegex *regexp.Regexp
var VendorRegex *regexp.Regexp

func InitEnv() {
	ListenAddress = os.Getenv("LISTEN_ADDRESS")
	FireflyAddress = os.Getenv("FIREFLY_ADDRESS")
	FireflyPAT = os.Getenv("FIREFLY_PAT")
	AuthKey = os.Getenv("AUTH_KEY")
	WithdrawalKeys = strings.Split(os.Getenv("WITHDRAWAL_KEYS"), ",")
	DepositKeys = strings.Split(os.Getenv("DEPOSIT_KEYS"), ",")
	VendorKeys = strings.Split(os.Getenv("VENDOR_KEYS"), ",")
	DateKeys = strings.Split(os.Getenv("DATE_KEYS"), ",")
	AmountKeys = strings.Split(os.Getenv("AMOUNT_KEYS"), ",")
	DateFormats = strings.Split(os.Getenv("DATE_FORMATS"), ",")

	currencyMap := os.Getenv("CURRENCY_MAP")
	if currencyMap != "" {
		CurrencyMap = make(map[string]string)
		for _, pair := range strings.Split(currencyMap, ",") {
			kv := strings.SplitN(strings.TrimSpace(pair), ":", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				val := strings.TrimSpace(kv[1])
				CurrencyMap[key] = val
			}
		}
	}
}

func InitRegex() {
	withdrawPattern := strings.Join(WithdrawalKeys, "|")
	depositPattern := strings.Join(DepositKeys, "|")
	amountPattern := strings.Join(AmountKeys, "|")
	datePattern := strings.Join(DateKeys, "|")
	vendorPattern := strings.Join(VendorKeys, "|")

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
