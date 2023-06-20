package main

import (
	"github.com/erfanmomeniii/user-management/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
