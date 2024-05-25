package main

import (
	"fmt"
	"os"

	"my-project/pkg/ritz"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: my-project <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "ritz":
		if len(os.Args) < 3 {
			fmt.Println("Usage: my-project ritz <subcommand>")
			os.Exit(1)
		}

		subcommand := os.Args[2]
		switch subcommand {
		case "init":
			ritz.Init()
		case "commit":
			ritz.Commit()
		default:
			fmt.Println("Invalid ritz subcommand:", subcommand)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid command:", command)
		os.Exit(1)
	}
}

