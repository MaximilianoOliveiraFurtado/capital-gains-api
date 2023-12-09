package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"capital-gains/internal/entity"
	"capital-gains/internal/service/operation"
	"capital-gains/internal/service/tax"
	"capital-gains/internal/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var lines []string
	taxSrvice = tax.NewService()
	operationService = operation.NewService(taxSrvice)

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

		operation, err = operation.NewService()

		utils.ParseEntity[entity.Operation](input)
		if err != nil {
			fmt.Println("Erro ao ler a linha x")
		}

		lines = append(lines, input)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
