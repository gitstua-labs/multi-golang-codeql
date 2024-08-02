package main

import (
	"fmt"
	"time"
)

func main() {
	currentDate := time.Now().Format("2006-01-02")
	fmt.Println("Current date:", currentDate)
}