package main

import "strings"

func flagEntry(l Log) bool {
	// harcoded conditions are temporary
	p := "sshd"
	m := "Failed password"

	if strings.HasPrefix(l.Message, m) && l.Process == p {
		return true
	}
	return false
}
