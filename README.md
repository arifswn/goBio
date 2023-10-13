# goBio

<!-- penjelasan -->

## BioApp adalah mini project dibuat untuk pembelajaran dan latihan, Golang sebagai Backend dan MySQL sebagai Database, Flutter sebagai Mobile App.

> REST API sederhana untuk menampilkan data biodata yang diambil dari database MySQL, menggunakan bahasa pemrograman Go dan framework Echo.

> Mobile App sederhana untuk menampilkan data biodata yang diambil dari REST API, menggunakan bahasa pemrograman Dart dan framework Flutter.

<!-- step by step -->

### Step 1: Installasi Go

1. Download Go di https://golang.org/dl/
2. Install Go
3. Cek versi Go dengan perintah `go version`
4. Cek apakah Go sudah terinstall dengan perintah `go env`
5. Buat folder untuk project Go, misalnya `goBio`

### Step 2: Installasi Library

1. Buka terminal
2. Masuk ke folder project Go, misalnya `cd goBio`
3. Install library GORM dengan perintah `go get gorm.io/gorm gorm.io/driver/mysql`
4. Install library Echo dengan perintah `go get github.com/labstack/echo/v4`
5. Cek apakah library sudah terinstall dengan perintah `go list -m all`
   <!-- note, untuk menuju ke panduan createGo -->
   > Untuk panduan lengkap membuat project Go, bisa dilihat di [sini](/panduan/createGo.md)

<!-- link ke panduan createDB -->

### Step 3: Buat database MySQL

1. Buka terminal
2. Masuk ke folder project Go, misalnya `cd goBio`
3. Buat database MySQL dengan perintah `mysql -u root -p`
4. Masukkan password MySQL
5. Buat database dengan perintah `create database go_bio`
6. Keluar dari MySQL dengan perintah `exit`
   <!-- note, untuk menuju ke panduan createDB -->
   > Untuk panduan lengkap membuat database MySQL, bisa dilihat di [sini](/panduan/createDB.md)

<!-- support and contact -->

## Support and Contact

<!-- talk about me -->

<!-- create smile and coffe -->
> Jika kalian merasa terbantu dengan adanya project ini, kalian bisa memberikan dukungan dengan cara memberikan bintang ⭐ pada project ini. Terima kasih atas dukungannya.

> Coffee ☕ adalah minuman yang sangat membantu saya dalam membuat project ini. Jika kalian ingin membelikan saya kopi, bisa melalui [sini](https://saweria.co/arifswn).

> Jika ada pertanyaan, saran, atau masukan, bisa menghubungi saya melalui email di [sini](mailto:hexabiner808@gmail.com) atau melalui [LinkedIn](https://www.linkedin.com/in/arifswn/).

<!-- license -->

## License

[MIT](https://choosealicense.com/licenses/mit/) License adalah lisensi perangkat lunak bebas yang sangat sederhana. Lisensi ini memberikan kepada penerima perangkat lunak hak untuk menggunakan perangkat lunak untuk tujuan apa pun, baik komersial maupun non-komersial, tanpa dikenakan biaya atau tanggung jawab apa pun. Perangkat lunak yang dilisensikan di bawah MIT License dapat disalin, dimodifikasi, diterbitkan, didistribusikan, digabungkan dengan perangkat lunak lain, dan/atau dijual dengan tanpa batasan.
