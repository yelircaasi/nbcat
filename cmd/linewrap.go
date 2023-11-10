/*
Copyright © 2023 Isaac Riley <isaac.r.riley@gmail.com>
*/
package cmd

import (
	"strings"
)

func WrapLine(str string, strLength int) string {
	return strings.Repeat(str, 40)
}
