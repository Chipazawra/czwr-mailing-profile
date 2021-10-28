package usecases

import (
	"context"
	"os"
	"testing"

	mongodriver "github.com/Chipazawra/czwr-mailing-profile/internal/profile/drivers/mongo"
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
