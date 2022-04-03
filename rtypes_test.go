package rtypes_test

import (
	"testing"

	"github.com/dewidyabagus/rtypes"
	"github.com/stretchr/testify/assert"
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

func TestConvertStructToMapInterface(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r.(string))
		}
	}()

	valueForTests := [][]string{}
	valueForTests = append(valueForTests, []string{"FirstName", "Joe Do", "first name must be same"})
	valueForTests = append(valueForTests, []string{"LastName", "Lorem", "last name must be same"})
	valueForTests = append(valueForTests, []string{"Username", "joedolorem", "username must be same"})
	valueForTests = append(valueForTests, []string{"Email", "joedolorem@example.com", "email must be same"})
	valueForTests = append(valueForTests, []string{"Address", "example address", "address must be same"})

	user := User{
		FirstName: "Joe Do",
		LastName:  "Lorem",
		Username:  "joedolorem",
		Email:     "joedolorem@example.com",
		Address:   "example address",
		MapData: map[string]interface{}{
			"Field1": true,
			"Field2": 1.3,
			"Field3": "Value Of Field3",
		},
		StructData: &struct {
			Field1 bool `map:"Field1"`
			Field2 int
			Field3 string
		}{false, 5, "lorem"},
	}

	response, err := rtypes.ConvertStructToMapInterface(&user)
	assert.Nil(t, err)
	for _, item := range valueForTests {
		assert.Equal(t, response[item[0]].(string), item[1], item[2])
	}
	assert.Equal(t, response["MapData"].(map[string]interface{})["Field1"].(bool), true, "MapData Field1 must be true value")
	assert.Equal(t, response["MapData"].(map[string]interface{})["Field2"].(float64), 1.3, "MapData Field2 must be 1.3 value")
	assert.Equal(t, response["MapData"].(map[string]interface{})["Field3"].(string), "Value Of Field3", "MapData Field3 must be 'Value Of Field3' value")
	assert.Equal(t, response["StructData"].(map[string]interface{})["Field1"].(bool), false, "StructData Field1 must be false value")
	_, found := response["StructData"].(map[string]interface{})["Field2"]
	assert.Equal(t, found, false, "StructData Field2 must not found key")

	// Error References Pointer Value
	newUser := &user
	_, err = rtypes.ConvertStructToMapInterface(&newUser)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "data type cannot be processed", "invalid argument entered")

	// Set Argument Nil Value
	resWithArgumentNilValue, _ := rtypes.ConvertStructToMapInterface(nil)
	assert.Equal(t, len(resWithArgumentNilValue), 0, "must be response 0 field")

	// Set Argument Non Struct
	_, err = rtypes.ConvertStructToMapInterface("string")
	assert.Contains(t, err.Error(), "data type cannot be processed", "invalid argument entered")
}
