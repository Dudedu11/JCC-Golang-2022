package library

import "fmt"

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi float64
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (l segitigaSamaSisi) luas() int {
	return (l.alas * l.tinggi) / 2
}

func (l segitigaSamaSisi) keliling() int {
	return l.alas * 3
}

func (l persegiPanjang) luas() int {
	return l.lebar * l.panjang
}

func (l persegiPanjang) keliling() int {
	return (l.lebar + l.panjang) * 2
}

func (l tabung) volume() float64 {
	return (22 * l.jariJari * l.jariJari * l.tinggi) / 7
}

func (l tabung) luasPermukaan() float64 {
	return ((22 * l.jariJari * l.jariJari) / 7) + ((2 * 22 * l.jariJari * l.tinggi) / 7)
}

func (l balok) volume() float64 {
	return l.lebar * l.panjang * l.tinggi
}

func (l balok) luasPermukaan() float64 {
	return 2 * (l.panjang*l.lebar + l.panjang*l.tinggi + l.lebar*l.tinggi)
}

func oy() {
	var segitiga hitungBangunDatar
	var persegi hitungBangunDatar
	var tab hitungBangunRuang
	var bal hitungBangunRuang

	segitiga = segitigaSamaSisi{10, 7}
	persegi = persegiPanjang{10, 10}
	tab = tabung{7, 10}
	bal = balok{10, 10, 10}

	fmt.Println("===== Segitiga Sama Sisi")
	fmt.Println("luas      :", segitiga.luas())
	fmt.Println("keliling  :", segitiga.keliling())
	fmt.Println("===== Persegi panjang")
	fmt.Println("luas      :", persegi.luas())
	fmt.Println("keliling  :", persegi.keliling())
	fmt.Println("===== Tabung")
	fmt.Println("volume       :", tab.volume())
	fmt.Println("l.permukaan  :", tab.luasPermukaan())
	fmt.Println("===== Balok")
	fmt.Println("volume       :", bal.volume())
	fmt.Println("l.permukaan  :", bal.luasPermukaan())
}
