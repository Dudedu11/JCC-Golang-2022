package main

import "fmt"

//struct soal 1
type buah struct {
	nama  string
	warna string
	biji  bool
	harga int
}

//struct soal 2

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luassegitiga() {
	fmt.Println((s.alas * s.tinggi) / 2)
}

func (s persegi) luaspersegi() {
	fmt.Println(s.sisi * s.sisi)
}

func (s persegiPanjang) luaspp() {
	fmt.Println(s.lebar * s.panjang)
}

func main() {
	// soal 1
	var john buah
	john.nama = "Nanas"
	john.warna = "Kuning"
	john.biji = false
	john.harga = 9000

	var john1 buah
	john1.nama = "Jeruk"
	john1.warna = "Orange"
	john1.biji = true
	john1.harga = 8000

	var john2 buah
	john2.nama = "Semangka"
	john2.warna = "Hijau & Merah"
	john2.biji = true
	john2.harga = 10000

	var john3 buah
	john3.nama = "Pisang"
	john3.warna = "Kuning"
	john3.biji = false
	john3.harga = 5000

	fmt.Println("Nama  :", john.nama)
	fmt.Println("Warna :", john.warna)
	fmt.Println("biji :", john.biji)
	fmt.Println("harga :", john.harga)
	fmt.Println("")
	fmt.Println("Nama  :", john1.nama)
	fmt.Println("Warna :", john1.warna)
	fmt.Println("biji :", john1.biji)
	fmt.Println("harga :", john1.harga)
	fmt.Println("")
	fmt.Println("Nama  :", john2.nama)
	fmt.Println("Warna :", john2.warna)
	fmt.Println("biji :", john2.biji)
	fmt.Println("harga :", john2.harga)
	fmt.Println("")
	fmt.Println("Nama  :", john3.nama)
	fmt.Println("Warna :", john3.warna)
	fmt.Println("biji :", john3.biji)
	fmt.Println("harga :", john3.harga)
	fmt.Println("")

	// soal 2
	var segi = segitiga{5, 10}
	var perse = persegi{10}
	var persep = persegiPanjang{10, 10}

	segi.luassegitiga()
	perse.luaspersegi()
	persep.luaspp()
}
