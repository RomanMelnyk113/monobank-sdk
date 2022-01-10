package monobank

type CardType string

const (
	Black    CardType = "black"
	White    CardType = "white"
	Platinum CardType = "platinum"
	FOP      CardType = "fop"
)

type UserInfo struct {
	ID         string   `json:"clientId"`
	Name       string   `json:"name"`
	WebHookURL string   `json:"webHookUrl"`
	Accounts   Accounts `json:"accounts"`
}

type Account struct {
	AccountID    string   `json:"id"`
	Balance      int64    `json:"balance"`
	CreditLimit  int64    `json:"creditLimit"`
	CurrencyCode int      `json:"currencyCode"`
	CashbackType string   `json:"cashbackType"` // enum: None, UAH, Miles
	CardMasks    []string `json:"maskedPan"`    // card number masks
	Type         CardType `json:"type"`
	IBAN         string   `json:"iban"`
}

type Accounts []Account

// Transaction - bank account statement
type Transaction struct {
	ID              string `json:"id"`
	Time            int64  `json:"time"`
	Description     string `json:"description"`
	MCC             int32  `json:"mcc"`
	OriginalMCC     int32  `json:"originalMcc"`
	Hold            bool   `json:"hold"`
	Amount          int64  `json:"amount"`
	OperationAmount int64  `json:"operationAmount"`
	CurrencyCode    int32  `json:"currencyCode"`
	CommissionRate  int64  `json:"commissionRate"`
	CashbackAmount  int64  `json:"cashbackAmount"`
	Balance         int64  `json:"balance"`
	Comment         string `json:"comment"`
	ReceiptID       string `json:"receiptId"`
	EDRPOU          string `json:"counterEdrpou"`
	IBAN            string `json:"counterIban"`
}

// Transactions - transactions
type Transactions []Transaction

type Currency struct {
	CurrencyCodeA int     `json:"currencyCodeA"`
	CurrencyCodeB int     `json:"currencyCodeB"`
	Date          int64   `json:"date"`
	RateSell      float64 `json:"rateSell"`
	RateBuy       float64 `json:"rateBuy"`
	RateCross     float64 `json:"rateCross"`
}

type Currencies []Currency

type Error struct {
	ErrorDescription string `json:"errorDescription"`
}

func (e Error) Error() string {
	return e.ErrorDescription
}
