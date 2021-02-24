package tool

import (
	"testing"
)

type T struct {
	Age int
}
type T2 struct {
	Age int32
}

type Test1 struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Hello2 string
}

type Test2 struct {
	Name  string `json:"name"`
	Hello string
	Age   int32 `json:"age"`
}

func TestEqualStruct(t *testing.T) {
	t2 := &Test2{}
	err := TranslateStruct(Test1{
		Name:   "11111",
		Age:    12,
		Hello2: "!!!!!",
	}, t2)
	t.Error(t2, err)
}
