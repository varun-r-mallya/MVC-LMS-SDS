package neem

import (
	"fmt"
	"time"
)

func Log(message string){
	now := time.Now()
	timestamp := now.Format("2006-01-02 15:04:05")
	text := timestamp + " " + f(Green, "Log: ") + f(LightCyan, message)
	fmt.Println(text)
}