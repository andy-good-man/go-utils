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

func TestToMgoBsonSweet(t *testing.T) {
	type User struct {
		ID   string `json:"ID,omitempty" bson:"_id"`
		Name string `bson:"Name"`
		Age  int32  `bson:"Age"`
		// Age  int32
	}

	user := &User{
		ID:   "qwertyuiop",
		Name: "小明",
		Age:  25,
	}

	// data := ToMgoBsonSweet(user, nil, nil)

	// checks := []string{"_id"}
	// checks := []string{"_id", "Name"}
	// checks := []string{"_id", "Name", "Age"}

	dels := []string{"Name"}

	// data := ToMgoBsonSweet(user, checks, nil)
	data := ToMgoBsonSweet(user, nil, dels)
	// data := ToMgoBsonSweet(user, checks, dels)

	fmt.Println("data  :", data)
}
