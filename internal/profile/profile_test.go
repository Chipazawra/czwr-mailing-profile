package profile

import (
	"testing"

	"github.com/Chipazawra/czwr-mailing-profile/internal/dbcontext/inmemoryctx"
)

func TestNew(t *testing.T) {

	r := New(inmemoryctx.New())

	if r == nil {
		t.Fatalf("r == nil")
	}

}

func TestCreate(t *testing.T) {

	r := New(inmemoryctx.New())

	idx, err := r.Create("usr", "res1")

	if idx != 0 || err != nil {
		t.Fatalf("r.Create(\"usr\", \"res1\") = %v, %v, want 0, <nil>", idx, err)
	}

	idx, err = r.Create("usr", "res2")

	if idx != 1 || err != nil {
		t.Fatalf("r.Create(\"usr\", \"res2\") = %v, %v, want 0, <nil>", idx, err)
	}

}

func TestUpdate(t *testing.T) {

	r := New(inmemoryctx.New())

	err := r.Update("usr", 0, "res1")
	if err.Error() != "There is no receiver list for user usr" {
		t.Fatalf("r.Update(\"usr\", 0, \"res1\") = %v, want 'There is no receiver list for user usr'", err)
	}

	_, err = r.Create("usr", "res1")
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = r.Update("usr", 0, "res2")
	if err != nil {
		t.Fatalf("r.Update(\"usr\", 0, \"res2\")=%v, want <nil>", err)
	}

}

func TestDelete(t *testing.T) {

	r := New(inmemoryctx.New())

	err := r.Delete("usr", 1)

	if err.Error() != "There is no receiver list for user usr" {
		t.Fatalf("r.Delete(\"usr\", 1) = %v, want 'There is no receiver list for user usr'", err)
	}

	_, err = r.Create("usr", "res1")
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = r.Delete("usr", 1)
	if err.Error() != "There is no receiver with index 1" {
		t.Fatalf("r.Delete(\"usr\", 1) = %v, want 'There is no receiver with index 1'", err)
	}

	err = r.Delete("usr", 0)

	if err != nil {
		t.Fatalf("r.Delete(\"usr\", 0) = %v, want <nil>", err)
	}

}
