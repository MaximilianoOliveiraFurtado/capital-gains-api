package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var lines []string

	for {
		input, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
			os.Exit(1)
		}

		input = strings.TrimSuffix(input, "\n")

		if input == "" {
			break
		}

		lines = append(lines, input)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
