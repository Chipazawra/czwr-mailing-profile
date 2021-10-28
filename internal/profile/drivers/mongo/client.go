package mongoclient

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
}

func New() *Client {
	return &Client{}
}

func (mc *Client) Connect(ctx context.Context, user, pass, clst string) error {
	var err error
	cOpts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%v:%v@%v/receivers?retryWrites=true&w=majority", user, pass, clst),
	)
	mc.client, err = mongo.Connect(ctx, cOpts)
	return err
}

func (mc *Client) Disonnect(ctx context.Context) {
	if err := mc.client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
