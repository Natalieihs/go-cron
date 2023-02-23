package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		fmt.Println("Task executed at", time.Now())
	})
	c.Start()

	defer c.Stop()

	select {}
}
