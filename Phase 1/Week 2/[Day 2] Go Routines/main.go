package main

import (
	"fmt"
	"sync"
)

// REVIEW MATERI
func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	fmt.Printf("Factorial of %d is %d\n", n, result)
}

func factorial1(n int) int {
	//defer wg.Done()

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	fmt.Printf("Factorial of %d is %d\n", n, result)
	return n
}

// //INCLASS
// func routine1() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("routine: ", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// // go routine send email

// func sendEmail() {
// 	for i := 0; i < 2; i++ {
// 		fmt.Println("Sending email: ", i)
// 		time.Sleep(3 * time.Second)
// 	}
// }

// func sendSMS() {
// 	for i := 0; i < 2; i++ {
// 		fmt.Println("Sending SMS: ", i)
// 		time.Sleep(2 * time.Second)
// 	}
// }

// func storeData() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Storing data: ", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }

func main() {

	//REVIEW MATERI
	numbers := []int{5, 7, 3, 10}

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for _, num := range numbers {
		go factorial(num, &wg)
	}

	wg.Wait()
	fmt.Println("All factorials calculated.")

	for _, num := range numbers {
		fact := factorial1(num)
		fmt.Printf("Calculated factorial %d\n", fact)
	}

	// //INCLASS
	// // start a go routine
	// go routine1()

	// //main routine
	// for i := 0; i < 2; i++ {
	// 	fmt.Println("main: ", i)
	// 	//time.Sleep(1 * time.Second)
	// }

	// time.Sleep(5 * time.Second)
	// fmt.Println("main function")

	// var wg sync.WaitGroup
	// wg.Add(2)

	// //go routine
	// go func() {
	// 	go sendEmail()
	// 	wg.Done()
	// }()

	// go func() {
	// 	go sendSMS()
	// 	wg.Done()
	// }()

	// wg.Wait()

	// storeData()

	// time.Sleep(6 * time.Second)
}

// go routine to send email
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func sendEmail() {
// 	for i := 0; i < 2; i++ {
// 		// SMTP code to send email
// 		time.Sleep(3 * time.Second)
// 		fmt.Println("Sending email: ", i)
// 	}
// }

// func sendSMS() {
// 	for i := 0; i < 2; i++ {
// 		// SMS code to send SMS
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("Sending SMS: ", i)
// 	}
// }

// func storeData() {
// 	for i := 0; i < 5; i++ {
// 		// Database code to store data
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Storing data: ", i)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go func() {
// 		sendEmail()
// 		wg.Done()
// 	}()

// 	go func() {
// 		sendSMS()
// 		wg.Done()
// 	}()

// 	wg.Wait()

// 	storeData()

// 	time.Sleep(6 * time.Second)
// }

// CATATAN
// asynchronous programming pasti delay dulu
