package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var text1 = "Jabar"
	var text2 = "Coding"
	var text3 = "Camp"
	var text4 = "Golang"
	var text5 = "2022"

	fmt.Println(text1, text2, text3, text4, text5)

	// soal 2
	halo := "Halo Dunia"
	find := "Dunia"
	replacewith := "Golang"

	newText := strings.Replace(halo, find, replacewith, 1)
	fmt.Println(newText)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var num1, err1 = strconv.Atoi(angkaPertama)
	var num2, err2 = strconv.Atoi(angkaKedua)
	var num3, err3 = strconv.Atoi(angkaKetiga)
	var num4, err4 = strconv.Atoi(angkaKeempat)

	total := num1 + num2 + num3 + num4

	if err1 == nil {
		if err2 == nil {
			if err3 == nil {
				if err4 == nil {
					fmt.Println(total)
				}
			}
		}
	}

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2022

	find1 := "halo"
	replacewith1 := "Hi"

	newText1 := strings.Replace(kalimat, find1, replacewith1, 2)
	fmt.Println(newText1, "-", angka)

}
