package usecases

import (
	"context"
	"fmt"
	"regexp"
	"text/template"

	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
)

type Templates struct {
	templates model.ITemplatesStorage
}

func NewTemplates(it model.ITemplatesStorage) *Templates {
	return &Templates{
		templates: it,
	}
}

func (t *Templates) UploadTemplate(ctx context.Context, raw string) (*model.Template, error) {

	tmpl := template.New("tmpl")
	_, err := tmpl.Parse(string(raw))

	if err != nil {
		return nil, fmt.Errorf("usecases Templates:%v", err)
	}

	rx, _ := regexp.Compile(`{{ *\.(.*?)}}`)
	prms := func(p []string) []string {
		for _, v := range rx.FindAllStringSubmatch(string(raw), -1) {
			p = append(p, v[1])
		}
		return p
	}(make([]string, 0))

	new := &model.Template{
		ID:     "",
		Raw:    raw,
		Params: prms,
	}

	new.ID, err = t.templates.Create(ctx, new)
	if err != nil {
		return nil, fmt.Errorf("usecases Templates:%v", err)
	}

	return new, nil

}
