# RTYPES (Reflection Types)
Simple package convertion for type struct to map with go reflection.

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go Version](https://img.shields.io/github/go-mod/go-version/dewidyabagus/rtypes)](https://github.com/dewidyabagus/rtypes)
[![codecov](https://codecov.io/gh/dewidyabagus/rtypes/branch/master/graph/badge.svg?token=ZV2OL9ULKO)](https://codecov.io/gh/dewidyabagus/rtypes)

---

* [Instalation](#instalation)
* [Example](#example)
* [Feedback](#feedback)
---

# Instalation
Install package in golang module
```consile
go get -u github.com/dewidyabagus/rtypes
```

# Example
Example of package usage
```go
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
```

See more [Example](https://github.com/dewidyabagus/rtypes/tree/master/example)!

# Feedback
I really hope for suggestions and feedback, you can contact me at [linkedin](https://www.linkedin.com/in/widya-ade-bagus-3a660716b/).
