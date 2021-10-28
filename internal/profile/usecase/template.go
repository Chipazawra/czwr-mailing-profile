package usecases

import (
	"context"

	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
)

type Templates struct {
	templates *model.ITemplatesStorage
}

func NewTemplates(it *model.ITemplatesStorage) *Templates {
	return &Templates{
		templates: it,
	}
}

func (p *Templates) UploadTemplate(ctx context.Context, raw string) (*model.Template, error) {
	return nil, nil
}
