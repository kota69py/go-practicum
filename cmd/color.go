package cmd

import (
	"runtime"
	"strings"
)

var (
	useColor = runtime.GOOS != "windows" || isModernTerminal()
)

func isModernTerminal() bool {
	return true
}

func colorGreen(s string) string {
	if !useColor {
		return s
	}
	return "\033[32m" + s + "\033[0m"
}

func colorRed(s string) string {
	if !useColor {
		return s
	}
	return "\033[31m" + s + "\033[0m"
}

func colorYellow(s string) string {
	if !useColor {
		return s
	}
	return "\033[33m" + s + "\033[0m"
}

func colorCyan(s string) string {
	if !useColor {
		return s
	}
	return "\033[36m" + s + "\033[0m"
}

func stars(n int) string {
	if n > 5 {
		n = 5
	}
	return strings.Repeat("★", n) + strings.Repeat("☆", 5-n)
}
