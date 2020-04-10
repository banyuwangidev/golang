# Tipe data

Tipe data menentukan bagaimana sebuah nilai akan mendapat perlakuan, maka dari itu kita perlu mengenal beberapa tipe data pada Go.

Karena GO dikompile ke _native executable_ maka tipe seperti integer dan float memiliki  jenis turunan berdasarkan lebar bit.

## Tipe Data Primitive
Tipe data dasar yang menampung nilai masukan primitive, seperti angka. karakter, dan lojik.
### Data Numerik

Data ini berisikan angka, baik dalam bilangan bulat dan desimal. Tipe ini : 
- `int` (integer/bilangan bulat)

Dalam terdapat 4 turunan integer :
	- int8
	- int16
	- int32
	- int64
	
masing - masing turunan menampung data berdasarkan bit biner 2 pangkat nomor dibelakang mereka.

Contoh:
`int8` mampu menampung  2^8 = 256 (0 sampai 255,sesuai kaidah biner).

- `float`(bilangan desimal)
Bilangan ini terbagi menjadi dua turunan :
	-`float32` (Definisi float harus menyertakan ukuran bitnya)
	-`float64`

ukuran bit 32 dan 64 merupakan kemampuan bilangan pecahan untuk menampung nilai dibelakang koma. 

### Data Karakter
[WIP]

### Data Lojik
[WIP]

## Operator 
[WIP]