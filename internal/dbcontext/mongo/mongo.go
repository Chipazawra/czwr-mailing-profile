package mongoctx

import (
	"context"
	"fmt"

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

func (mc *MongoClinet) Connect(ctx context.Context, user, pass string) error {
	var err error
	cOpts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%v:%v@czwrmongo.yrzjn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", user, pass),
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

	res, err := mCollection.InsertOne(ctx, Receiver{
		ID:   [12]byte{},
		User: usr,
		Name: receiver,
	})
	if err != nil {
		panic(err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (mc *MongoClinet) ReceiverRead(ctx context.Context, usr string) ([]string, error) {
	return nil, nil
}
func (mc *MongoClinet) ReceiverUpdate(ctx context.Context, usr string, id string, receiver string) error {
	return nil
}
func (mc *MongoClinet) ReceiverDelete(ctx context.Context, usr string, id string) error {
	return nil
}
func (mc *MongoClinet) TemplateCreate(ctx context.Context, raw string, params []string) error {
	return nil
}
