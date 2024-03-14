package arraylist_test

import (
	"reflect"
	"testing"

	"github.com/alcocom/base-go/list/arraylist"
)

func TestMethod_New(t *testing.T) {
	paramNotExist := arraylist.New[int]()
	paramExistFixedType := arraylist.New("H", "E", "LLO")
	paramExistAnyType := arraylist.New[any](1, 2, "h", "ello")

	if reflect.TypeOf(paramNotExist).String() != "*arraylist.ArrayList[int]" { // slice == []int
		t.Error()
	}

	if reflect.TypeOf(paramExistFixedType).String() != "*arraylist.ArrayList[string]" { // slice == []string
		t.Error()
	}

	if reflect.TypeOf(paramExistAnyType).String() != "*arraylist.ArrayList[interface {}]" { // slice == "[]interface {}"
		t.Error()
	}
}
