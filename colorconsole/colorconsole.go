package colorconsole

import (
	"fmt"
	"runtime"
)

type TextColor int

const (
	Black TextColor = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func SetTextColor(color TextColor) {
	if runtime.GOOS == "windows" {
		fmt.Printf("\x1b[38;5;%dm", color)
	} else {
		fmt.Printf("\x1b[0;3%dm", color)
	}
}

func SetBackgroundColor(color TextColor) {
	if runtime.GOOS == "windows" {
		fmt.Printf("\x1b[48;5;%dm", color)
	} else {
		fmt.Printf("\x1b[4%dm", color)
	}
}

func ResetColor() {
	if runtime.GOOS == "windows" {
		fmt.Print("\x1b[0m")
	} else {
		fmt.Print("\x1b[0;0m")
	}
}

func ColoredText(text string, textColor TextColor, backgroundColor TextColor) string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("\x1b[38;5;%dm\x1b[48;5;%dm%s\x1b[0m", textColor, backgroundColor, text)
	} else {
		return fmt.Sprintf("\x1b[0;3%dm\x1b[4%dm%s\x1b[0;0m", textColor, backgroundColor, text)
	}
}
