package neem

import (
	"fmt"
	"time"
	"os"
)

func Log(message string){
	if os.Getenv("MODE") == "prod" {
		return
	} else if os.Getenv("MODE") == "dev" {
		now := time.Now()
		timestamp := now.Format("2006-01-02 15:04:05")
		text := timestamp + " " + f(Green, "Log: ") + f(LightCyan, message)
		fmt.Println(text)	
		return
	}
}