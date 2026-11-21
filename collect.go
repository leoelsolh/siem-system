package main

import (
	"bufio"
	"os"
)

func collectLogs(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	slice := []string{}
	for s.Scan() {
		slice = append(slice, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return slice, nil
}
