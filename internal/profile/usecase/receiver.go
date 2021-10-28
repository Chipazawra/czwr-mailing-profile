package usecases

import (
	"context"

	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
)

type Receivers struct {
	receivers *model.IReceiverStorage
}

func NewReceivers(ir *model.IReceiverStorage) *Receivers {
	return &Receivers{
		receivers: ir,
	}
}

func (p *Receivers) Create(ctx context.Context, user, name string) (*model.Receiver, error) {
	return nil, nil
}
func (p *Receivers) Read(ctx context.Context, user string) ([]model.Receiver, error) {
	return nil, nil
}
func (p *Receivers) Update(ctx context.Context, id, user, name string) (*model.Receiver, error) {
	return nil, nil
}
func (p *Receivers) Delete(ctx context.Context, id string) error {
	return nil
}
