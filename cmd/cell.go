/*
Copyright © 2023 Isaac Riley <isaac.r.riley@gmail.com>
*/
package cmd

import (
	"strings"

	sf "github.com/sa-/slicefunk"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MakeCodeLine(codeStr string, lineLength int) string {

	var textLength int = lineLength - 4
	codeStr = strings.Trim(codeStr[0:min(len(codeStr), textLength)], "\n")
	var interior = codeStr + strings.Repeat(" ", textLength-len(codeStr))
	var line = vertical + " " + interior + " " + vertical

	return line
}

func FrameStrings(codeStrings []string, lineLength int) string {
	var line = strings.Repeat(horizontal, lineLength-2)
	var topLine = topLeft + line + topRight
	var bottomLine = bottomLeft + line + bottomRight

	var codeLines = sf.Map(
		codeStrings,
		func(str string) string { return MakeCodeLine(str, lineLength) },
	)
	var codeBlock = strings.Join(codeLines, "\n")
	if len(codeStrings) > 0 {
		codeBlock += "\n"
	}

	cell := topLine + "\n" + codeBlock + bottomLine

	return cell
}
