package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"capital-gains/internal/entity"
	"capital-gains/internal/service/operation"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var operationsTax []entity.Tax
	operationService := operation.NewService()

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

		operations, err := entity.ParseOperations(input)
		if err != nil {
			fmt.Fprintln(os.Stdout, []any{"Erro ao ler a linha %s: {%s}", operations, err}...)
		}

		for _, operationInputed := range operations {
			operationResult := operationService.OperationInput(operationInputed)
			operationsTax = append(operationsTax, operationResult)
		}

		//operationResult := operationService.OperationInput(input)
		//operationsTax = append(operationsTax, operationResult)

		// for _, operationInputed := range input {
		// 	operationResult := operationService.OperationInput(operationInputed)
		// 	operationsTax = append(operationsTax, operationResult)
		// }

	}

	fmt.Println(operationsTax)

	// for _, operation := range operationsTax {
	// 	fmt.Println(operation)
	// }
}
