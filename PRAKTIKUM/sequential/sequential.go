package main

import "fmt"

type arrStr [1234]string

func SeqSearch_1(T arrStr, n int, X string) int {
	var found int = -1
	var j int = 0

	for j < n && found == -1 {
		if T[j] == X {
			found = j
		}
		j = j + 1
	}
	return found
}
func main() {
	var data arrStr
	var n int

	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scanln(&n)

	fmt.Println("Masukkan elemen array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i+1)
		fmt.Scanln(&data[i])
	}
	var search string
	fmt.Print("Masukkan elemen yang ingin dicari: ")
	fmt.Scanln(&search)

	result := SeqSearch_1(data, n, search)

	if result == -1 {
		fmt.Println("Elemen tidak ditemukan dalam array.")
	} else {
		fmt.Printf("Elemen ditemukan pada indeks: %d\n", result)
	}
}
