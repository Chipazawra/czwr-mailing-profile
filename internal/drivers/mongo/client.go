package mongodriver

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Driver struct {
	client *mongo.Client
}

func New() *Driver {
	return &Driver{}
}

func (dr *Driver) Client() *mongo.Client {
	return dr.client
}

func (dr *Driver) Connect(ctx context.Context, user, pass, clst string) error {
	var err error
	cOpts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%v:%v@%v/receivers?retryWrites=true&w=majority", user, pass, clst),
	)
	dr.client, err = mongo.Connect(ctx, cOpts)
	return err
}

func (dr *Driver) Disonnect(ctx context.Context) {
	if err := dr.client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
