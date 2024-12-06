package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}
func posisi(k, n int) int {
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i + 1
		}
	}
	return -1
}
func main() {
	var n, k int
	fmt.Println("Masukkan jumlah elemen array (n) dan elemen yang dicari (k):")
	fmt.Scan(&n, &k)
	if n < 1 || n > NMAX {
		fmt.Println("Jumlah elemen array harus antara 1 hingga 1000000.")
		return
	}
	isiArray(n)
	pos := posisi(k, n)
	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Printf("%d\n", pos)
	}
}
