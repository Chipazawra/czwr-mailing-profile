package model

import (
	"context"
)

type Template struct {
	ID     string
	Raw    string
	Params []string
}

type ITemplatesUserCase interface {
	UploadTemplate(ctx context.Context, raw string) (*Template, error)
}

type ITemplatesStorage interface {
	Create(ctx context.Context, template *Template) (string, error)
}
