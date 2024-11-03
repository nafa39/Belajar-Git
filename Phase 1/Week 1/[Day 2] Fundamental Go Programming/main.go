package main

import (
	"fmt"
)

// NON GRADED CHALLENGE LOOPING 2
// make median function
// func median(data []float64) float64 {
// 	dataCopy := make([]float64, len(data))
// 	copy(dataCopy, data)

// 	sort.Float64s(dataCopy)

// 	var median float64
// 	l := len(dataCopy)
// 	if l == 0 {
// 		return 0
// 	} else if l%2 == 0 {
// 		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
// 	} else {
// 		median = dataCopy[l/2]
// 	}

// 	return median
// }

// // NON GRADED CHALLENGE LOGIC 1 - PALINDROME
// func reverse(word string) string {
// 	runes := []rune(word)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// // NON GRADED CHALLENGE FUNCTION 1
// func AlayGen(data2 []string) string {
// 	var word string
// 	for i := range data2 {
// 		runes := []rune(data2[i])
// 		for j := 0; j < len(runes); j++ {
// 			runes[j] = []rune(strings.Replace(string(runes[j]), "a", "4", -1))[0]
// 			runes[j] = []rune(strings.Replace(string(runes[j]), "i", "!", -1))[0]
// 			runes[j] = []rune(strings.Replace(string(runes[j]), "e", "3", -1))[0]
// 			runes[j] = []rune(strings.Replace(string(runes[j]), "l", "1", -1))[0]
// 			runes[j] = []rune(strings.Replace(string(runes[j]), "s", "5", -1))[0]
// 			runes[j] = []rune(strings.Replace(string(runes[j]), "n", "N", -1))[0]
// 			runes[j] = []rune(strings.Replace(string(runes[j]), "x", "*", -1))[0]
// 		}
// 		data2[i] = string(runes)
// 		word += data2[i] + " "
// 	}
// 	return word
// }

// NON GRADED CHALLENGE FUNCTION 2
func fibonacci(n int) int {
	var result int
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			result = i
		} else {
			result = fibonacci(n-1) + fibonacci(n-2)
		}
	}
	return result
}

// NON GRADED CHALLENGE TRIBONACCI
func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	}
}

func main() {

	//INCLASS DAY 2

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

	//var student1, student2, student3 string = "satu", "dua", "tiga"
	//var first, second, third int

	//first, second, third = 1, 2, 3

	//fmt.Println(student1, student2, student3)
	//fmt.Println(first, second, third)

	//var name, age, address = "Airell", 23, "Jakarta"
	//first, second, third := "1", 2, "3"

	//fmt.Println(name, age, address)
	//fmt.Println(first, second, third)

	//fmt.Println("nama saya >", name)   // sama seperti console.log ("nama saya >" + name)
	//fmt.Printf("nama saya > %s", name) //di dalam string
	//fmt.Println("\n")

	//variable underscore

	// 	_ = "nafa"

	// 	hasilPenjumlahan, _ := penjumlahan()
	// 	fmt.Println(hasilPenjumlahan)

	// 	var first = 89
	// 	var second = -12

	// 	fmt.Printf("tipe data first = %T\n", first)
	// 	fmt.Printf("bilangan second = %T\n", second)

	// 	decimal := 3.63
	// 	fmt.Printf("decimal = %.3f\n", decimal) // set 3 digit di belakang koma

	// 	var message = `selamat datang di "Hacktiv8".
	// Salam Kenal:).
	// Mari belajar "Scalable web service with go".`

	// 	fmt.Println(message)

	//konstanta
	//const nama = "Airell"

	// var amount = 6 / 2 * (1 + 2)

	// fmt.Println(amount)

	//condition

	// var currentYear = 2024

	// if age := currentYear - 1998; age < 18 {
	// 	fmt.Println("Underage")
	// } else {
	// 	fmt.Println("Adult")
	// }

	// var score = 6

	// switch {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// 	fallthrough
	// case (score < 8) && (score > 3):
	// 	fmt.Println("Awesome")
	// case score < 5:
	// 	fmt.Println("You need to learn more")
	// default:
	// 	fmt.Println("Not bad")
	// 	fmt.Println("Good job")
	// }

	// //looping hal 17
	// for i := 1; i <= 20; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("angka", i)
	// }

	//EXERCISE DAY 2
	//PR
	//soal 1

	//length string pakai len(string)
	//T
	//UG
	//ASM

	//nanti diganti pake *
	//AS*

	//soal 3

	//bikin tribonacci setelah fibonacci

	//hal 20 outerloop

	// var name = "NAFATULMIARAHMAWATI"
	// for i := 1; len(name) > 0; i++ {
	// 	if i > len(name) {
	// 		// When remaining characters are less than i, print the rest with asterisks
	// 		fmt.Println(name, strings.Replace(fmt.Sprintf("%*s", i-len(name), ""), " ", "*", -1))
	// 		//fmt.Println(name, "*")
	// 		break
	// 	}
	// 	fmt.Println(name[:i])
	// 	name = name[i:]
	// }

	//soal 2

	//make the pyramids like this for the output if the input is 5

	//00000
	//0000
	//000
	//00
	//0

	// a := "0"
	// input := 7
	// for i := input; i > 0; i-- {
	// 	fmt.Println(strings.Repeat(a, i))
	// }

	// //NON GRADED CHALLENGE VARIABLE 1
	// var myNum int32 = 50
	// fmt.Println(myNum)

	// var myNum2 float32 = 51
	// fmt.Println(myNum2)

	// var myNumStr string = "50"
	// fmt.Println(myNumStr)

	// //NON GRADED CHALLENGE VARIABLE 2
	// var x int32 = 5
	// var y int32 = 10

	// z := x + y
	// fmt.Println(z)

	// //NON GRADED CHALLENGE CLI
	// fmt.Println("Masukan Nama:")
	// var nama string
	// fmt.Scanln(&nama)
	// fmt.Println("Halo", nama)

	// //NON GRADED CHALLENGE ARRAY & SLICE 1
	// var people = []string{"Walt", "Jesse", "Skyler", "Saul"}
	// fmt.Println(len(people))
	// fmt.Println(people)
	// people = append(people, "Hank", "Marry")
	// fmt.Println(len(people))
	// fmt.Println(people)

	// //NON GRADED CHALLENGE ARRAY & SLICE 2
	// var data []map[string]string
	// data = append(data, map[string]string{"name": "Hank", "gender": "M"})
	// data = append(data, map[string]string{"name": "Heisenberg", "gender": "M"})
	// data = append(data, map[string]string{"name": "Skyler", "gender": "F"})

	// fmt.Println(data)

	// for i := range data {
	// 	if data[i]["gender"] == "M" {
	// 		data[i]["name"] = "Mr." + data[i]["name"]
	// 	} else if data[i]["gender"] == "F" {
	// 		data[i]["name"] = "Mrs." + data[i]["name"]

	// 	}
	// }
	// fmt.Println(data)

	// //NON GRADED CHALLENGE CONDITIONAL 1
	// var Nama string
	// fmt.Println("Masukan Nama:")
	// fmt.Scanln(&Nama)
	// var RandomNumber int = rand.Intn(100)
	// switch {
	// case RandomNumber > 80:
	// 	fmt.Println("Selamat ", Nama, ", anda sangat beruntung")
	// case RandomNumber > 60:
	// 	fmt.Println("Selamat ", Nama, ", anda beruntung")
	// case RandomNumber > 40:
	// 	fmt.Println("Mohon maaf ", Nama, ", anda kurang beruntung")
	// default:
	// 	fmt.Println("Mohon maaf ", Nama, ", anda sial")
	// }
	// fmt.Println("Random Number:", RandomNumber)

	// //NON GRADED CHALLENGE CONDITIONAL 2
	// var nama2 string
	// fmt.Println("Masukan Nama:")
	// fmt.Scanln(&nama2)
	// var umur int
	// fmt.Println("Masukan Umur:")
	// fmt.Scanln(&umur)

	// if _, err := fmt.Sscanf(fmt.Sprint(umur), "%d", &umur); err != nil {
	// 	fmt.Println("Umur harus berupa angka")
	// 	return
	// }

	// if umur < 0 {
	// 	fmt.Println("Umur tidak boleh negatif")
	// } else if umur > 100 {
	// 	fmt.Println("Umur maksimal 100 tahun")
	// } else if umur > 18 {
	// 	fmt.Println("Silahkan masuk")
	// } else {
	// 	fmt.Println("Dilarang masuk, maksimal umur 19 tahun")
	// }

	// //NON GRADED CHALLENGE LOOPING 1
	// var data2 = []map[string]string{}
	// data2 = append(data2, map[string]string{"name": "Hank", "age": "50", "Job": "Polisi"})
	// data2 = append(data2, map[string]string{"name": "Heisenberg", "age": "52", "Job": "Ilmuwan"})
	// data2 = append(data2, map[string]string{"name": "Skyler", "age": "48", "Job": "Akuntan"})

	// for i := range data2 {
	// 	fmt.Println("Hi perkenalkan, Nama saya ", data2[i]["name"], ", umur saya ", data2[i]["age"], " tahun, dan saya bekerja sebegai ", data2[i]["Job"])
	// }

	// //NON GRADED CHALLENGE LOOPING 2
	// var slice1 = []float64{1, 5, 7, 8, 10, 24, 33}
	// var slice2 = []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	// var sum, sum2 float64
	// for _, value := range slice1 {
	// 	sum += value
	// }

	// for _, value := range slice2 {
	// 	sum2 += value
	// }

	// var average float64 = sum / float64(len(slice1))
	// fmt.Println("Rata-rata slice1 adalah", average)
	// var average2 float64 = sum2 / float64(len(slice2))
	// fmt.Println("Rata-rata slice2 adalah", average2)
	// fmt.Println("Jumlah slice1 adalah", sum)
	// fmt.Println("Jumlah slice2 adalah", sum2)
	// fmt.Println("Median slice1 adalah", median(slice1))
	// fmt.Println("Median slice2 adalah", median(slice2))

	// //NON GRADED CHALLENGE LOGIC 1 - PALINDROME
	// var word string
	// fmt.Println("Masukan kata palindrome:")
	// fmt.Scanln(&word)
	// if word == reverse(word) {
	// 	fmt.Println("Kata", word, "adalah palindrome")
	// } else {
	// 	fmt.Println("Kata", word, "bukan palindrome")
	// }

	// //NON GRADED CHALLENGE LOGIC 2 - XOXO
	// var word2 string
	// fmt.Println("Masukan kata:")
	// fmt.Scanln(&word2)
	// var countX, countO int
	// for _, value := range word2 {
	// 	if value == 'x' {
	// 		countX++
	// 	} else if value == 'o' {
	// 		countO++
	// 	}
	// }
	// if countX == countO {
	// 	fmt.Println("True")
	// } else {
	// 	fmt.Println("False")
	// }

	// //NON GRADED CHALLENGE LOGIC 3 - XOXO
	// var swap = []int{12, 1, 3, 4, 5}
	// for i := 0; i < len(swap)-1; i++ {
	// 	if swap[i] > swap[i+1] {
	// 		swap[i], swap[i+1] = swap[i+1], swap[i]
	// 	}
	// }
	// fmt.Println(swap)

	// //NON GRADED CHALLENGE FUNCTION 1
	// var data3 = []string{"what", "personal", "morality", "complex", "judgements", "this", "is", "a", "test"}
	// fmt.Println(AlayGen(data3))

	//NON GRADED CHALLENGE FUNCTION 2
	var data4 int
	fmt.Println("Masukan angka:")
	fmt.Scanln(&data4)
	fmt.Println(fibonacci(data4))

	//NON GRADED CHALLENGE TRIBONACCI
	var data5 int
	fmt.Println("Masukan angka:")
	fmt.Scanln(&data5)
	fmt.Println("Tribonacci ke-", data5, "adalah", tribonacci(data5))

}

//CATATAN
//harus mereturn error

// func penjumlahan() (int, error) {
// 	//proses logika penjumlahan
// 	return 5, nil
// }

// func coba(x, y int) (int, int) {
// 	return 100, 50
// }
