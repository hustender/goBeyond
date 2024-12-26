//go:build !windows

package main

import (
	"fmt"
	"github.com/hustender/goBeyond/cmd"
	"os"
	"runtime"
)

func main() {
	fmt.Println(cmd.Error(fmt.Sprintf("This program only runs on Windows. Your current OS is: %s\n", runtime.GOOS)))
	os.Exit(1)
}
