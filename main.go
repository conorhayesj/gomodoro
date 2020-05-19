package main

import (
	"fmt"
	"time"
	"flag"
)

func main() {

	minutePtr := flag.Int("m", 25, "Default 25 mins")
	taskPtr := flag.String("t", "No task", "Task to be completed")

	flag.Parse()

	fmt.Printf("Starting Pomodoro timer for %v mins.\n", *minutePtr)
	fmt.Printf("Task to be completed:\n%v\n", *taskPtr)
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("Times Up.")
}
