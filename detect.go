package main

import "strings"

func Flag(l Log) bool {
	// hardcode the conditions for now
	// p := "sshd"
	m := "Failed password"

	if strings.HasPrefix(l.Message, m) {
		return true
	}
	return false
}
