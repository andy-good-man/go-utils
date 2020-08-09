package gostructs

import (
	"fmt"
	"testing"
)

func TestToMgomap(t *testing.T) {
	type User struct {
		ID   string `json:"ID,omitempty" bson:"_id"`
		Name string `bson:"Name"`
		Age  int32
	}

	user := &User{
		ID:   "qwertyuiop",
		Name: "小明",
		Age:  25,
	}

	data := ToMgoBson(user, nil)
	fmt.Println("data  :", data)
}
