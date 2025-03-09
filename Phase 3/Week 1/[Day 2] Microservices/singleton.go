package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	var once sync.Once

	task := func() {
		fmt.Println("executing task")
	}

	goroutine(&once, task)
	goroutine(&once, task)

	time.Sleep(1)
}

func goroutine(once *sync.Once, onceBody func()) {
	once.Do(onceBody)
}
