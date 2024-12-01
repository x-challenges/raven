package main

import (
	"context"
	"log"

	"github.com/x-challenges/raven/stateless"
)

type Status = string

const (
	New    Status = "NEW"
	Closed Status = "CLOSED"
)

type Order struct {
	Status Status
}

var (
	Open  stateless.Trigger = "Open"
	Close stateless.Trigger = "Close"
)

func main() {
	var order = Order{
		Status: New,
	}

	var fsm = stateless.New[*Order](stateless.State(New))

	fsm.
		Configure(stateless.State(New)).
		Permit(stateless.Trigger(Close), stateless.State(Closed)).
		OnEntry(func(ctx context.Context, args *Order) error { args.Status = Closed; return nil })

	fsm.
		Configure(stateless.State(Closed)).
		Permit(stateless.Trigger(Open), stateless.State(New)).
		OnEntry(func(ctx context.Context, args *Order) error { args.Status = Closed; return nil })

	log.Print(order.Status)

	log.Println(fsm.Fire(context.TODO(), stateless.Trigger(Close), &order))

	log.Print(order.Status)
}
