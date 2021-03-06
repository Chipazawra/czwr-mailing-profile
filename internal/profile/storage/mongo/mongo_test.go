package mongostorage

import (
	"context"
	"log"
	"os"
	"testing"

	mongodriver "github.com/Chipazawra/czwr-mailing-profile/internal/drivers/mongo"
	"github.com/Chipazawra/czwr-mailing-profile/internal/profile/model"
)

var (
	rStorage *Receivers
	tStorage *Templates
)

func TestMain(m *testing.M) {
	//before test
	ctx := context.TODO()

	mDriver := mongodriver.New()
	err := mDriver.Connect(ctx, "admin", "admin", "czwrmongo.yrzjn.mongodb.net")
	defer mDriver.Disonnect(ctx)

	if err != nil {
		panic(err)
	}

	rStorage = NewReceivers(mDriver.Client())
	tStorage = NewTemplates(mDriver.Client())
	//run test
	exitVal := m.Run()
	//after test
	os.Exit(exitVal)
}

func TestTemplatesCreate(t *testing.T) {

	ctx := context.TODO()

	tmpl := &model.Template{
		ID:     "",
		Raw:    "<title>{{.Title}}</title>",
		Params: []string{"Title"},
	}

	if _, err := tStorage.Create(ctx, tmpl); err != nil {
		t.Errorf("Err = %v", err)
	}

}
func TestReceiverCreate(t *testing.T) {

	ctx := context.TODO()

	rcvr := &model.Receiver{
		ID:   "",
		User: "usr",
		Name: "TestReceiverCreate_receiver",
	}

	if id, err := rStorage.Create(ctx, rcvr); err != nil {
		t.Errorf("Err = %v", err)
	} else {
		log.Printf("id = %v\n", id)
	}

}
func TestReceiverRead(t *testing.T) {

	ctx := context.TODO()
	rcvr := &model.Receiver{
		User: "usr",
		Name: "TestReceiverRead_receiver",
	}

	_, err := rStorage.Create(ctx, rcvr)
	if err != nil {
		t.Errorf("Err = %v", err)
	}

	_, err = rStorage.Read(ctx, rcvr.User)
	if err != nil {
		t.Errorf("err = %v", err)
	}

}
func TestReceiverUpdate(t *testing.T) {

	ctx := context.TODO()
	rcvr := &model.Receiver{
		User: "usr",
		Name: "TestReceiverUpdate_receiver",
	}

	id, err := rStorage.Create(ctx, rcvr)
	if err != nil {
		panic(err)
	}

	rcvrupd := &model.Receiver{
		ID:   id,
		User: "usr",
		Name: "TestReceiverUpdate_receiver_TestReceiverUpdate_receiver",
	}

	err = rStorage.Update(ctx, rcvrupd)
	if err != nil {
		t.Errorf("err = %v", err)
	}

}
func TestReceiverDelete(t *testing.T) {

	ctx := context.TODO()
	rcvr := &model.Receiver{
		User: "usr",
		Name: "TestReceiverDelete_receiver",
	}

	id, err := rStorage.Create(ctx, rcvr)
	if err != nil {
		t.Errorf("Err = %v", err)
	}

	err = rStorage.Delete(ctx, id)
	if err != nil {
		t.Errorf("err = %v", err)
	}

}
