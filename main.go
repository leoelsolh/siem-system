package main

import (
	"fmt"
	"log"
)

func main() {
	l, err := collectLogs("logs/sample.log")
	if err != nil {
		log.Printf("Error while collecting logs:  %s", err)
		return
	}
	p := parseLogs(l)
	for i, entry := range p {
		found := Flag(entry)
		if found {
			fmt.Printf("Alert on line %d\n", i+1)
		}
	}
}
