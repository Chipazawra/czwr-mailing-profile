package usecases

import (
	"context"
	"log"
	"os"
	"testing"

	mongodriver "github.com/Chipazawra/czwr-mailing-profile/internal/drivers/mongo"
	mongostorage "github.com/Chipazawra/czwr-mailing-profile/internal/profile/storage/mongo"
)

var (
	uReceivers *Receivers
	uTemplates *Templates
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

	rStorage := mongostorage.NewReceivers(mDriver.Client())
	tStorage := mongostorage.NewTemplates(mDriver.Client())

	uReceivers = NewReceivers(rStorage)
	uTemplates = NewTemplates(tStorage)

	//run test
	exitVal := m.Run()
	//after test
	os.Exit(exitVal)
}

func TestTemplatesCreate(t *testing.T) {

	ctx := context.TODO()
	template, err := uTemplates.UploadTemplate(ctx, "<title>{{.Title}}</title>")

	if err != nil {
		t.Errorf("Err = %v", err)
	} else {
		log.Printf("Receiver = %v", template)
	}

}
func TestReceiverCreate(t *testing.T) {

	ctx := context.TODO()
	receiver, err := uReceivers.Create(ctx, "uscase_user", "usecase_reveiver")

	if err != nil {
		t.Errorf("Err = %v", err)
	} else {
		log.Printf("Receiver = %v", receiver)
	}

}

func TestReceiverRead(t *testing.T) {

	ctx := context.TODO()

	receiver, err := uReceivers.Create(ctx, "uscase_user", "usecase_reveiver")
	if err != nil {
		t.Errorf("Err = %v", err)
	}

	_, err = uReceivers.Read(ctx, receiver.User)
	if err != nil {
		t.Errorf("Err = %v", err)
	}

}

func TestReceiverUpdate(t *testing.T) {

	ctx := context.TODO()
	receiver, err := uReceivers.Create(ctx, "uscase_user", "usecase_reveiver")
	if err != nil {
		t.Errorf("Err = %v", err)
	} else {
		log.Printf("Receiver = %v", receiver)
	}

	receiver, err = uReceivers.Update(ctx, receiver.ID, "uscase_user_upd", "usecase_reveiver_upd")
	if err != nil {
		t.Errorf("err = %v", err)
	} else {
		log.Printf("Receiver = %v", receiver)
	}

}

func TestReceiverDelete(t *testing.T) {

	ctx := context.TODO()
	receiver, err := uReceivers.Create(ctx, "uscase_user", "usecase_reveiver_deleted")
	if err != nil {
		t.Errorf("Err = %v", err)
	} else {
		log.Printf("Receiver = %v", receiver)
	}

	err = uReceivers.Delete(ctx, receiver.ID)
	if err != nil {
		t.Errorf("err = %v", err)
	} else {
		log.Printf("Receiver = %v", receiver)
	}

}
