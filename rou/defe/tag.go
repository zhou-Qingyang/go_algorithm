package defe

import "fmt"

type User struct {
	ID   int    `json:"id" db:"user_id"`
	Name string `json:"username" db:"user_name"`
}

func init() {
	fmt.Println("defe GetTag")
}

func GetTag() {
	// user := User{
	// 	ID:   1,
	// 	Name: "John",
	// }
	// fmt.Printf("p: %v\n", user)
	// fmt.Printf("p: %+v\n", user)
	// fmt.Printf("p: %#v\n", user)

}
