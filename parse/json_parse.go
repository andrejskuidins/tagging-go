package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parse(path string) *string {
	file, err := os.Open("")
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return nil
}

func main() {
	parse("../workfile.txt")
}
