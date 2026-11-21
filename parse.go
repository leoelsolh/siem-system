package main

import (
	"log"
	"strconv"
	"strings"
)

type Log struct {
	Timestamp string
	Host      string
	Process   string
	PID       int
	Message   string
}

func parseLogs(lines []string) []Log {
	slice := []Log{}
	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 5 {
			log.Printf("Malformed log line on line: %d\n", i+1) // TODO:
			// Printing index here is unreliable in real logs, since line numbers change
			// change this to display line content instead
			continue
		}
		timestamp := strings.Join(fields[0:3], " ")
		host := fields[3]
		message := strings.Join(fields[5:], " ")
		process, rest, found := strings.Cut(fields[4], "[")

		var pidStr string
		var pid int

		if !found {
			process, _, _ = strings.Cut(fields[4], ":")
		} else {
			pidStr, _, _ = strings.Cut(rest, "]")
			pid, _ = strconv.Atoi(pidStr)
		}
		entry := Log{
			Timestamp: timestamp,
			Host:      host,
			Process:   process,
			PID:       pid,
			Message:   message,
		}
		slice = append(slice, entry)
	}
	return slice
}
