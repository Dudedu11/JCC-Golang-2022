package main

import (
	"errors"
	"fmt"
)

// function soal 1
func logging(kal string, taun int) {
	fmt.Printf(kal)
	fmt.Println(" ", taun)
}

func runApplication(kal string, taun int) {
	defer logging(kal, taun)
	fmt.Println("Halo!")
}

// function soal 2
func kelilingSegitigaSamaSisi(nilai int, status bool) (string, error) {
	if nilai == 0 && status == true {
		return string("0"), errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if nilai == 0 && status == false {
		return string("0"), errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if nilai != 0 && status == true {
		var keliling int = nilai * 3
		// str, _ := strconv.Itoa(nilai)
		// str1, _ := strconv.Itoa(keliling)
		//var oy string = "keliling segitiga sama sisinya dengan sisi", nilai, " cm adalah", keliling, "cm"
		return string(keliling), nil
	} else {
		var keliling int = nilai * 3
		// var oy sting = "keliling segitiga sama sisinya dengan sisi" + nilai + " cm adalah", keliling, "cm"
		return string(keliling), nil

	}
}

// function soal 3
func tambahAngka(nilai int, angka *int) {
	*angka = nilai
}

func cetakAngka(angka *int) {
	fmt.Println(angka)
}

func main() {
	// soal 1
	var kal string = "Golang Backend Development"
	var taun int = 2021
	runApplication(kal, taun)

	// soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// soal 3
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)
}
