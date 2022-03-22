package main

import (
	"fmt"
	"sync"
	"time"
)

// function soal 1
func uy(s []string, wg *sync.WaitGroup) {
	for i := 0; i < len(s); i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i+1, ".", s[i])
	}
	wg.Done()
}

// function soal 2
func getMovies(ch chan string, movies []string) {
	for x := 0; x < len(movies); x++ {
		ch <- movies[x]
	}
	close(ch)
}

// function soal 3
func luaslingkaran(jariJari []int, ch chan int) {
	for i := 0; i < len(jariJari); i++ {
		ch <- (22 * jariJari[i] * jariJari[i]) / 7
	}
	close(ch)
}

func kelilinglingkaran(jariJari []int, ch chan int) {
	for i := 0; i < len(jariJari); i++ {
		ch <- (22 * 2 * jariJari[i]) / 7
	}
	close(ch)
}

func volumetabung(jariJari []int, tinggi int, ch chan int) {
	for i := 0; i < len(jariJari); i++ {
		ch <- ((22 * jariJari[i] * jariJari[i]) / 7) * tinggi
	}
	close(ch)
}

// function soal 4
func persegipanjang(panjang int, lebar int, ch chan int) {
	ch <- panjang * lebar
	close(ch)
}

func kelilingpersegi(panjang int, lebar int, ch chan int) {
	ch <- 2 * (panjang + lebar)
	close(ch)
}

func volumebalok(panjang int, lebar int, tinggi int, ch chan int) {
	ch <- panjang * lebar * tinggi
	close(ch)
}

func main() {
	// soal 1
	var wg sync.WaitGroup
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	wg.Add(1)
	go uy(phones, &wg)
	wg.Wait()

	// soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	var jariJari = []int{8, 14, 20}
	var tinggib = 10

	luas := make(chan int)
	keliling := make(chan int)
	vtabung := make(chan int)

	go luaslingkaran(jariJari, luas)

	fmt.Println("Luas lingkaran")
	for value := range luas {
		fmt.Println(value)
	}

	go kelilinglingkaran(jariJari, keliling)
	fmt.Println("Keliling lingkaran")
	for value1 := range keliling {
		fmt.Println(value1)
	}

	go volumetabung(jariJari, tinggib, vtabung)
	fmt.Println("Volume Tabung")
	for value2 := range vtabung {
		fmt.Println(value2)
	}

	// soal 4
	var panjang = 10
	var lebar = 10
	var tinggi = 10

	luasp := make(chan int)
	kelilingp := make(chan int)
	volumeb := make(chan int)
	go persegipanjang(panjang, lebar, luasp)
	go kelilingpersegi(panjang, lebar, kelilingp)
	go volumebalok(panjang, lebar, tinggi, volumeb)
	fmt.Println("Luas persegi panjang")
	for value3 := range luasp {
		fmt.Println(value3)
	}
	fmt.Println("Keliling lingkaran ")
	for value4 := range kelilingp {
		fmt.Println(value4)
	}
	fmt.Println("Volume balok ")
	for value5 := range volumeb {
		fmt.Println(value5)
	}

}
