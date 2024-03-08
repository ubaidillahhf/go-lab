package pointerx

import "fmt"

func Basic2() {
	//!
	/**
	contoh penggunaan pada function change number
	*/

	number1 := 1
	fmt.Println("alamat 1:", &number1)
	changeNumber(&number1, 200)
	fmt.Println("nilai awal:", number1)
}

func changeNumber(number1 *int, number2 int) {
	fmt.Println("alamat 2:", number1)
	*number1 = number2
	fmt.Println(*number1)
}
