package tests

import (
	"github.com/JekaMas/easyjson"
	"bytes"
	"testing"
)

func TestIncludeFields(t *testing.T) {
	n := ExcludeFields{1, 2, []string{"Int2"}}

	res, err := easyjson.Marshal(n)
	if err != nil {
		t.Fatal("Got an error", err)
	}

	expected := `{"Int2":2}`
	if !bytes.Equal(res, []byte(expected)) {
		t.Fatal("Got", string(res), "\t Expected", expected)
	}
}

func TestIncludeFields_FirstField(t *testing.T) {
	n := ExcludeFields{1, 2, []string{"Int1"}}

	res, err := easyjson.Marshal(n)
	if err != nil {
		t.Fatal("Got an error", err)
	}

	expected := `{"Int1":1}`
	if !bytes.Equal(res, []byte(expected)) {
		t.Fatal("Got", string(res), "\t Expected", expected)
	}
}

func TestIncludeFields_All_EmptyIncludeFields(t *testing.T) {
	n := ExcludeFields{1, 2, []string{}}

	res, err := easyjson.Marshal(n)
	if err != nil {
		t.Fatal("Got an error", err)
	}

	expected := `{"Int1":1,"Int2":2}`
	if !bytes.Equal(res, []byte(expected)) {
		t.Fatal("Got", string(res), "\t Expected", expected)
	}
}

func TestIncludeFields_All_AllFields(t *testing.T) {
	n := ExcludeFields{1, 2, []string{"Int1", "Int2"}}

	res, err := easyjson.Marshal(n)
	if err != nil {
		t.Fatal("Got an error", err)
	}

	expected := `{"Int1":1,"Int2":2}`
	if !bytes.Equal(res, []byte(expected)) {
		t.Fatal("Got", string(res), "\t Expected", expected)
	}
}