package main

import (
	"fmt"
	"log"
)

func main() {
	slice, err := collectLogs("logs/sample.log")
	if err != nil {
		log.Printf(
			"Something went wront while collecting logs:  %s",
			err,
		)
	}
	for i, thing := range slice {
		fmt.Printf("[%d]  %s\n", i+1, thing)
	}
}
