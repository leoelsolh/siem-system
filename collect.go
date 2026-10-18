package main

import (
	"bufio"
	"os"
)

func collectLogs(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
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
