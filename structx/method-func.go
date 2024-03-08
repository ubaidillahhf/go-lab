package structx

import "fmt"

/**
* So jadi ini buat method(function yg menjadi milik dr objek(class)/bukan function bebas), di golang biasanya di deklare seperti ini.
* funct (user User) display(){}
 */

func (student Student) Display() (result string) {
	result = fmt.Sprintf("Name: %s %s, Email: %s", student.FirstName, student.LastName, student.Email)
	return
}
