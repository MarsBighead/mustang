package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	evbus "github.com/asaskevich/EventBus"
)

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func main() {
	scan()
	bus := evbus.New()

	bus.Subscribe("main:calculator", calculator)
	bus.Publish("main:calculator", 20, 40)
	bus.Unsubscribe("main:calculator", calculator)
}

func scan() {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 64)
		}
		fmt.Printf("Split: %d, %s\n", advance, string(token))
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)

	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
