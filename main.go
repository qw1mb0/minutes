package main

import (
	"bufio"
	"fmt"
	"time"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("failed to read stdin: %s", err)
	}
	input = strings.TrimSuffix(input, "\n")
	bro, err := time.ParseDuration(input)
	if err != nil {
		log.Fatal("failed to parse: %s", err)
	}
	fmt.Println(bro.Minutes())
}
