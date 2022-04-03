package main

import (
	"fmt"
	"log"

	"github.com/dewidyabagus/rtypes"
)

type User struct {
	FirstName  string                 `map:"FirstName"`
	LastName   string                 `map:"LastName"`
	Username   string                 `map:"Username"`
	Email      string                 `map:"Email"`
	Address    string                 `map:"Address"`
	MapData    map[string]interface{} `map:"MapData"`
	StructData *struct {
		Field1 bool `map:"Field1"`
		Field2 int
		Field3 string
	} `map:"StructData"`
}

func main() {
	user := &User{
		FirstName: "Joe Do",
		LastName:  "Lorem",
		Username:  "joedolorem",
		Email:     "joedolorem@example.com",
		Address:   "example address",
		MapData: map[string]interface{}{
			"MapField1": 0,
			"MapField2": true,
			"MapField3": "Field3",
		},
		StructData: &struct {
			Field1 bool `map:"Field1"`
			Field2 int
			Field3 string
		}{true, 1, "Testing"},
	}
	response, err := rtypes.ConvertStructToMapInterface(user)
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println(response)
	}
}
