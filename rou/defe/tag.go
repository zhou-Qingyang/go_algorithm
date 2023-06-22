package defe

import (
	"fmt"
)

type User struct {
	ID   int    `json:"id" db:"user_id"`
	Name string `json:"username" db:"user_name"`
}

func GetTag() {
	user := User{
		ID:   1,
		Name: "John",
	}

	// t := reflect.TypeOf(user)
	// for i := 0; i < t.NumField(); i++ {
	// 	flag := t.Field(i)
	// 	fmt.Printf("tag的 Name:%s,ID : %s,DB : %s\n", flag.Name, flag.Tag.Get("json"), flag.Tag.Get("db"))
	// }
	fmt.Printf("p: %v\n", user)
	fmt.Printf("p: %+v\n", user)
	fmt.Printf("p: %#v\n", user)

}
