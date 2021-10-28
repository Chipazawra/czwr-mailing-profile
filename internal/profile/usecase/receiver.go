package usecases

import (
	"context"

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
		return nil, err
	}

	return new, err
}
func (r *Receivers) Read(ctx context.Context, user string) ([]*model.Receiver, error) {
	return nil, nil
}
func (r *Receivers) Update(ctx context.Context, id, user, name string) (*model.Receiver, error) {
	return nil, nil
}
func (r *Receivers) Delete(ctx context.Context, id string) error {
	return nil
}
