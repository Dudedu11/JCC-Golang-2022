package main

import "fmt"

func main() {
	var segitiga library.hitungBangunDatar
	var persegi library.hitungBangunDatar
	var tab library.hitungBangunRuang
	var bal library.hitungBangunRuang

	segitiga = library.segitigaSamaSisi{10, 7}
	persegi = library.persegiPanjang{10, 10}
	tab = library.tabung{7, 10}
	bal = library.balok{10, 10, 10}

	fmt.Println("===== Segitiga Sama Sisi")
	fmt.Println("luas      :", library.segitiga.luas())
	fmt.Println("keliling  :", library.segitiga.keliling())
	fmt.Println("===== Persegi panjang")
	fmt.Println("luas      :", library.persegi.luas())
	fmt.Println("keliling  :", library.persegi.keliling())
	fmt.Println("===== Tabung")
	fmt.Println("volume       :", library.tab.volume())
	fmt.Println("l.permukaan  :", library.tab.luasPermukaan())
	fmt.Println("===== Balok")
	fmt.Println("volume       :", library.bal.volume())
	fmt.Println("l.permukaan  :", library.bal.luasPermukaan())
}
