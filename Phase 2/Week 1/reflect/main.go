package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `required:"true" max:"100"`
	Age   int    `required:"true"`
	Level string `required:"false"`
}

func ValidatePersonData(s interface{}) error {
	t := reflect.TypeOf(s)
	fmt.Println("Number of fields:", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(s).Field(i).Interface()
			if value == "" {
				return fmt.Errorf("Field %s is required", field.Name)
			}
		}
	}
	return nil
}

func main() {
	// Create a new User
	u := User{
		Name: "John",
		Age:  25,
	}

	t := reflect.TypeOf(u)  // Get the type of the User
	k := t.Kind()           // Get the kind of the User
	v := reflect.ValueOf(u) // Get the value of the User

	fmt.Println("Type:", v)
	fmt.Println("Type:", t)
	fmt.Println("Type:", k)

	s := User{

		Age:   25,
		Level: "admin",
	}
	p := ValidatePersonData(s)
	fmt.Println(p)
}
