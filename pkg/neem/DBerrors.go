package neem

import (
	"fmt"
)

func DBError(message string, err error) {
	text := f(Magenta, "Database Error: ") + f(White, message) + "\n" + f(LightMagenta, err.Error())
	fmt.Println(text)
}