package main

import "fmt"

func main() {
	var nama string = "Elisabeth Kartini"
	umur := 12
	umur = 13
	fmt.Println("Namaku", nama, "berumur", umur, "Tahun")

	// Pointers
	nama = "E. Kartini"
	var alamat *string = &nama
	fmt.Println(*alamat)
	fmt.Println(*alamat == nama) // true

	// manipulasi nilai *alamat
	*alamat = "Elisabeth Kartini"
	fmt.Println(nama == *alamat) // true

	// *alamat dan nama ini terhubung.
}
