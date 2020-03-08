package main

import (
	"fmt"
	"github.com/rickshawdriver/somebody/cmd"
	"os"
)

func main() {
	rootCmd := cmd.GetRootCmd(os.Args[:1])
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
