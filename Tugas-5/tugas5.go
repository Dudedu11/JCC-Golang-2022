package main

import (
	"fmt"
	"strings"
)

// func soal 1
func luasPersegiPanjang(panjang int, lebar int) (hasil int) {
	hasil = panjang * lebar
	return
}

func kelilingPersegiPanjang(panjang int, lebar int) (hasil int) {
	hasil = 2 * (panjang + lebar)
	return
}

func volumeBalok(panjang int, lebar int, tinggi int) (hasil int) {
	hasil = panjang * lebar * tinggi
	return
}

// function soal 2

func introduce(nama string, jk string, profesi string, umur string) string {
	var panggil string
	if strings.Contains(jk, "laki") == true {
		panggil = "Pak "
	} else {
		panggil = "Bu "
	}
	// 	return panggil, nama + "adalah seorang", profesi, "yang berusia", umur, "tahun"
	var oy string = panggil + "" + nama + " adalah seorang " + profesi + " yang berusia " + umur
	return oy
}

// function soal 3
func buahFavorit(name string, buah ...string) {
	fmt.Printf("halo nama saya %s", name)
	fmt.Printf(" dan buah favorit saya adalah ")
	for _, fruit := range buah {
		fmt.Printf(fruit)
		fmt.Printf(", ")
	}
}

// function soal 4
// func tambahDataFilm(judul string, durasi string, genre string, tahun string){

// }

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// kataPertama, kataKedua, angka := tampilkanKataDanAngka()
	// fmt.Println(kataPertama, kataKedua, angka)

	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavorit("john", buah...)

	// soal 4
	// var dataFilm = []map[string]string{}
	// func(){
	// 	dataFilm["genre"] = "action"
	// }

	// tambahDataFilm("LOTR", "2 jam", "action", "1999")
	// tambahDataFilm("avenger", "2 jam", "action", "2019")
	// tambahDataFilm("spiderman", "2 jam", "action", "2004")
	// tambahDataFilm("juon", "2 jam", "horror", "2004")

	// for _, item := range dataFilm {
	// fmt.Println(item)
	// }

}
