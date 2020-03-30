# Variable

> Variable merupakan wadah yang kita deklarasikan untuk menampung nilai dengan tipe tertentu. Layaknya sebuah wadah, kita perlu mengetahui wadah yang tepat untuk menampung barang agar muat. Karena setiap wadah akan menyimpan dan dibaca sesuai data kebutuhan pengguna, hal ini juga memiliki pengaruh saat bagaimana nilai akan difungsikan.

## Deklarasi Variable
Go memiliki dua cara untuk mendeklarasikan variable : statik dan dinamik.

Perbedaanya ialah dekralasi tipe data secara ekspisit/implisit terhadap suatunilai.

Deklarsi **Statis** (eksplisit):

```go
var <nama> <tipe data> = <nilai>
```
Contoh:
```go
var  nama  string  =  "Elisabeth K"
```

dan ada cara untuk deklarasi secara **dinamis** : tipe data secara langsung ditentukan oleh tipe nilai yang diberikan. Deklarasi dinamis mengunakan _walrus operator_  (`:=`)

Deklarasi **dinamis**:
```go
<nama> := <nilai>
```
Contoh :
```go
umur := 12 // tipe data: int
```
> // merupakan komentar

#### Merubah nilai variable 
Dengan melakukan operator assignment (`=`) dengan nama variable bersangkutan dan nilai baru.
```go
umur := 12 // int
fmtPrintln(umur) // 12 
umur = 13
fmtPrintln(umur) // 13
```
dan menulis nama suatu variable untuk memanggil nilai dari suatu variable, cara kerja variable serupa dengan variable matematika : substitusi nama variable dengan nilai dalam variable.

**Contoh Penuh**
```go
var nama string = "Elisabeth K"
umur := 12
umur = 13
fmt.Println("Namaku",nama,"berumur",umur,"Tahun")
// Namaku Elisabeth K dan berumur 13
```

### Tipe Data
Tipe data merupakan _flag_(bendera) untuk menandakan tipe nilai yang akan ditampung. Apakah nilai tersebut bilangan, karakter , pointer atau tipe komposit, menentukan bagaimana suatu nilai untuk diberi tindakan.

Berikut beberapa contoh tipe data dasar :
Tipe |Nilai | 
-|-
`bool` | `true` dan `false`
`string` | kumpulan karakter
`int` | bilangan bulat : 1,2 ...
`float32` | bilangan pecahaan

### Deklarasi Paralel
Deklarasi Paralel merupakan deklarasi variable dengan banyak nama dan nilai dalam satu baris, namun dengan tipe sejenis. Setiap nama dan nilai dipisahkan dengan koma(`,`) dan dirangkai sesuai posisinya.

```go
var umur,tinggi int = 21,172
// umur = 21
// tinggi = 172
```
### Konstanta
Konstanta serupa variable , namun sekali variable diberi nilai maka nilai variable tersebut tidak bisa diubah.
```go
const <nama> <tipe data> = <nilai>
```

### Variable Pembuangan
Pada Go tidak boleh ada nilai/variable yang tidak digunakan, bila ada maka Go tidak akan jalan. Variable pembuangan biasanya ditemui untuk menangani fungsi yang mengembalikan nilai lebih dari satu. Deklarasi variable pembuangan dengan mengunakan nama variable berupa underscore (`_`).

```go
aku,_ := "Namaku El", "dan Aku Cantiqueee"
```
Kenapa variable `_` disebut pembuangan ? karena saat kita meletakan nilai didalamnya nilai tersebut tidak bisa di akses (nilainya dimakan), ini serupa kerja `/dev/null` di Linux.

### Pointer

[WIP]