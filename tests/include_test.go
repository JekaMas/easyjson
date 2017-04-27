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
