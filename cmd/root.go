//go:build windows

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goBeyond",
	Short: "goBeyond is a powerful and lightweight tool to uncover your system's specifications, written in Go.",
	Long:  "goBeyond is a powerful and lightweight tool to uncover your system's specifications, written in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available components: \n" +
			"[-] cpu \n" +
			"[-] memory \n" +
			"[-] gpu \n" +
			"[-] all \n" +
			"Usage: goBeyond [component]")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occured executing goBeyond \n'%s'\n", err)
		os.Exit(1)
	}
}
