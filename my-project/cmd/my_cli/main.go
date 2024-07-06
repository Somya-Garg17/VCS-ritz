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
        }

        subcommand := os.Args[2]
        switch subcommand {
        case "init":
            ritz.Init()
        case "add":
            if len(os.Args) < 4 {
                fmt.Println("Usage: ./my_cli ritz add <file>")
                os.Exit(1)
            }

            files := os.Args[3:]
            ritz.Add(files)
        case "commit":
            if len(os.Args) < 5 || os.Args[3] != "-m" {
                fmt.Println("Usage: ./my_cli ritz commit -m <message>")
                os.Exit(1)
            }
            message := os.Args[4]
            ritz.Commit(message)
        case "status":
            ritz.Status()
        case "diff":
            if len(os.Args) < 4 {
                fmt.Println("Usage: ./my_cli ritz diff <file>")
                os.Exit(1)
            }
            filePath := os.Args[3]
            ritz.Diff(filePath)
        case "branch":
            handleBranchCommand(os.Args[3:])
        case "checkout":
            handleCheckoutCommand(os.Args[3:])
        default:
            fmt.Println("Invalid ritz subcommand:", subcommand)
            os.Exit(1)
        }
    default:
        fmt.Println("Invalid command:", command)
        os.Exit(1)
    }
}

func handleBranchCommand(args []string) {
    if len(args) == 0 {
        err := ritz.ListBranches()
        if err != nil {
            fmt.Println("Error listing branches:", err)
        }
        return
    }

    switch args[0] {
    case "-d":
        if len(args) != 2 {
            fmt.Println("Usage: ./my_cli ritz branch -d <branch>")
            return
        }
        err := ritz.DeleteBranch(args[1])
        if err != nil {
            fmt.Println("Error deleting branch:", err)
        }
    case "-m":
        if len(args) != 3 {
            fmt.Println("Usage: ./my_cli ritz branch -m <old> <new>")
            return
        }
        err := ritz.RenameBranch(args[1], args[2])
        if err != nil {
            fmt.Println("Error renaming branch:", err)
        }
    default:
        err := ritz.CreateBranch(args[0])
        if err != nil {
            fmt.Println("Error creating branch:", err)
        }
    }
}

func handleCheckoutCommand(args []string) {
    if len(args) != 1 {
        fmt.Println("Usage: ./my_cli ritz checkout <branch>")
        return
    }

    err := ritz.CheckoutBranch(args[0])
    if err != nil {
        fmt.Println("Error checking out branch:", err)
    }
}

