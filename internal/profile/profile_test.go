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
	err := mClient.Connect(ctx, "admin", "admin", "czwrmongo.yrzjn.mongodb.net")
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

func TestReceiversCreate(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*100))

	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
		}
		cancel()
	}()
	for i := 0; i < 1000; i++ {
		if _, err := service.receivers.Create(ctx, fmt.Sprintf("user"), fmt.Sprintf("receiver - %v", i)); err != nil {
			panic(err)
		}
	}
}

func TestReceiversRead(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))

	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
		}
		cancel()
	}()

	if receivers, err := service.receivers.Read(ctx, "user"); err != nil {
		panic(err)
	} else {
		fmt.Printf("%v\n", receivers)
	}
}

func TestReceiversUpdate(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))

	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
		}
		cancel()
	}()

	if err := service.receivers.Update(ctx, "61783fc030ef8ccf833d7a0d", "noname"); err != nil {
		panic(err)
	}
}

func TestReceiversDelete(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))

	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
		}
		cancel()
	}()

	if err := service.receivers.Delete(ctx, "user", "61783fc030ef8ccf833d7a0d"); err != nil {
		panic(err)
	}
}

func TestTemplateCreate(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))

	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
		}
		cancel()
	}()

	if _, err := service.template.Create(ctx, "<div><h1>{{ .Title}}</h1><p>{{ .Message}}</p></div>"); err != nil {
		panic(err)
	}

	if _, err := service.template.Create(ctx, "<div><h1></h1><p></p></div>"); err != nil {
		panic(err)
	}

}
