package transaction

import (
	"time"

	"github.com/nerlihmax/te/pkg/account"
)

type Transaction struct {
	ID          int32
	Debit       account.Account
	Credit      account.Account
	Amount      int64 // TODO: bigint
	Description string
	Timestamp   time.Time
}
