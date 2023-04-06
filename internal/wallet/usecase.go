package wallet

import "context"

type InitWithdrawUseCase interface {
	Init(context.Context, InitWithdrawInput) error
}

type InitWithdrawInput struct{}
