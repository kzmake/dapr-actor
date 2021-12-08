package main

import (
	"context"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"

	"github.com/kzmake/dapr-actor/api"
)

func main() {
	ctx := context.Background()

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	actor := new(api.VendingMachineActor)
	client.ImplActorClientStub(actor)

	actor.Drop(ctx, api.Yen10)
	actor.Drop(ctx, api.Yen10)
	actor.Drop(ctx, api.Yen100)
	vm, _ := actor.Get(ctx)
	fmt.Println("get a vending machine: ", vm)

	actor.Drop(ctx, api.Yen500)
	vm, _ = actor.Get(ctx)
	fmt.Println("get a vending machine: ", vm)

	coins, _ := actor.Return(ctx)
	fmt.Println("return conins: ", coins)

	vm, _ = actor.Get(ctx)
	fmt.Println("get a vending machine: ", vm)
}
