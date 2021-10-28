package model

import (
	"context"
)

type Receiver struct {
	ID   string
	User string
	Name string
}

type IReceiverUserCase interface {
	Create(ctx context.Context, user, name string) (*Receiver, error)
	Read(ctx context.Context, user string) ([]Receiver, error)
	Update(ctx context.Context, id, user, name string) (*Receiver, error)
	Delete(ctx context.Context, id string) error
}

type IReceiverStorage interface {
	Create(ctx context.Context, receiver *Receiver) (string, error)
	Read(ctx context.Context, usr string) ([]Receiver, error)
	Update(ctx context.Context, receiver *Receiver) error
	Delete(ctx context.Context, id string) error
}
