package main

import (
	"fmt"
	"os/exec"
)

func main() {
	inputFile := "../downloaded_file.csv"
	outputFile := "sorted.csv"
	columnName := "ENTRY_DATE"

	cmd := exec.Command("python3", "../../sort_csv.py", inputFile, outputFile, columnName)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %s\n", err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
