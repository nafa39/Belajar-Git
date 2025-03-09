package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("59 0 * * *", func() {
		fmt.Println("Every day at", time.Now())
	})
	c.Start()
	select {}
}
