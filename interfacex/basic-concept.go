package interfacex

import "fmt"

/**
MASALAHNYA
*/

type BangunDatarInt interface {
	HitungLuas() uint
}

type BangunDatarFloat interface {
	HitungLuas() float32
}

type Persegi struct {
	Sisi uint
}

func (persegi Persegi) HitungLuas() uint {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang uint
	Lebar   uint
}

func (persegiPanjang PersegiPanjang) HitungLuas() uint {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Segitiga struct {
	Alas   float32
	Tinggi float32
}

func (segitiga Segitiga) HitungLuas() float32 {
	return 0.5 * segitiga.Alas * segitiga.Tinggi
}

// Masalahnya adalah Persegi harus general (akomodir semua struct) => Ubah ke interface!!
func SeberapaLuasInt(bangunDatar BangunDatarInt) uint {
	return bangunDatar.HitungLuas()
}

func SeberapaLuasFloat(bangunDatar BangunDatarFloat) float32 {
	return bangunDatar.HitungLuas()
}

func Basic5() {
	var LuasBangunDatarInt uint
	var LuasBangunDatarFloat float32

	persegi1 := Persegi{4}
	LuasBangunDatarInt = SeberapaLuasInt(persegi1)
	fmt.Println(LuasBangunDatarInt)

	persegiPanjang1 := PersegiPanjang{4, 2}
	LuasBangunDatarInt = SeberapaLuasInt(persegiPanjang1)
	fmt.Println(LuasBangunDatarInt)

	segitiga1 := Segitiga{4.2, 2.2}
	LuasBangunDatarFloat = SeberapaLuasFloat(segitiga1)
	fmt.Println(LuasBangunDatarFloat)
}
