package main

import (
	"fmt"
	"time"
)

func LogInfo(message string, args ...interface{}) {
	dt := time.Now()

	fmt.Printf("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	fmt.Printf("\033[0m"+message+"\r\n", args...)
}

func LogWarning(message string, args ...interface{}) {
	dt := time.Now()

	fmt.Printf("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	fmt.Printf("\033[93mWARNING: "+message+"\r\n", args...)
}

func LogError(message string, args ...interface{}) {
	dt := time.Now()

	fmt.Printf("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	fmt.Printf("\033[1;31mERROR: "+message+"\r\n", args...)
}
