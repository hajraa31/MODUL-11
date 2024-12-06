package main

import "fmt"

func main() {
	var suara int
	hitungSuara := make(map[int]int)
	jumlahData := 0
	fmt.Println("Masukkan angka suara satu per satu (antara 1-20). Akhiri dengan angka 0:")
	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		if suara < 1 || suara > 20 {
			fmt.Println("Masukkan angka valid antara 1 hingga 20, atau 0 untuk mengakhiri.")
			continue
		}
		hitungSuara[suara]++
		jumlahData++
	}
	ketua := 0
	wakilKetua := 0
	maxSuaraKetua := 0
	maxSuaraWakil := 0
	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > maxSuaraKetua {
			wakilKetua = ketua
			maxSuaraWakil = maxSuaraKetua
			ketua = i
			maxSuaraKetua = hitungSuara[i]
		} else if hitungSuara[i] > maxSuaraWakil {
			wakilKetua = i
			maxSuaraWakil = hitungSuara[i]
		} else if hitungSuara[i] == maxSuaraKetua && i < ketua {
			wakilKetua = ketua
			maxSuaraWakil = maxSuaraKetua
			ketua = i
		} else if hitungSuara[i] == maxSuaraWakil && i < wakilKetua {
			wakilKetua = i
		}
	}
	fmt.Println("\nJumlah data suara yang terbaca:", jumlahData)
	fmt.Println("Distribusi suara:")
	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			fmt.Printf("Calon %d: %d suara\n", i, hitungSuara[i])
		}
	}
	fmt.Printf("\nKetua RT: %d\n", ketua)
	fmt.Printf("Wakil Ketua RT: %d\n", wakilKetua)
}
