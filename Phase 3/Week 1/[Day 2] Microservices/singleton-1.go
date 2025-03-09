package main

import (
	"fmt"
	"sync"
	"time"
)

type DB1 struct {
	Conn string
}

var singleDB1 *DB1
var mu sync.Mutex

func getInstance1() *DB1 {
	mu.Lock()
	defer mu.Unlock()
	if singleDB1 == nil {
		singleDB1 = &DB1{Conn: "postgresql:5432"} //instance
		fmt.Println("Database instance created")
	}
	return singleDB1
}

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(3000 * time.Millisecond)
			_ = getInstance1()
			fmt.Println("get instance", i)
		}
	}()
	go func() {
		fmt.Println("goroutine 2")
		_ = getInstance1()
	}()

	go func() {
		fmt.Println("goroutine 3")
		_ = getInstance1()
	}()

	fmt.Scanln()
}
