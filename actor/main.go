package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/actor"
	daprd "github.com/dapr/go-sdk/service/http"

	"github.com/kzmake/dapr-actor/api"
)

type PiggyBankActor struct {
	actor.ServerImplBase
}

func NewPiggyBankActor() func() actor.Server {
	return func() actor.Server {
		return &PiggyBankActor{}
	}
}

func (a *PiggyBankActor) Type() string {
	return "PiggyBank"
}

func (a *PiggyBankActor) Drop(ctx context.Context, coin api.Coin) error {
	log.Println("Actor: ", a.Type(), "/", a.ID(), " drop a coin: ", coin)
	pg, err := a.get()
	if err != nil {
		return err
	}

	pg.Coins = append(pg.Coins, coin)

	err = a.set(pg)
	if err != nil {
		return err
	}

	return nil
}
func (a *PiggyBankActor) Return(context.Context) ([]api.Coin, error) {
	log.Println("Actor: ", a.Type(), "/", a.ID(), " return coins")
	pg, err := a.get()
	if err != nil {
		return nil, err
	}

	new := &api.PiggyBank{
		ID:    a.ID(),
		Coins: []api.Coin{},
	}
	if err := a.set(new); err != nil {
		return nil, err
	}

	return pg.Coins, nil
}
func (a *PiggyBankActor) Get(context.Context) (*api.PiggyBank, error) {
	log.Println("Actor: ", a.Type(), "/", a.ID(), " get a piggy bank")
	pg, err := a.get()
	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (a *PiggyBankActor) get() (*api.PiggyBank, error) {
	pg := &api.PiggyBank{
		ID:    a.ID(),
		Coins: []api.Coin{},
	}

	if found, err := a.GetStateManager().Contains("piggy-bank"); err != nil {
		return nil, err
	} else if found {
		if err := a.GetStateManager().Get("piggy-bank", pg); err != nil {
			return nil, err
		}
	}

	return pg, nil
}

func (a *PiggyBankActor) set(pg *api.PiggyBank) error {
	if err := a.GetStateManager().Set("piggy-bank", pg); err != nil {
		return err
	}

	return nil
}

func main() {
	s := daprd.NewService(":8080")
	s.RegisterActorImplFactory(NewPiggyBankActor())

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}
