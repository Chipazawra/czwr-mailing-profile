package usecases

import (
	"context"
	"fmt"

	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
)

type Receivers struct {
	storage model.IReceiverStorage
}

func NewReceivers(ir model.IReceiverStorage) *Receivers {
	return &Receivers{
		storage: ir,
	}
}

func (r *Receivers) Create(ctx context.Context, user, name string) (*model.Receiver, error) {

	new := &model.Receiver{
		ID:   "",
		User: user,
		Name: name,
	}

	var err error

	new.ID, err = r.storage.Create(ctx, new)

	if err != nil {
		return nil, fmt.Errorf("usecases Receivers:%v", err)
	}

	return new, nil
}

func (r *Receivers) Read(ctx context.Context, user string) ([]*model.Receiver, error) {
	receivers, err := r.storage.Read(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("usecases Receivers:%v", err)
	}
	return receivers, nil
}

func (r *Receivers) Update(ctx context.Context, id, user, name string) (*model.Receiver, error) {

	upd := &model.Receiver{
		ID:   id,
		User: user,
		Name: name,
	}

	err := r.storage.Update(ctx, upd)

	if err != nil {
		return nil, fmt.Errorf("usecases Receivers:%v", err)
	}

	return upd, nil
}

func (r *Receivers) Delete(ctx context.Context, id string) error {

	err := r.storage.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("usecases Receivers:%v", err)
	}
	return nil
}
