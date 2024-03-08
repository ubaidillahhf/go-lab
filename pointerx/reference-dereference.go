package pointerx

import "fmt"

type Mahasiswa struct {
	NIM      uint
	Nama     string
	Fakultas string
}

func Basic() {
	string1 := "satu"
	string2 := string1

	fmt.Println(string1)
	fmt.Println(string2)

	// satu
	// satu

	//!
	/**
	Masih terlihat sama bukan? cuman jika dilihat dr alokasi memory string1 dan string2 berbeda (copy as value).
	untuk membuktikannya, maka bisa kita print alamat memorinya dengan "reference memory/&"
	*/

	// fmt.Println(&string1)
	// fmt.Println(&string2)

	// 0x14000010230
	// 0x14000010240

	//!
	/**
	Berbeda bukan? nah kalo semisal di assign nilai baru ke string2 maka string1 tidak akan terpengaruh (Beda kaya NodeJs).
	*/

	// string2 = "berubah!"

	// fmt.Println(string1)
	// fmt.Println(string2)

	// satu
	// berubah!

	//!
	/**
	Pertanyannya? gimana ngerubah biar pass by reference? tahukah kamu kenapa memakai pass by reference? THIS!!
	--"pass by reference when the value is large."--
	*/

	// string1 := "satu"
	// string2 := &string1
	// ATAU
	// var string1 string = "satu";
	// var string2 *string = &string1;

	// *string2 = "berubah!"

	// fmt.Println(string1)
	// fmt.Println(*string2)

}
