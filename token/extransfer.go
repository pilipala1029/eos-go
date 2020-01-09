package token

import eos "github.com/eoscanada/eos-go"

func NewExTransfer(from, to eos.AccountName, quantity eos.ExtendedAsset, memo string) *eos.Action {
	return &eos.Action{
		Account: AN("eosio.token"),
		Name:    ActN("extransfer"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(ExTransfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `eosio.token` contract.
type ExTransfer struct {
	From     eos.AccountName   `json:"from"`
	To       eos.AccountName   `json:"to"`
	Quantity eos.ExtendedAsset `json:"quantity"`
	Memo     string            `json:"memo"`
}
