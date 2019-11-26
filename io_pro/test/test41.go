package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type StructA struct {
	Username string `json:"user"`
	Process  string `json:"process"`
}

func main345676() {

	var test1 StructA
	err := json.Unmarshal([]byte(`{"user": "user123", "process": "something"}`), &test1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n",test1)
	// do some work with test1

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	//encoder.SetEscapeHTML()

	//无法更改struct的tag值，只能传创建struct
	//fmt.Printf("%v\n",&test1.Username)
	//elem := reflect.TypeOf(&test1).Elem().Field(0)
	//
	//fmt.Println(elem.Tag)
	//reflect.ValueOf(&elem.Tag).Elem().SetString("user12345")
	//fmt.Println(elem.Tag)
	//fmt.Printf("%v\n",&test1.Username)
	//
	//fmt.Println(reflect.TypeOf(test1).Field(0).Tag)
	//fmt.Println(reflect.TypeOf(&test1).Elem().Field(0).Tag.Lookup(`user`))
	//
	//
	//fmt.Println(reflect.TypeOf(test1).Field(0).Tag)

	err = encoder.Encode(&test1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buffer)


	jsonByte, err := json.Marshal(&test1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonByte))

}

func (u *StructA) MarshalJSON() ([]byte, error) {
	type Alias StructA
	return json.Marshal(&struct {
		Username string `json:"username"`
		*Alias
	}{
		Username: u.Username,
		Alias:    (*Alias)(u),
	})
}