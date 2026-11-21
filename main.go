package main

import (
	"fmt"
	"log"
)

func main() {
	l, err := collectLogs("logs/sample.log")
	if err != nil {
		log.Printf("Error while collecting logs:  %v", err)
		return
	}
	p := parseLogs(l)
	flagged := filterEvents(p)
	fmt.Println("[ flagged events ]")
	for _, f := range flagged {
		fmt.Println(f.Message)
	}
}
