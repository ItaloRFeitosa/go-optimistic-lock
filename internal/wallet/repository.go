package wallet

import "context"

type AccountRepository interface {
	Save(context.Context, *Account) error
}
