package main

import (
	"fmt"
)

func main() {
	var panjang int = 5
	var tinggi int = 5

	cetakGambar(panjang, tinggi)
}
func cetakGambar(panjang int, tinggi int) {

	// validasi angka harus ganjil
	if panjang%2 == 0 || tinggi%2 == 0 {
		fmt.Println("error: angka panjang harus ganjil")
		return
	}

	fmt.Println("--- panjang ---")

	for i := 0; i < panjang; i++ {
		for j := 1; j <= tinggi; j++ {
			// print * jika tinggi ditengah
			if i == tinggi/2 {
				fmt.Print(" * ")
				continue
			}
			// print * diawal panjang dan diakhir panjang
			if j == 1 || j == panjang {
				fmt.Print(" * ")
				continue
			}
			// selain perjikaan di atas print =
			fmt.Print(" = ")
		}
		fmt.Println()
	}
}
