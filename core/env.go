package core

import (
	"os"
	"strings"
)

// Firefly / Networking Settings
var ListenAddress string
var FireflyAddress string
var FireflyPAT string
var FireflyAccount string
var FireflyBudget string
var AuthKey string
var VerificationKey string
var JWTSigningSecret string

// Regex Settings
var WithdrawalKeys []string
var DepositKeys []string
var WithdrawalVendorKeys []string
var DepositVendorKeys []string
var DateKeys []string
var AmountKeys []string
var DateFormats []string
var BlacklistKeys []string
var CurrencyMap map[string]string

func InitEnv() {
	// Firefly / Networking Settings
	ListenAddress = os.Getenv("LISTEN_ADDRESS")
	FireflyAddress = os.Getenv("FIREFLY_ADDRESS")
	FireflyPAT = os.Getenv("FIREFLY_PAT")
	FireflyAccount = os.Getenv("FIREFLY_ACCOUNT")
	FireflyBudget = os.Getenv("FIREFLY_BUDGET")
	AuthKey = os.Getenv("AUTH_KEY")
	VerificationKey = os.Getenv("VERIFICATION_KEY")
	JWTSigningSecret = os.Getenv("JWT_SIGNING_SECRET")

	// Regex Settings
	WithdrawalKeys = strings.Split(os.Getenv("WITHDRAWAL_KEYS"), ",")
	DepositKeys = strings.Split(os.Getenv("DEPOSIT_KEYS"), ",")
	WithdrawalVendorKeys = strings.Split(os.Getenv("WITHDRAWAL_VENDOR_KEYS"), ",")
	DepositVendorKeys = strings.Split(os.Getenv("DEPOSIT_VENDOR_KEYS"), ",")
	DateKeys = strings.Split(os.Getenv("DATE_KEYS"), ",")
	AmountKeys = strings.Split(os.Getenv("AMOUNT_KEYS"), ",")
	BlacklistKeys = strings.Split(os.Getenv("BLACKLIST_KEYS"), ",")
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
