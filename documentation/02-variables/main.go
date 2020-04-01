package main

import "fmt"

func main() {
	var nama string = "Elisabeth Kartini"
	umur := 12
	umur = 13
	fmt.Println("Namaku", nama, "berumur", umur, "Tahun")

	// Variable Pembuangan
	aku, _ := "Namaku El", "dan Aku Cantiqueee"
	fmt.Println(aku)
	// fmt.Println(_) : Error

	// Pointers
	nama = "E. Kartini"
	var alamat *string = &nama
	fmt.Println(*alamat)
	fmt.Println(*alamat == nama) // true

	// manipulasi nilai *alamat
	*alamat = "Elisabeth Kartini"
	fmt.Println(nama == *alamat) // true

	// variable *alamat dan nama saling terhubung.
}
