package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now().Format("15:04:05")
	fmt.Println("Current time:", currentTime)
}