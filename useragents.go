package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

type UAList struct {
	userAgents []string
}

func (l *UAList) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l.userAgents = append(l.userAgents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	log.Printf("%d user agents loaded from %s", len(l.userAgents), fileName)
	return nil
}

func (l *UAList) Next() string {
	return l.userAgents[rand.Int()%len(l.userAgents)]
}
