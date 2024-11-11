// handling error in golang
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// type MyError struct {
// 	Msg string
// 	Err error
// 	Code int
// }

func insertData(nama string, umur int) (inserted_id int, err error) {
	// insert data to database
	// if success return inserted_id
	// if failed return error
	fmt.Println(nama, umur)
	return 0, errors.New("Error") // ini belom ngerti
}

func main() {
	var str1 string = "123"
	var str2 string = "123.123"

	var int1 int
	var int2 int

	var err error

	_, err = insertData("John Doe", 20)
	if err != nil {
		fmt.Println(err)
	}

	// Convert string to int
	int1, err = strconv.Atoi(str1)
	if err != nil {
		fmt.Println(err)
		// } else {
		// 	fmt.Println(int1)
	}

	// Convert string to int
	int2, err = strconv.Atoi(str2)
	if err != nil {
		fmt.Println(err)
		// } else {
		// 	fmt.Println(int2)
	}

	fmt.Println(int1)
	fmt.Println(int2)

}

//case kak dharma
// 	package main
// import (
//     "errors"
//     "fmt"
// )
// func division(n1, n2 int) (int, error) {
//     //validasi input kalo 0
//     if n2 == 0 {
//         return 0, errors.New("cannot divide by zero")
//     }
//     //bagiin
//     result := n1 / n2
//     //balikin result
//     return result, nil
// }
// func main() {
//     n1 := 1
//     n2 := 0
//     //manggil sama bikin variable error
//     _, err := division(n1, n2)
//     //kalo error ngapain kalo success ngapain
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Success")
//     }
// }

//contoh kak dharma lagi
// package main
// import (
//     "errors"
//     "fmt"
// )
// func division(n1, n2 int) (int, error) {
//     //validasi input kalo 0
//     if n2 == 0 {
//         return 0, errors.New("cannot divide by zero")
//     }
//     //bagiin
//     result := n1 / n2
//     //balikin result
//     return result, nil
// }
// func main() {
//     n1 := 1
//     n2 := 0
//     //manggil sama bikin variable error
//     _, err := division(n1, n2)
//     //kalo error ngapain kalo success ngapain
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Success")
//     }
// }

//recover belom aman, defer belom ngerti, fmt.Errorf belom ngerti
//defer func() // kalo panic tetep jalan tapi ga masuk ke recover
//if recover() != nil { => ini buat ngambil errornya
