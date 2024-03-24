package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"capital-gains/internal/entity"
	"capital-gains/internal/service/operation"
	"capital-gains/internal/service/tax"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var operationsTaxes []entity.Tax
	taxService := tax.NewService(&entity.Finstate{})
	operationService := operation.NewService(taxService)

	for {
		input, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading from stdin:", err)
			os.Exit(1)
		}

		input = strings.TrimSuffix(input, "\n")

		if input == "" {
			break
		}

		operations := operationService.InputParseOperation(input)

		for _, operationInputed := range operations {
			operationTax := operationService.OperationTax(&operationInputed)
			operationsTaxes = append(operationsTaxes, *operationTax)
		}

		jsonData, err := json.Marshal(operationsTaxes)
		if err != nil {
			fmt.Println("error converting tax list to json:", err)
			return
		}

		fmt.Println(string(jsonData))

	}

}
