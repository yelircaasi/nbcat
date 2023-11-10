/*
Copyright © 2023 Isaac Riley <isaac.r.riley@gmail.com>
*/
package cmd

import (
	"fmt"
	"strings"

	sf "github.com/sa-/slicefunk"
)

func StringifyRawNotebook(raw RawNotebook, lineLength int) string {

	var header = []string{"\n\n"}
	var segments = sf.Map(raw.Cells, func(cell NBCell) string { return StringifySegment(cell, lineLength) })

	return strings.Join(append(header, segments...), "\n")
}

func StringifySegment(cell NBCell, lineLength int) string {
	var outputString string

	if cell.CellType == "code" {
		outputString = MakeCodeCellString(cell, lineLength)
	} else {
		outputString = MakeMarkdownString(cell, lineLength)
	}

	return outputString
}

func MakeCodeCellString(cell NBCell, lineLength int) string {
	var output = strings.Join(
		sf.Map(
			cell.Outputs,
			func(s Output) string { return strings.Join(s.Text, "") },
		),
		"\n",
	)

	return fmt.Sprintf("[%d]\n", cell.ExecutionCount) + FrameStrings(cell.Source, lineLength) + "\n" + output
}

func MakeMarkdownString(cell NBCell, lineLength int) string {
	var line = strings.Repeat(horizontal, lineLength) + "\n\n"
	return line + strings.Join(cell.Source, "") + "\n"
}
