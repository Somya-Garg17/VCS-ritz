package main

import (
	"fmt"
	"os"

	"my-project/pkg/ritz"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./my_cli <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "ritz":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./my_cli ritz <subcommand>")
			os.Exit(1)
			return
		}

		subcommand := os.Args[2]
		switch subcommand {
		case "init":
			ritz.Init()
		case "add":
			if len(os.Args) < 4 {
				fmt.Println("Usage: ./my_cli ritz add <file>")
				os.Exit(1)
				return
			}

			files := os.Args[3:]
			ritz.Add(files)
		case "commit":
			if len(os.Args) < 4 {
				fmt.Println("Usage: ./my_cli ritz commit -m <message>")
				os.Exit(1)
				return
			}

			if os.Args[3] != "-m" {
				fmt.Println("Usage: ./my_cli ritz commit -m <message>")
				os.Exit(1)
				return
			}

			message := os.Args[4]
			ritz.Commit(message)
		default:
			fmt.Println("Invalid ritz subcommand:", subcommand)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid command:", command)
		os.Exit(1)
	}
}

