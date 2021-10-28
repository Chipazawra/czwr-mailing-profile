package datastore

import (
	"context"

	"github.com/Chipazawra/czwr-mailing-profile/internal/entities"
)

type Receivers interface {
	Create(ctx context.Context, receiver entities.Receiver) (entities.Receiver, error)
	Read(ctx context.Context, usr string) ([]entities.Receiver, error)
	Update(ctx context.Context, id string, receiver entities.Receiver) (entities.Receiver, error)
	Delete(ctx context.Context, id string) error
}

type Templates interface {
	Create(ctx context.Context, template entities.Template) (entities.Template, error)
}
