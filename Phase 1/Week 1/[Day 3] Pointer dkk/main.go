package main

import (
	"fmt"
	"learn-pointer/helpers"
)

// Pointer adalah sebuah variabel yang menyimpan alamat memori dari sebuah variabel lain
// Pointer digunakan untuk mengakses data yang disimpan di alamat memori yang sama
// pointer pakai function

// func change(name *string) {
// 	*name = "budiman"
// }

// func change2(name string) {
// 	name = "budiman"
// }

//latihan soal

//Sistem Data Pegawai
// Buatlah program yang memiliki struktur data untuk menyimpan informasi pegawai. Setiap pegawai memiliki informasi berikut:

// ID (integer)
// Nama (string)
// Umur (integer)
// Gaji (float64)

// Implementasikan fungsi berikut:

// Buat fungsi TambahPegawai untuk menambahkan data pegawai baru ke dalam slice pegawai.
// Buat metode NaikkanGaji yang menerima persentase kenaikan gaji untuk pegawai tertentu berdasarkan ID. Gunakan pointer receiver untuk memperbarui gaji pegawai.
// Buat fungsi CariPegawai yang menerima ID pegawai dan mengembalikan pointer ke struct pegawai yang sesuai. Jika tidak ditemukan, kembalikan nil.

// output

// Contoh waktu panggil functionnya :

// pegawai := TambahPegawai(1, "Andi", 25, 5000000.0)
// pegawai.NaikkanGaji(10)
// p := CariPegawai(1)
// fmt.Println(p.Gaji) // Harus mencetak gaji setelah kenaikan

func main() {

	pegawai := helpers.TambahPegawai(1, "Andi", 25, 5000000.0)

	fmt.Println(pegawai)
	pegawai.NaikkanGaji(10)

	p := helpers.CariPegawai(1)
	fmt.Println(p)
	fmt.Println(p.Gaji)

	// var name string = "John Doe"

	// // Pointer
	// var nama2 *string = &name //& ngambil alamat dari variabel name
	// //fmt.Println("Hello, my name is ", name)  // John Doe
	// //fmt.Println("Hello, my name is ", &name) // 0xc000010200
	// //fmt.Println("Hello, my name is ", nama2) // 0xc000010200
	// //fmt.Println("Hello, my name is ", *nama2)

	// //*nama2 = "Doe John" //mangubah string yang ada di alamat nama2
	// change(nama2)
	// //change2(*nama2)                          //tidak berubah karena tidak menggunakan pointer, panggil function change2 pakai *nama2
	// fmt.Println("Hello, my name is ", &name) // Doe John

	// //& untuk mengambil alamat memori yang bukan pointer
	// //* untuk mengambil data yang ada di alamat memori yang sudah diambil oleh pointer

	// //struct

	// type Address struct {
	// 	City, Province, Country string
	// }

	// //embedded struct
	// type Person struct {
	// 	Name    string
	// 	Address Address
	// }

	// var dia Address = Address{
	// 	City:     "Subang",
	// 	Province: "Jawa Barat",
	// 	Country:  "Indonesia",
	// }

	// fmt.Println(dia)
	// fmt.Println(dia.City)

	// //slice of struct

	// var allPerson = []Person{
	// 	Person{
	// 		Name: "John Doe",
	// 		Address: Address{
	// 			City:     "Subang",
	// 			Province: "Jawa Barat",
	// 			Country:  "Indonesia",
	// 		},
	// 	},
	// 	Person{
	// 		Name: "John Doe",
	// 		Address: Address{
	// 			City:     "Subang",
	// 			Province: "Jawa Barat",
	// 			Country:  "Indonesia",
	// 		},
	// 	},
	// }

	// fmt.Println(allPerson)
}
