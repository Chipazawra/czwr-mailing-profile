package profile

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	mongoctx "github.com/Chipazawra/czwr-mailing-profile/internal/dbcontext/mongo"
)

var (
	service *Profile
)

func TestMain(m *testing.M) {
	//before test
	ctx := context.TODO()
	mClient := mongoctx.New()
	err := mClient.Connect(ctx, "admin", "admin")
	defer mClient.Disonnect(ctx)
	if err != nil {
		panic(err)
	}
	service = New(mClient)
	//run test
	exitVal := m.Run()
	//after test
	os.Exit(exitVal)
}

func TestNew(t *testing.T) {

}

func TestReceiversCreate(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*100000))

	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
			cancel()
		}
	}()
	for i := 0; i < 1000; i++ {
		if _, err := service.receivers.Create(ctx, fmt.Sprintf("user - %v", i), fmt.Sprintf("receiver - %v", i)); err != nil {
			panic(err)
		}
	}
}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
