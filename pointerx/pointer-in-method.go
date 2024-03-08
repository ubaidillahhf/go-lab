package pointerx

import "fmt"

func Basic3() {
	//!
	/**
	contoh penggunaan pada function dengan struct
	*/

	Mahasiswa1 := Mahasiswa{
		NIM:      1551506021097007,
		Nama:     "Ubaidillah Hakim Fadly",
		Fakultas: "FILKOM",
	}
	fmt.Println(Mahasiswa1.Nama)

	// Cara pemanggilan 1
	// graduated1(&Mahasiswa1)
	// Cara pemanggilan dengan method
	Mahasiswa1.graduated()

	fmt.Println("Gelar: ", Mahasiswa1.Nama)
}

// Func basic
func graduated1(mahasiswa *Mahasiswa) {
	mahasiswa.Nama = mahasiswa.Nama + " S.Kom"
}

// Func With method
func (mahasiswa *Mahasiswa) graduated() {
	mahasiswa.Nama = mahasiswa.Nama + " S.Kom"
}
