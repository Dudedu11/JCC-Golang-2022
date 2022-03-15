package main

import (
	"fmt"
	"strings"
)

// function soal 1
func hitung(luas *float64, keliling *float64, jari float64) {
	*luas = (22 * jari * jari) / 7
	*keliling = (2 * 22 * jari) / 7
}

// function soal 2
func introduce(sentence *string, nama string, jk string, profesi string, umur string) {
	var panggil string
	if strings.Contains(jk, "laki") == true {
		panggil = "Pak "
	} else {
		panggil = "Bu "
	}
	// 	return panggil, nama + "adalah seorang", profesi, "yang berusia", umur, "tahun"
	*sentence = panggil + "" + nama + " adalah seorang " + profesi + " yang berusia " + umur
}

// function soal 3
// func input(buah *[]string, nama string) {
// 	buah *[] =  append(&buah, nama)
// }

// function soal 4

func main() {
	// soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64

	hitung(&luasLigkaran, &kelilingLingkaran, 7)
	fmt.Println(luasLigkaran)
	fmt.Println(kelilingLingkaran)

	// soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	// var buah = *[]string{}

	// input(&buah, Jeruk)
	// input(&buah, Semangka)
	// input(&buah, Mangga)
	// input(&buah, Strawberry)
	// input(&buah, Durian
	// input(&buah, Manggis)
	// input(&buah, Alpukat)

	// for y := 0; y < len(buah); y++ {
	// 	var z = y + 1
	// 	fmt.Println(z, ". "+buah[y])
	// }

	// soal 4

	// var dataFilm = []map[string]string{}

	// tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	// tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	// tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	// tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// isi dengan jawaban anda untuk menampilkan data

}
