package mongostorage

import (
	"context"

	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Templates struct {
	mClient *mongo.Client
}

func NewTemplates(mClient *mongo.Client) *Templates {
	return &Templates{mClient: mClient}
}

func (t *Templates) Create(ctx context.Context, template *model.Template) error {

	mCollection := t.mClient.Database("profile").Collection("templates")
	_, err := mCollection.InsertOne(ctx,
		struct {
			ID     primitive.ObjectID `bson:"_id,omitempty"`
			Raw    string             `bson:"raw"`
			Params []string           `bson:"params"`
		}{
			ID:     [12]byte{},
			Raw:    template.Raw,
			Params: template.Params,
		},
	)

	return err
}
