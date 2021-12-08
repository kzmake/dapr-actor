package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/actor"
	daprd "github.com/dapr/go-sdk/service/http"

	"github.com/kzmake/dapr-actor/api"
)

type VendingMachineActor struct {
	actor.ServerImplBase
}

func NewVendingMachineActor() func() actor.Server {
	return func() actor.Server {
		return &VendingMachineActor{}
	}
}

func (a *VendingMachineActor) Type() string {
	return "vendingmachine"
}

func (a *VendingMachineActor) Drop(ctx context.Context, coin api.Coin) error {
	log.Println("drop coin: ", coin)
	vm, err := a.get()
	if err != nil {
		return err
	}
	log.Println("get vm: ", vm)
	vm.CoinSlot = append(vm.CoinSlot, coin)
	log.Println("append coin: ", vm.CoinSlot)
	err = a.set(vm)
	if err != nil {
		return err
	}

	log.Println("set vm: ", vm)
	return nil
}
func (a *VendingMachineActor) Return(context.Context) ([]api.Coin, error) {
	vm, err := a.get()
	if err != nil {
		return nil, err
	}

	new := &api.VendingMachine{
		ID:       a.ID(),
		CoinSlot: []api.Coin{},
	}
	if err := a.set(new); err != nil {
		return nil, err
	}

	return vm.CoinSlot, nil
}
func (a *VendingMachineActor) Get(context.Context) (*api.VendingMachine, error) {
	vm, err := a.get()
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func (a *VendingMachineActor) get() (*api.VendingMachine, error) {
	vm := &api.VendingMachine{
		ID:       a.ID(),
		CoinSlot: []api.Coin{},
	}

	if found, err := a.GetStateManager().Contains(a.ID()); err != nil {
		fmt.Println("state manager call contains with " + a.ID() + "err = " + err.Error())

		return nil, err
	} else if found {
		if err := a.GetStateManager().Get(a.ID(), vm); err != nil {
			fmt.Println("state manager call get with " + a.ID() + "err = " + err.Error())

			return nil, err
		}
	}

	return vm, nil
}

func (a *VendingMachineActor) set(vm *api.VendingMachine) error {
	if err := a.GetStateManager().Set(a.ID(), vm); err != nil {
		fmt.Println("state manager call save with " + a.ID() + "err = " + err.Error())

		return err
	}

	return nil
}

func main() {
	s := daprd.NewService(":8080")
	s.RegisterActorImplFactory(NewVendingMachineActor())

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}
