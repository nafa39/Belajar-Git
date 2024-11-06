package main

import (
	"fmt"
	"learn-pointer/helpers"
)

// Pointer adalah sebuah variabel yang menyimpan alamat memori dari sebuah variabel lain
// Pointer digunakan untuk mengakses data yang disimpan di alamat memori yang sama
// pointer pakai function

// INCLASS EXERCISE

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

//INDIVIDUAL EXERCISE

// SumInt sums up all integers in a slice.
// func SumInt(numbers []int) int {
// 	sum := 0
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	return sum
// }

//GROUP EXERCISE

//NON GRADED CHALLENGE

func main() {

	// pegawai := helpers.TambahPegawai(1, "Andi", 25, 5000000.0)

	// fmt.Println(pegawai)
	// pegawai.NaikkanGaji(10)

	// p := helpers.CariPegawai(1)
	// fmt.Println(p)
	// fmt.Println(p.Gaji)

	//INCLASS EXERCISE

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

	//INDIVIDUAL EXERCISE

	// fmt.Println("Individu Exercise")
	// fmt.Println(" ")

	// fmt.Println("Hello Nafa")

	// //string manipulation
	// var name string = " john doe "
	// fmt.Println(name)

	// //uppercase
	// fmt.Println(strings.ToUpper(name))

	// //titlecase
	// fmt.Println(strings.Title(name))

	// //lowercase
	// fmt.Println(strings.ToLower(name))

	// //trim
	// fmt.Print("Trim:", name, "Something")
	// fmt.Println(" ")
	// fmt.Print("Trim:", strings.TrimRight(name, " "), "Something")
	// fmt.Println(" ")
	// fmt.Print("Trim:", strings.TrimLeft(name, " "), "Something")
	// fmt.Println(" ")

	// nama := "NAFATUL"

	// result := strings.Split(nama, "")
	// fmt.Println(result)

	// //sum int
	// result2 := SumInt([]int{1, 2, 3, 4, 5})

	// fmt.Println(result2)

	// //GROUP EXERCISE

	// fmt.Println(" ")
	// fmt.Println("Group Exercise")
	// fmt.Println(" ")

	// var emp = helpers.Pegawai{
	// 	ID:   1,
	// 	Nama: "Andi",
	// 	Umur: 25,
	// 	Gaji: 5000000.0,
	// }

	// //print data pegawai sebelum di ubah
	// fmt.Println(emp.ID, emp.Nama, emp.Umur)
	// fmt.Printf("Gaji sebelum naik: %.2f\n", emp.Gaji)

	// //print data pegawai setelah di ubah
	// helpers.Ubah(&emp)
	// fmt.Println(emp.ID, emp.Nama, emp.Umur)
	// fmt.Printf("Gaji setelah naik: %.2f\n", emp.Gaji)

	//NON GRADED CHALLENGE

	//struct and method

	var person = helpers.Person{
		Name: "nafa",
		Age:  45,
		Job:  "free style",
	}

	//print data person sebelum diubah
	fmt.Println(person)

	for i := 0; i < 5; i++ {
		helpers.AddYear(&person)
	}
	fmt.Println(person)

	//struct & method 2
	var org = []helpers.Person{
		{Name: "dicki", Age: 51, Job: "anak tiri"},
		{Name: "billa", Age: 8, Job: "sparepart"},
		{Name: "sulhan", Age: 9, Job: "mekanik"},
	}

	for _, o := range org {
		helpers.GetInfo(&o)
	}
}
