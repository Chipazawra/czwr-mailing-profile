package profile

import (
	"context"

	"github.com/Chipazawra/czwr-mailing-profile/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Receivers struct {
	mClient *mongo.Client
}

func New(mClient *mongo.Client) *Receivers {
	return &Receivers{mClient: mClient}
}
func (r *Receivers) Create(ctx context.Context, receiver entities.Receiver) (entities.Receiver, error) {
	mCollection := r.mClient.Database("profile").Collection("receivers")
	_, err := mCollection.InsertOne(ctx,
		receiver,
	)
	if err != nil {
		panic(err)
	}

	return receiver, nil
}
func (r *Receivers) Read(ctx context.Context, usr string) ([]string, error) {
	mCollection := r.mClient.Database("profile").Collection("receivers")
	res, err := mCollection.Find(ctx,
		bson.M{"user": usr},
		options.Find().SetProjection(
			bson.D{
				primitive.E{Key: "name", Value: 1},
				primitive.E{Key: "_id", Value: 1},
			},
		),
	)
	if err != nil {
		panic(err)
	}

	var result []string
	for res.Next(ctx) {
		result = append(result, res.Current.String())
	}

	return result, nil
}
func (r *Receivers) Update(ctx context.Context, id string, receiver entities.Receiver) error {
	mCollection := r.mClient.Database("profile").Collection("receivers")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = mCollection.UpdateOne(ctx,
		bson.D{
			primitive.E{Key: "_id", Value: objectID},
		},
		bson.D{
			primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "name", Value: receiver},
			}},
		},
		options.Update().SetUpsert(false),
	)

	return err
}
func (r *Receivers) Delete(ctx context.Context, id string) error {

	mCollection := r.mClient.Database("profile").Collection("receivers")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = mCollection.DeleteOne(ctx,
		bson.D{
			primitive.E{Key: "_id", Value: objectID},
		},
	)

	return err
}
