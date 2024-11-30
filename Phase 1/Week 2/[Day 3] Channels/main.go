package main

import (
	"fmt"
	"os"
)

//INCLASS

// //unbuffered channel
// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hai, my name is %s", student)

// 	c <- result
// }

// // producer and consumer
// func producer(c chan int) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("producing: %d\n", i)
// 		c <- i
// 		time.Sleep(time.Second * 1)
// 	}
// 	close(c)
// }

// func consumer(c chan int, done chan bool) {
// 	for msg := range c {
// 		fmt.Printf("consuming: %d\n", msg)
// 	}
// 	done <- true
// }

// // producer and consumer with pointer
// func producer(c chan *int) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("producing: %d\n", i)
// 		c <- &i
// 		time.Sleep(time.Second * 1)
// 	}
// 	close(c)
// }

// func consumer(c chan *int, done chan *bool) {
// 	for msg := range c {
// 		fmt.Printf("consuming: %d\n", *msg)
// 	}
// 	done <- new(bool)
// }

// // defer
// func deferFunc() {
// 	fmt.Println("closing file")
// }

// exit
func deferFunc() {
	fmt.Println("closing file")
}

func main() {

	//INCLASS

	// //unbuffered channel
	// c := make(chan string)

	// go introduce("Airell", c)
	// go introduce("Nanda", c)
	// go introduce("Mailo", c)

	// msg1 := <-c
	// fmt.Println(msg1)

	// msg2 := <-c
	// fmt.Println(msg2)

	// msg3 := <-c
	// fmt.Println(msg3)

	// close(c)

	// //buffered channel
	// cl := make(chan int, 3)

	// go func(c chan int) {
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Printf("func go routine %d before sending data into channel \n", i)
	// 		cl <- i
	// 		fmt.Printf("func go routine %d after sending data into channel \n", i)
	// 	}
	// 	close(c)
	// }(cl)

	// fmt.Println("main goroutine sleeping for 2 seconds")
	// time.Sleep(time.Second * 2)

	// for msg := range cl {
	// 	fmt.Println("main go routine received data from channel ", msg)
	// }

	// //producer and consumer
	// c := make(chan int)
	// done := make(chan bool)

	// go producer(c)
	// go consumer(c, done)

	// <-done
	// fmt.Println("done")

	// //producer and consumer with pointer
	// c := make(chan *int)
	// done := make(chan *bool)

	// go producer(c)
	// go consumer(c, done)

	// <-done
	// fmt.Println("done")

	// //defer
	// //defer fmt.Println("deferred")
	// defer deferFunc()
	// fmt.Println("not deferred")

	// exit
	defer func() {
		fmt.Println("deferred")
		deferFunc()
	}()
	fmt.Println("before exit")

	os.Exit(0) //force close the program

	fmt.Println("after exit")
}
