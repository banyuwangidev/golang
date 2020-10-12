
# Tipe data

Tipe data menentukan bagaimana sebuah nilai akan mendapat perlakuan, maka dari itu kita perlu mengenal beberapa tipe data pada Go.

Karena GO dikompile ke _native executable_ maka tipe seperti integer dan float memiliki jenis turunan berdasarkan lebar bit.


## Tipe Data Primitive

Tipe data dasar yang menampung nilai masukan primitive, seperti angka. karakter, dan lojik.

### Data Numerik

Data ini berisikan angka, baik dalam bilangan bulat dan desimal. Tipe ini :

-  `int` (integer/bilangan bulat)

Dalam terdapat 4 turunan integer :

- int8

- int16

- int32

- int64

masing - masing turunan menampung data berdasarkan bit biner 2 pangkat nomor dibelakang mereka.

Contoh:

-  `int8` mampu menampung 2^8 = 256 (0 sampai 255,sesuai kaidah biner).

-  `float`(bilangan desimal)

Bilangan ini terbagi menjadi dua turunan :

-`float32` (Definisi float harus menyertakan ukuran bitnya)

-`float64`

ukuran bit 32 dan 64 merupakan kemampuan bilangan pecahan untuk menampung nilai dibelakang koma.

 
### Data Karakter
Karakter merupakan tipe yang terdiri dari alphanumeric dan symbol: `string`.

```go
var  slogan  string  =  "Tiada hari Tanpa ngoding"
```

### Data Lojik
Go juga dapat menyimpan berupa logika benar atau salah : `bool`.

```go
var  akuNolep  bool  = true
```
## Ekspresi

Ekpresi merupakan cara mengunkapkan pernyataan, mulai dari menyatakan rasa suka, penghormatan, matematis, data hingga kode yang kita tulis merupakan bentuk ekspresi tertentu.

> Ekspresi = operator/fungsi dan operan 

Data yang tidak dioperasikan maka jangan dideklarasikan, untuk itu kita akan mengenal bagaimana mana beragam tipe data ditindak lanjuti. 

Untuk mengoperasikan suatu nilai kita perlu operator yang sesuai dengan kaidah expresi berikut.

### Ekspresi Aritmatika

> float32 dipakai untuk mewakili semua jenis float : float32/float64

Ekspresi ini umum dijumpai , seperti menghitung domba saat mau tidur dan menambah hari libur.

Berikut contoh ekspresi arimatika pada golang:
| Tipe Data | Symbol | Arti | Contoh|
|--|--|--|--|
|`int`,`float32`,`string`  | `+` | penjumlahan(`int`&`float32`) concat(`string`)| `"a"+"b"`,`2+3`|
|`int`,`float32`| `-`|pengurangan| `2-1`,`2.1-0.5` |
|`int`,`float32`| `*`|perkalian| `2*1`,`3.14*22` |
|`int`,`float32`| `/`|pembagian| `2/1`,`2.3/4` |
|`int`| `%`|modulus| `2%1` |

### Expresi Relasional

Ekspresi ini melihat hubungan terhadap operan kiri dan operan kanan, ekspresi akan menghasil nilai bool dari hasil operasi.
Berbeda dengan ekspresi matematik yang merupakan hasil manipulasi dan perubahan dari data/nilai operan.

| Tipe Data | Symbol | Arti | Contoh|
|--|--|--|--|
|`int`,`float32`|`>`|Lebih besar dari|`2 > 1`|
|`int`,`float32`|`<`|Lebih kecil dari|`3 < 4` |
|`int`,`float32`|`>=`|Lebih besar dari sama dengan|`4 >= 5`|
|`int`,`float32`|`<=`|Lebih kecil dari sama dengan|` 6 <= 7` |
|`int`,`float32`,`string`|`==`|Sama dengan| `"Kamu" == "Aku"` |
|`int`,`float32`,`string`|`!=`|Tidak sama dengan| `"Mantan" != "Kenangan"`|

### Expresi Lojik
Expresi ini mengevaluasi kondisi lojik menjadi hasil berupa data lojik, umumnya digunakan untuk pengolahan keputusan.

| Tipe Data | Symbol | Arti | Contoh|
|--|--|--|--|
|`bool`| `&`|Logika Dan| `true && false` |
|`bool`| `|`|Logika Atau| `true || false` |
|`bool`| `!`|Negasi| `!true` |

Untuk mudah mengingat. Jika terdapat operan false pada logika dan (`&&`) maka akan munghasilkan false begitu sebaliknya dengan logika atau (`||`).
