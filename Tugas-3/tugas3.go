package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var num1, err1 = strconv.Atoi(panjangPersegiPanjang)
	var num2, err2 = strconv.Atoi(lebarPersegiPanjang)

	var num3, err3 = strconv.Atoi(alasSegitiga)
	var num4, err4 = strconv.Atoi(tinggiSegitiga)

	if err1 == nil {
		if err2 == nil {
			if err3 == nil {
				if err4 == nil {
					var kelilingPersegiPanjang int = 2 * (num1 + num2)
					var luasSegitiga int = (num3 * num4) / 2
					fmt.Println(kelilingPersegiPanjang)
					fmt.Println(luasSegitiga)
				}
			}
		}
	}

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("indeksnya A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("indeksnya B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("indeksnya C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("indeksnya D")
	} else if nilaiJohn < 50 {
		fmt.Println("indeksnya E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("indeksnya A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("indeksnya B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("indeksnya C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("indeksnya D")
	} else if nilaiDoe < 50 {
		fmt.Println("indeksnya E")
	}

	// soal 3
	var tanggal = 1
	var bulan = 1
	var tahun = 2001

	switch bulan {
	case 1:
		bulan := "Januari"
		fmt.Println(tanggal, bulan, tahun)
	case 2:
		bulan := "Februari"
		fmt.Println(tanggal, bulan, tahun)
	case 3:
		bulan := "Maret"
		fmt.Println(tanggal, bulan, tahun)
	case 4:
		bulan := "April"
		fmt.Println(tanggal, bulan, tahun)
	case 5:
		bulan := "Mei"
		fmt.Println(tanggal, bulan, tahun)
	case 6:
		bulan := "Juni"
		fmt.Println(tanggal, bulan, tahun)
	case 7:
		bulan := "Juli"
		fmt.Println(tanggal, bulan, tahun)
	case 8:
		bulan := "Agustus"
		fmt.Println(tanggal, bulan, tahun)
	case 9:
		bulan := "September"
		fmt.Println(tanggal, bulan, tahun)
	case 10:
		bulan := "Oktober"
		fmt.Println(tanggal, bulan, tahun)
	case 11:
		bulan := "November"
		fmt.Println(tanggal, bulan, tahun)
	case 12:
		bulan := "Desember"
		fmt.Println(tanggal, bulan, tahun)
	default:
		fmt.Println("Tidak terjadi apa-apa")
	}

	// soal 4
	var kelahiran = 2001
	if kelahiran >= 1944 && kelahiran <= 1964 {
		fmt.Println("Anda generasi Baby boomer")
	} else if kelahiran >= 1965 && kelahiran <= 1979 {
		fmt.Println("Anda generasi X")
	} else if kelahiran >= 1980 && kelahiran <= 1994 {
		fmt.Println("Anda generasi Y(Millenials)")
	} else if kelahiran >= 1965 && kelahiran <= 2015 {
		fmt.Println("Anda generasi Z")
	}

}
