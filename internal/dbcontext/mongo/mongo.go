package mongoctx

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClinet struct {
	client *mongo.Client
}

type Receiver struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	User string             `bson:"user"`
	Name string             `bson:"name"`
}

type Template struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Raw    string             `bson:"raw"`
	Params []string           `bson:"params"`
}

func New() *MongoClinet {
	return &MongoClinet{}
}

func (mc *MongoClinet) Connect(ctx context.Context, user, pass, clst string) error {
	var err error
	cOpts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%v:%v@%v/receivers?retryWrites=true&w=majority", user, pass, clst),
	)
	mc.client, err = mongo.Connect(ctx, cOpts)
	return err
}

func (mc *MongoClinet) Disonnect(ctx context.Context) {
	if err := mc.client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func (mc *MongoClinet) ReceiverCreate(ctx context.Context, usr, receiver string) (string, error) {
	mCollection := mc.client.Database("profile").Collection("receivers")
	res, err := mCollection.InsertOne(ctx,
		Receiver{
			ID:   [12]byte{},
			User: usr,
			Name: receiver,
		},
	)
	if err != nil {
		panic(err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (mc *MongoClinet) ReceiverRead(ctx context.Context, usr string) ([]string, error) {
	mCollection := mc.client.Database("profile").Collection("receivers")
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
func (mc *MongoClinet) ReceiverUpdate(ctx context.Context, id string, receiver string) error {
	mCollection := mc.client.Database("profile").Collection("receivers")
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
func (mc *MongoClinet) ReceiverDelete(ctx context.Context, usr string, id string) error {

	mCollection := mc.client.Database("profile").Collection("receivers")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = mCollection.DeleteOne(ctx,
		bson.D{
			primitive.E{Key: "_id", Value: objectID},
			primitive.E{Key: "user", Value: usr},
		},
	)

	return err

}
func (mc *MongoClinet) TemplateCreate(ctx context.Context, raw string, params []string) (string, error) {
	mCollection := mc.client.Database("profile").Collection("templates")
	res, err := mCollection.InsertOne(ctx,
		Template{
			ID:     [12]byte{},
			Raw:    raw,
			Params: params,
		},
	)
	if err != nil {
		panic(err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
