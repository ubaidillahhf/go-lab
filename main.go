package main

import (
	"fmt"
	"go/contex"
)

func main() {
	//! Case 1
	// tes := structx.StructInStructExample()
	// fmt.Println(tes)

	//! Case 2
	// user1 := structx.DisplayUser(structx.Student{1, "Ubaidillah", "hakim", "ubaidillahhf@gmail.com", true})
	// user2 := structx.DisplayUser(structx.Student{2, "Inamul", "fadly", "inamul.fadly@gmail.com", true})

	// fmt.Println(user1)
	// fmt.Println(user2)

	// NB: Kalo dibikin kaya poin 4 tapi tanpa param method, maka akan kena error:
	// user1.DisplayUser undefined (type structx.Student has no field or method DisplayUser)

	//! Case 3
	// type AddressList struct {
	// 	Name    string
	// 	Address []structx.User
	// }
	// user1 := structx.User{1, "Ubaidillah", "hakim", structx.Address{Province: "East Java"}}
	// user2 := structx.User{2, "Inamul", "fadly", structx.Address{Province: "West Java"}}

	// users := []structx.User{user1, user2}

	// sector1 := AddressList{"SECTOR I", users}
	// fmt.Println(sector1)

	//! Case 4 (Method in func)
	// user1 := structx.Student{1, "Ubaidillah", "hakim", "ubaidillahhf@gmail.com", true}
	// user2 := structx.Student{2, "Inamul", "fadly", "inamul.fadly@gmail.com", true}

	// fmt.Println(user1.Display())
	// fmt.Println(user2.Display())

	//! Case 5 (Pointer)
	// interfacex.Basic5()

	//! Case 6 (Context)
	dataNumbers := []int{7, 2, 8, -9, 4, 0}

	// Ini test aja sih sebenernya
	channel := make(chan int)
	go contex.Sum(dataNumbers[:len(dataNumbers)/2], channel)
	go contex.Sum(dataNumbers[len(dataNumbers)/2:], channel)
	x, y := <-channel, <-channel

	fmt.Println(x, y, x+y)

}
