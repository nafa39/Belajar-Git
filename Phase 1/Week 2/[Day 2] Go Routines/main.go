package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
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

// // Asynchronous email sending
// type Notification struct {
// 	UserID  int
// 	Message string
// }

// func sendEmailAsync(userID int, message string) {
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("Email notification sent to user %d: %s\n", userID, message)
// }

// // Images processing services
// func processImage(imageURL string) {
// 	fmt.Printf("Processing image: %s\n", imageURL)

// 	time.Sleep(3 * time.Second)

// 	fmt.Printf("Image processing completed: %s\n", imageURL)
// }

// // Task scheduler
// func scheduleTask(task func()) {
// 	go task()
// }

// download manager
func downloadFile(url string, destination string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading file from %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(destination)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", destination, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Error writing to file %s: %s\n", destination, err)
		return
	}
	fmt.Printf("Downloaded file from %s to %s\n", url, destination)
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

	// // Asynchronous email sending
	// now := time.Now()
	// notifications := []Notification{
	// 	{UserID: 101, Message: "Your order has been confirmed."},
	// 	{UserID: 202, Message: "Your account has been created."},
	// 	{UserID: 303, Message: "Your payment was successful."},
	// }

	// for _, notification := range notifications {
	// 	go sendEmailAsync(notification.UserID, notification.Message)
	// }

	// fmt.Println("Main application continous...")

	// time.Sleep(3 * time.Second)

	// fmt.Println("Main application finished.")
	// diff := time.Since(now)
	// fmt.Println(diff)

	// // Images processing services
	// now := time.Now()

	// images := []string{
	// 	"https://example.com/image1.jpg",
	// 	"https://example.com/image2.jpg",
	// 	"https://example.com/image3.jpg",
	// 	"https://example.com/image4.jpg",
	// }

	// for _, image := range images {
	// 	go processImage(image)
	// }

	// fmt.Println("Image processing started, main application continues...")

	// time.Sleep(5 * time.Second)

	// fmt.Println("All image processing completed.")
	// diff := time.Since(now)
	// fmt.Println(diff)

	// // task scheduler
	// now := time.Now()
	// task1 := func() {
	// 	fmt.Println("Task 1 is being executed.")
	// }

	// task2 := func() {
	// 	fmt.Println("Task 2 is being executed.")
	// }

	// task3 := func() {
	// 	fmt.Println("Task 3 is being executed.")
	// }

	// scheduleTask(task1)
	// scheduleTask(task2)
	// scheduleTask(task3)

	// fmt.Println("Main application continues...")

	// var wg sync.WaitGroup
	// wg.Add(3)
	// go func() {
	// 	wg.Wait()
	// 	fmt.Println("All task completed.")
	// }()

	// wg.Done()
	// wg.Done()
	// wg.Done()

	// diff := time.Since(now)
	// fmt.Println(diff)

	// //the output is different asyncron each other

	//download manager
	now := time.Now()
	downloadJobs := map[string]string{
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4": "video1.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_2mb.mp4": "video2.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_5mb.mp4": "video3.mp4",
	}

	var wg sync.WaitGroup
	wg.Add(len(downloadJobs))

	for url, destination := range downloadJobs {
		go func(u, d string) {
			defer wg.Done()
			downloadFile(u, d)
		}(url, destination)
	}

	wg.Wait()

	fmt.Println("All files downloaded.")

	diff := time.Since(now)
	fmt.Println(diff)

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

//time.Sleep(6 * time.Second)

//go routine to send email
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
