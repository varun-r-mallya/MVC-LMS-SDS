package neem

import (
	"fmt"
	"time"
)

type Color int

const (
	Black Color = iota
	Red
	Green
	YellowError
	Blue
	Magenta
	Cyan
	LightGray
	DarkGray
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

var colorCodes = [...]string{
	"\033[0;30m",
	"\033[0;31m",
	"\033[0;32m",
	"\033[0;33m",
	"\033[0;34m",
	"\033[0;35m",
	"\033[0;36m",
	"\033[0;37m",
	"\033[1;30m",
	"\033[1;31m",
	"\033[1;32m",
	"\033[1;33m",
	"\033[1;34m",
	"\033[1;35m",
	"\033[1;36m",
	"\033[1;37m",
}

func f(color Color, message string) string {
	return (colorCodes[color] + message + "\033[0m")
}

// A spotlight error is a non fatal error
func Spotlight(err error, message string) {
	if err != nil {
		text := f(Cyan, "Spotlit Error: ") + f(White, message) + "\n" + f(LightBlue, err.Error())
		fmt.Println(text)
	}
}

// Critical is a program stopping error. This is logged and the program is stopped
func Critial(err error, message string) {
	now := time.Now()
	timestamp := now.Format("2006-01-02 15:04:05")
	text := timestamp + " " + f(Red, "Critial Error: ") + f(LightRed, message) + "\n" + f(DarkGray, err.Error())
	panic(text)
}
