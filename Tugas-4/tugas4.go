package main

import (
	"fmt"
)

func main() {
	// soal 1
	for i := 1; i <= 20; i++ {
		if i%2 != 0 && i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else if i%2 != 0 {
			fmt.Println(i, "- JCC")
		} else {
			fmt.Println(i, "- Candradimuka")
		}
	}

	// soal 2
	for r := 1; r <= 7; r++ {
		for f := 1; f <= r; f++ {
			fmt.Printf("#")
		}
		fmt.Println("")
	}

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	for x := 0; x < len(kalimat); x++ {
		fmt.Printf(kalimat[x] + " ")
	}
	fmt.Println("")

	// soal 4
	var sayuran = [7]string{}
	sayuran[0] = "Bayam"
	sayuran[1] = "Buncis"
	sayuran[2] = "Kangkung"
	sayuran[3] = "Kubis"
	sayuran[4] = "Seledri"
	sayuran[5] = "Tauge"
	sayuran[6] = "Timun"

	for y := 0; y < len(sayuran); y++ {
		var z = y + 1
		fmt.Println(z, ". "+sayuran[y])
	}

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for key, val := range satuan {
		fmt.Println(key, "=", val)
	}
	fmt.Println("Volume Balok =", satuan["panjang"]*satuan["lebar"]*satuan["tinggi"])

}
