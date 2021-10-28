package mongostorage

import (
	"context"

	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Receivers struct {
	mClient *mongo.Client
}

type mReceiver struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	User string             `bson:"user"`
	Name string             `bson:"name"`
}

func NewReceivers(mClient *mongo.Client) *Receivers {
	return &Receivers{mClient: mClient}
}

//					Create(ctx context.Context, receiver *Receiver) (string, error)
func (r *Receivers) Create(ctx context.Context, receiver *model.Receiver) (string, error) {

	mCollection := r.mClient.Database("profile").Collection("receivers")
	res, err := mCollection.InsertOne(ctx,
		&mReceiver{
			ID:   [12]byte{},
			User: receiver.User,
			Name: receiver.Name,
		},
	)

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

//					Read(ctx context.Context, usr string) ([]*Receiver, error)
func (r *Receivers) Read(ctx context.Context, usr string) ([]*model.Receiver, error) {
	mCollection := r.mClient.Database("profile").Collection("receivers")
	res, err := mCollection.Find(ctx,
		bson.M{"user": usr},
		options.Find().SetProjection(
			bson.D{
				primitive.E{Key: "_id", Value: 1},
				primitive.E{Key: "name", Value: 1},
				primitive.E{Key: "user", Value: 1},
			},
		),
	)
	if err != nil {
		panic(err)
	}

	var result []*model.Receiver
	for res.Next(ctx) {
		var mReceiver mReceiver
		err = res.Decode(&mReceiver)
		if err != nil {
			return nil, err
		}
		result = append(result, func() *model.Receiver {
			return &model.Receiver{
				ID:   mReceiver.ID.Hex(),
				User: mReceiver.User,
				Name: mReceiver.Name,
			}
		}())
	}

	return result, nil
}

//					Update(ctx context.Context, receiver *Receiver) error
func (r *Receivers) Update(ctx context.Context, receiver *model.Receiver) error {
	mCollection := r.mClient.Database("profile").Collection("receivers")
	objectID, err := primitive.ObjectIDFromHex(receiver.ID)
	if err != nil {
		return err
	}

	_, err = mCollection.UpdateOne(ctx,
		bson.D{
			primitive.E{Key: "_id", Value: objectID},
		},
		bson.D{
			primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "name", Value: receiver.Name},
				primitive.E{Key: "user", Value: receiver.User},
			}},
		},
		options.Update().SetUpsert(false),
	)

	return err
}

//					Delete(ctx context.Context, id string) error
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
