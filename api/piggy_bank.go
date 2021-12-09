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
	Yen500 Coin = 500
)

type PiggyBank struct {
	ID    string `json:"id"`
	Coins []Coin `json:"coins"`
}

type PiggyBankActor struct {
	Id string `json:"id"`

	Drop   func(context.Context, Coin) error
	Return func(context.Context) ([]Coin, error)
	Get    func(context.Context) (*PiggyBank, error)
}

func NewPiggyBankActor() *PiggyBankActor {
	return &PiggyBankActor{Id: ulid.MustNew(ulid.Timestamp(time.Now()), crand.Reader).String()}
}

func (a *PiggyBankActor) Type() string {
	return "PiggyBank"
}

func (a *PiggyBankActor) ID() string {
	return a.Id
}
