package main

import (
	"fmt"
	"time"
)

type DB struct {
	Conn string
}

var singleDB *DB

func getInstance() *DB {
	if singleDB == nil {
		singleDB = &DB{Conn: "postgresql:5432"} //instance
		fmt.Println("Database instance created")
	}
	return singleDB
}

func main1() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(3000 * time.Millisecond)
			_ = getInstance()
			fmt.Println("get instance", i)
		}
	}()
	go func() {
		fmt.Println("goroutine 2")
		_ = getInstance()
	}()

	fmt.Scanln()
}
