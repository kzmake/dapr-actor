package main

import (
	"context"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"

	"github.com/kzmake/dapr-actor/api"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	actor := api.NewPiggyBankActor()
	client.ImplActorClientStub(actor)

	ctx := context.Background()

	actor.Drop(ctx, api.Yen10)
	actor.Drop(ctx, api.Yen10)
	actor.Drop(ctx, api.Yen100)
	pg, _ := actor.Get(ctx)
	fmt.Println("get a piggy bank: ", pg)

	actor.Drop(ctx, api.Yen500)
	pg, _ = actor.Get(ctx)
	fmt.Println("get a piggy bank: ", pg)

	coins, _ := actor.Return(ctx)
	fmt.Println("return conins: ", coins)
	pg, _ = actor.Get(ctx)
	fmt.Println("get a piggy bank: ", pg)

	actor.Drop(ctx, api.Yen10)
	pg, _ = actor.Get(ctx)
	fmt.Println("get a piggy bank: ", pg)
}
