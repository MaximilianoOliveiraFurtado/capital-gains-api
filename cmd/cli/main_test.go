package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func runMainUseCasesTest(t *testing.T, inputFilePath, expectedOutput string) {
	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		t.Fatalf("error loading file: %v", err)
	}

	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("error creating pipe: %v", err)
	}

	go func() {
		defer w.Close()
		w.Write(input)
	}()

	os.Stdin = r

	stdout := new(bytes.Buffer)
	os.Stdout = os.NewFile(0, "/dev/null")
	rStdout, wStdout, err := os.Pipe()
	if err != nil {
		t.Fatalf("error creating stdout pipe: %v", err)
	}
	os.Stdout = wStdout

	main()

	wStdout.Close()
	io.Copy(stdout, rStdout)

	output := strings.TrimSpace(stdout.String())
	expectedOutput = strings.TrimSpace(expectedOutput)

	if output != expectedOutput {
		t.Errorf("expected %v, found %v", expectedOutput, output)
	}
}

func TestMainApplicationUseCase_1(t *testing.T) {
	runMainUseCasesTest(t, "../../test/integration/data/case_1.txt", `[{"tax":0},{"tax":0},{"tax":0}]`)
}

func TestMainApplicationUseCase_2(t *testing.T) {
	runMainUseCasesTest(t, "../../test/integration/data/case_2.txt", `[{"tax":0},{"tax":10000},{"tax":0}]`)
}

func TestMainApplicationUseCase_3(t *testing.T) {
	runMainUseCasesTest(t, "../../test/integration/data/case_3.txt", `[{"tax":0},{"tax":0},{"tax":1000}]`)
}

func TestMainApplicationUseCase_4(t *testing.T) {
	runMainUseCasesTest(t, "../../test/integration/data/case_4.txt", `[{"tax":0},{"tax":0},{"tax":0}]`)
}

func TestMainApplicationUseCase_5(t *testing.T) {
	runMainUseCasesTest(t, "../../test/integration/data/case_5.txt", `[{"tax":0},{"tax":0},{"tax":0},{"tax":10000}]`)
}

func TestMainApplicationUseCase_6(t *testing.T) {
	runMainUseCasesTest(t, "../../test/integration/data/case_6.txt", `[{"tax":0},{"tax":0},{"tax":0},{"tax":0},{"tax":3000}]`)
}

func TestMainApplicationUseCase_7(t *testing.T) {
	runMainUseCasesTest(t, "../../test/integration/data/case_7.txt",
		`[{"tax":0},{"tax":0},{"tax":0},{"tax":0},{"tax":3000},{"tax":0},{"tax":0},{"tax":3700},{"tax":0}]`)
}
