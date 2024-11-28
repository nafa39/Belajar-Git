package main

import (
	"fmt"
	"time"
)

// REVIEW MATERI

// // hands-on-lab
// func factorial(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	result := 1
// 	for i := 2; i <= n; i++ {
// 		result *= i
// 	}

// 	fmt.Printf("Factorial of %d is %d\n", n, result)
// }

// func factorial1(n int) int {
// 	//defer wg.Done()

// 	result := 1
// 	for i := 2; i <= n; i++ {
// 		result *= i
// 	}

// 	fmt.Printf("Factorial of %d is %d\n", n, result)
// 	return n
// }

// // goroutines
// func goroutine() {
// 	fmt.Println("Hello")
// }

// // goroutines (Asynchronous process #1)
// func firstProcess(index int) {
// 	fmt.Println("First process func started")
// 	for i := 1; i <= index; i++ {
// 		fmt.Println("i= ", i)
// 	}
// 	fmt.Println("First process func ended")
// }

// func secondProcess(index int) {
// 	fmt.Println("Second process func started")
// 	for j := 1; j <= index; j++ {
// 		fmt.Println("j= ", j)
// 	}
// 	fmt.Println("Second process func ended")
// }

// Asynchronous email sending
type Notification struct {
	UserID  int
	Message string
}

func sendEmailAsync(userID int, message string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Email notification sent to user %d: %s\n", userID, message)
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

	// //hands-on-lab
	// numbers := []int{5, 7, 3, 10}

	// var wg sync.WaitGroup
	// wg.Add(len(numbers))

	// for _, num := range numbers {
	// 	go factorial(num, &wg)
	// }

	// wg.Wait()
	// fmt.Println("All factorials calculated.")

	// for _, num := range numbers {
	// 	fact := factorial1(num)
	// 	fmt.Printf("Calculated factorial %d\n", fact)
	// }

	// //go routines
	// go goroutine()

	// //goroutines (Asynchronous process #1)
	// fmt.Println("main execution started")

	// go firstProcess(8)

	// secondProcess(8)

	// fmt.Println("No. of Goroutines: ", runtime.NumGoroutine())

	// //goroutines (Asynchronous process #3)
	// time.Sleep(time.Second * 2)

	// fmt.Println("Main execution ended")

	// // RUNTIME
	// //synchronous
	// now := time.Now()

	// for i := 0; i < 10; i++ {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("Helloworld")
	// }

	// diff := time.Since(now)
	// fmt.Println(diff)

	// //Asynchronous
	// wg := &sync.WaitGroup{}
	// now := time.Now()

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println("Helloworld")
	// 		wg.Done()
	// 	}()
	// }

	// wg.Wait()
	// diff := time.Since(now)
	// fmt.Println(diff)

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
