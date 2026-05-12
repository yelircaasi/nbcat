package cmd

import (
	"strings"

	sf "github.com/sa-/slicefunk"
)

func MakeCodeLine(codeStr string, lineLength int) string {
	textLength := lineLength - 4

	trimmedRaw := strings.Trim(codeStr[0:min(len(codeStr), textLength)], "\n")
	padding := strings.Repeat(" ", textLength-len(trimmedRaw))

	highlighted := HighlightSource(trimmedRaw)

	return vertical + " " + highlighted + padding + " " + vertical
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
