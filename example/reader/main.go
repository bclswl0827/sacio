package main

import (
	"fmt"
	"reflect"

	"github.com/bclswl0827/sacio"
)

func main() {
	var sacData sacio.SACData

	// Read SAC file
	err := sacData.Read("./testdata.sac")
	if err != nil {
		panic(err)
	}

	// Print structured data
	printFields(sacData)
}

func printFields(obj any) {
	value := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	if typ.Kind() != reflect.Struct {
		fmt.Println("Object is not a struct")
		return
	}

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := typ.Field(i)

		fmt.Printf("%s: %v\n", fieldType.Name, fieldValue.Interface())
	}
}
