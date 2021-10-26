package profile

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	receivers *ReceiversStorage
	template  *TemplatesStorage
}

func New(dbCtx IDBctx) *Profile {
	return &Profile{
		&ReceiversStorage{dbCtx: dbCtx},
		&TemplatesStorage{dbCtx: dbCtx},
	}
}

func (p *Profile) Register(g *gin.Engine) *gin.RouterGroup {

	profile := g.Group("/profile")
	profile.GET("/i", p.iHandler)
	profile.GET("/me", p.meHandler)

	reciviers := profile.Group("/reciviers")
	reciviers.POST("/:usr/:receiver", p.CreateHandler)
	reciviers.GET("/:usr", p.ReadHandler)
	reciviers.PATCH("/:usr/:id/:receiver", p.UpdateHandler)
	reciviers.DELETE("/:usr/:id", p.DeleteHandler)
	reciviers.POST("/upload_template", p.UploadTemplateHandler)

	return profile
}

type IDBctx interface {
	ReceiverCreate(ctx context.Context, usr, receiver string) (string, error)
	ReceiverRead(ctx context.Context, usr string) ([]string, error)
	ReceiverUpdate(ctx context.Context, usr string, id string, receiver string) error
	ReceiverDelete(ctx context.Context, usr string, id string) error
	TemplateCreate(ctx context.Context, raw string, params []string) error
}

type ReceiversStorage struct {
	dbCtx IDBctx
}

func (r *ReceiversStorage) Create(ctx context.Context, usr, receiver string) (string, error) {

	id, err := r.dbCtx.ReceiverCreate(ctx, usr, receiver)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *ReceiversStorage) Read(ctx context.Context, usr string) ([]string, error) {

	lst, err := r.dbCtx.ReceiverRead(ctx, usr)
	if err != nil {
		return nil, err
	}
	return lst, nil
}

func (r *ReceiversStorage) Update(ctx context.Context, usr string, id string, receiver string) error {

	err := r.dbCtx.ReceiverUpdate(ctx, usr, id, receiver)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReceiversStorage) Delete(ctx context.Context, usr string, id string) error {

	err := r.dbCtx.ReceiverDelete(ctx, usr, id)
	if err != nil {
		return err
	}
	return nil
}

type TemplatesStorage struct {
	dbCtx IDBctx
}

func (t *TemplatesStorage) Create(ctx context.Context, raw string, params []string) {
	t.dbCtx.TemplateCreate(ctx, raw, params)
}
