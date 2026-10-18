package main

import (
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
	for _, line := range lines {
		fields := strings.Fields(line)

		timestamp := strings.Join(fields[0:3], " ")
		host := fields[3]
		message := strings.Join(fields[5:], " ")
		process, rest, found := strings.Cut(fields[6], "[")

		var pidStr string
		var pid int

		if !found {
			_, process, _ = strings.Cut(fields[4], ":")
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
