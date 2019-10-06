# Bahasa Go

Golang atau juga dikenal Go adalah bahasa statically typed, compiled programming language yang dirancang pada 2007 oleh Robert Griesemer, Rob Pike dan Ken Thompson di Google dan Go diumumkan kepada publik pada November 2009.

> progress

## Persiapan

Hal yang paling penting sebelum belajar adalah mempersiapkan alat-alat yang dibutuhkan
sehingga proses belajar dapat berlangsung dengan baik, silahkan lewati bagian ini jika kamu
sudah mempunyai alat-alat yang dibutuhkan.

#### Text Editor dan Plugin

Kami menggunakan [Visual Studio Code](https://code.visualstudio.com/download) untuk menuliskan kode, silahkan menyesuaikan dengan patform OS masing-masing, dan kami menggunakan plugin [Go](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go) plugin ini sangat membantu penulisan kode Go nantinya.

#### Instalasi Go

Sebelum kita memulai belajar Go tentu kita harus meng-install terlebih dahulu, agar kita bisa menggunakanya pada komputer kita. Panduan penginstallan sebenarnya sudah tersedia pada [situs resmi](https://golang.org/doc/install) golang, namun disini akan memberikan panduan penginstallan yang singkat, jika ingin mengetahui panduan lengkapnya kamu dapat melihat disitus resminya, pada buku ini kita menggunakan versi `go1.13.1`.

#### Linux

1. Buka terminal kamu lalu ketikan perintah berikut pada terminalmu.
   `$ sudo tar -C /usr/local -xzf go1.13.1.linux-amd64.tar.gz`
   Atau menyesuaikan dengan sistem operasi yang kamu gunakan 32bit atau memiliki arch berbeda dengan kami, maka bisa menggunakan perintah ini.
   `$ tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`
   Dimana `$VERSION=1.13.1 $OS=linux $ARCH=amd64 atau sesuaikan architecture masing-masing cth: amd64, 386, arm, arm64, s390x, ppc64le`

2. Setelah melalukan perintah diatas kita harus mengekspor path-nya, dengan menggunakan perintah dibawah ini.
   `$ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.zshrc`
   `$ source ~/.zshrc # Mengeksekusi file tanpa harus merestart terminal`
   atau sesuaikan dengan `.profile` kamu sendiri jika kamu menggunakan bash terminal kamu bisa menggunakan `~/.bashrc` sebagai gantinya.

3. Selanjutnya langkah terakhir adalah memastikan bahwa Go sudah terinstall dengan benar dengan menggunakan perintah `$ go version` output yang diharapkan adalah `$ go version go1.13.1`

#### MacOS

#### Windows
