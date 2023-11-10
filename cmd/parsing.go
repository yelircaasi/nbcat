/*
Copyright © 2023 Isaac Riley <isaac.r.riley@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ParseNotebookFile(notebookPath string) RawNotebook {
	print("Parsing " + notebookPath + "\n")
	jsonFile, err := os.Open(notebookPath)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var rawNotebook RawNotebook
	json.Unmarshal(byteValue, &rawNotebook)

	fmt.Println("Successfully parsed " + notebookPath + "\n")
	// fmt.Println(rawNotebook.Cells[0].CellType)

	defer jsonFile.Close()

	return rawNotebook
}

type RawNotebook struct {
	Cells    []NBCell       `json:"cells"`
	Metadata MetadataStruct `json:"metadata"`
}

type NBCell struct {
	Source         []string `json:"source"`
	CellType       string   `json:"cell_type"`
	OutputType     string   `json:"output_type,omitempty"`
	Outputs        []Output `json:"outputs,omitempty"`
	ExecutionCount int      `json:"execution_count"`
}

type Output struct {
	Name string   `json:"name"`
	Type string   `json:"output_type"`
	Text []string `json:"text"`
}

type MetadataStruct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
