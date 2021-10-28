package profile

import (
	"context"

	"github.com/Chipazawra/czwr-mailing-profile/internal/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Templates struct {
	mClient *mongo.Client
}

func New(mClient *mongo.Client) *Templates {
	return &Templates{mClient: mClient}
}

func (t *Templates) TemplateCreate(ctx context.Context, template entities.Template) (string, error) {
	mCollection := t.mClient.Database("profile").Collection("templates")
	res, err := mCollection.InsertOne(ctx,
		template,
	)
	if err != nil {
		panic(err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
