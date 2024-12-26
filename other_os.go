//go:build !windows

package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("This program only runs on Windows. Your current OS is: %s\n", runtime.GOOS)
	os.Exit(1)
}
