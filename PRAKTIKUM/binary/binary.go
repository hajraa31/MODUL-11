package main

import "fmt"

type arrInt [4321]int

func BinarySearch_2(T arrInt, n int, X int) int {
	var found int = -1
	var med int
	var kr int = 0
	var kn int = n - 1

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if X > (med) {
			kn = med - 1
		} else if X < (med) {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

func main() {
	var data arrInt
	var n int
	fmt.Print("Masukkan jumlah elemen dalam array:")
	fmt.Scanln(&n)
	fmt.Println("Masukkan elemen array (harus terurut menurun):")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i+1)
		fmt.Scanln(&data[i])
	}

	var search int
	fmt.Print("Masukkan elemen yang ingin dicari:")
	fmt.Scanln(&search)
	result := BinarySearch_2(data, n, search)

	if result == -1 {
		fmt.Println("Elemen tidak ditemukan dalam array.")
	} else {
		fmt.Printf("Elemen ditemukan pada indeks: %d\n", result)
	}
}
