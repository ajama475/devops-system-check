package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("--- DEVOPS SYSTEM CHECK ---")
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("System Status: READY FOR AUTOMATION")
}
