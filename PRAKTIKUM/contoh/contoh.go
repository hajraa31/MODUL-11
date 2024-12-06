package main

import (
	"fmt"
)

// Definisi struct mahasiswa
type mahasiswa struct {
	nama    string
	nim     string
	kelas   string
	jurusan string
	ipk     float64
}

// Definisi array bertipe struct mahasiswa
type arrMhs [2023]mahasiswa

// Sequential Search berdasarkan nama
func SeqSearch_3(T arrMhs, n int, X string) int {
	/* Mengembalikan indeks mahasiswa dengan nama X,
	   atau -1 apabila tidak ditemukan pada array T
	   yang berisi n data mahasiswa */
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if T[j].nama == X {
			found = j
		}
		j = j + 1
	}
	return found
}

// Binary Search berdasarkan nim
func BinarySearch_3(T arrMhs, n int, X string) int {
	/* Mengembalikan indeks mahasiswa dengan nim X,
	   atau -1 apabila tidak ditemukan pada array T
	   yang berisi n data mahasiswa dan terurut membesar berdasarkan nim */
	var found int = -1
	var med int
	var kr int = 0
	var kn int = n - 1

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if X < T[med].nim {
			kn = med - 1
		} else if X > T[med].nim {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

// Fungsi utama untuk demonstrasi
func main() {
	var data arrMhs
	var n int

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scanln(&n)

	// Input data mahasiswa
	fmt.Println("Masukkan data mahasiswa:")
	for i := 0; i < n; i++ {
		fmt.Printf("Mahasiswa %d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scanln(&data[i].nama)
		fmt.Print("NIM: ")
		fmt.Scanln(&data[i].nim)
		fmt.Print("Kelas: ")
		fmt.Scanln(&data[i].kelas)
		fmt.Print("Jurusan: ")
		fmt.Scanln(&data[i].jurusan)
		fmt.Print("IPK: ")
		fmt.Scanln(&data[i].ipk)
	}

	// Pencarian berdasarkan nama
	var cariNama string
	fmt.Print("Masukkan nama mahasiswa yang ingin dicari: ")
	fmt.Scanln(&cariNama)
	posNama := SeqSearch_3(data, n, cariNama)
	if posNama == -1 {
		fmt.Println("Mahasiswa dengan nama tersebut tidak ditemukan.")
	} else {
		fmt.Printf("Mahasiswa ditemukan pada indeks: %d\n", posNama)
	}

	// Pencarian berdasarkan nim
	var cariNIM string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari: ")
	fmt.Scanln(&cariNIM)
	posNIM := BinarySearch_3(data, n, cariNIM)
	if posNIM == -1 {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
	} else {
		fmt.Printf("Mahasiswa ditemukan pada indeks: %d\n", posNIM)
	}
}
