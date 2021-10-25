package inmemoryctx

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	ctx := New()

	if ctx == nil {
		t.Fatalf("ctx == nil")
	}

}

func TestCreate(t *testing.T) {

	ctx := New()
	_, err := ctx.Create("foo", "bar")
	if err != nil {
		t.Fatalf("Expected nil value in err, val = %v", err)
	}

}

func TestDelete(t *testing.T) {

	ctx := New()

	idx, err := ctx.Create("foo", "bar")
	if err != nil {
		t.Fatal(err)
	}

	err = ctx.Delete("foo", idx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {

	ctx := New()
	idx, err := ctx.Create("foo", "bar")
	if err != nil {
		t.Fatal(err)
	}

	err = ctx.Update("foo", idx, "bar1")
	if err != nil {
		t.Fatal(err)
	}

}

func TestRead(t *testing.T) {

	ctx := New()
	_, err := ctx.Create("foo", "bar")

	if err != nil {
		t.Fatal(err)
	}

	receiversList, _ := ctx.Read("foo")

	fmt.Printf("%v\n", receiversList)

}
