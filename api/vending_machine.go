package api

import (
	"context"
	crand "crypto/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

type Coin uint64

const (
	Yen10  Coin = 10
	Yen50  Coin = 50
	Yen100 Coin = 100
	Yen500 Coin = 100
)

type VendingMachine struct {
	ID       string `json:"id"`
	CoinSlot []Coin `json:"coin_slot"`
}

type VendingMachineActor struct {
	Drop   func(context.Context, Coin) error
	Return func(context.Context) ([]Coin, error)
	Get    func(context.Context) (*VendingMachine, error)
}

func (a *VendingMachineActor) Type() string {
	return "vendingmachine"
}

func (a *VendingMachineActor) ID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), crand.Reader).String()
}
