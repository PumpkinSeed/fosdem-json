package main

import (
	"encoding/json"
	"fmt"
)

type TestStruct struct {
	Data NullString `json:"data,omitempty"`
}

type NullString struct {
	Valid bool
	Value string
}

func (x *NullString) IsZero() bool {
	return !x.Valid
}

func main() {
	var s = TestStruct{
		Data: NullString{
			Valid: false,
			Value: "empty",
		},
	}

	result, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
