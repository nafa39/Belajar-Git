package main

import "fmt"

func main() {
	//var name string = "Airell"
	//var age int = 23

	//fmt.Println("Ini namanya adalah ==>", name)
	//fmt.Println("Ini adalah umurnya ==>", age)

	//var age = 23
	//var name = "Airell"

	//fmt.Printf("Ini namanya adalah ==> %s\n", name)
	//fmt.Printf("%T, %T", name, age)

	//name := "Airell"
	//age := 23

	//fmt.Printf("Ini namanya adalah ==> %s\n", name)
	//fmt.Printf("%T, %T", name, age)
	//fmt.Println("\n")

	var student1, student2, student3 string = "satu", "dua", "tiga"
	//var first, second, third int

	//first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	//fmt.Println(first, second, third)

	//var name, age, address = "Airell", 23, "Jakarta"
	//first, second, third := "1", 2, "3"

	//fmt.Println(name, age, address)
	//fmt.Println(first, second, third)

	//fmt.Println("nama saya >", name)   // sama seperti console.log ("nama saya >" + name)
	//fmt.Printf("nama saya > %s", name) //di dalam string
	//fmt.Println("\n")

	//variable underscore

	_ = "nafa"

	hasilPenjumlahan, _ := penjumlahan()
	fmt.Println(hasilPenjumlahan)

	var first = 89
	var second = -12

	fmt.Printf("tipe data first = %T\n", first)
	fmt.Printf("bilangan second = %T\n", second)

	decimal := 3.63
	fmt.Printf("decimal = %.3f\n", decimal) // set 3 digit di belakang koma

	var message = `selamat datang di "Hacktiv8".
Salam Kenal:).
Mari belajar "Scalable web service with go".`

	fmt.Println(message)

	//konstanta
	//const nama = "Airell"

	var amount = 6 / 2 * (1 + 2)

	fmt.Println(amount)

	//condition

	var currentYear = 2024

	if age := currentYear - 1998; age < 18 {
		fmt.Println("Underage")
	} else {
		fmt.Println("Adult")
	}

	var score = 6

	switch {
	case score == 8:
		fmt.Println("Perfect")
		fallthrough
	case (score < 8) && (score > 3):
		fmt.Println("Awesome")
	case score < 5:
		fmt.Println("You need to learn more")
	default:
		fmt.Println("Not bad")
		fmt.Println("Good job")
	}

	//looping hal 17
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("angka", i)
	}

	//PR
	//soal 1

	//length string pakai len(string)
	//T
	//UG
	//ASM

	//nanti diganti pake *
	//AS*

	//soal 2

	//00000
	//0000
	//000
	//00
	//0

	//soal 3

	//bikin tribonacci setelah fibonacci

	//hal 20 outerloop

	//punya dharma
	var name = "DHARMA"
	for i := 1; len(name) > 0; i++ {
		fmt.Println(name[:i])
		name = name[i:]
	}

}

//CATATAN
//harus mereturn error

func penjumlahan() (int, error) {
	//proses logika penjumlahan
	return 5, nil
}

func coba(x, y int) (int, int) {
	return 100, 50
}

//PR jawaban obie
// func printWord(word string) (string, error){
// 	var wordPyramid = "";

// 	for i := 0; i < len(word); i++ {
// 		for j:= 0; j <= i; j++{
// 			wordPyramid += string(word[j]) + " "
// 		}

// 		if (i != len(word)-1){
// 			wordPyramid += "\n"
// 		}
// 	}

// 	return wordPyramid, nil
// }
// // ``

// // usage di main functionnya:

// // var triangle, _ = printWord("Tugas");
// // fmt.Println(triangle)
