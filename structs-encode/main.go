package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`               // this field tag use for we don't want to encode field
	Permissions permissions `json:"perms,omitempty"` // omitempty is used to not encode the nil value
}

func main() {

	users := []user{
		{"emirhan", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "666", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	//json package only imports exported fields
	fmt.Println(string(out))

}
