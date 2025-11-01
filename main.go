package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second)
	fmt.Println("Hello, World!")
	fmt.Println(time.Since(start))
}
