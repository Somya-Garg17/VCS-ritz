package main

import (
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) < 2 {
        println("Usage: ritz <command>")
        os.Exit(1)
    }

    command := os.Args[2]


        switch command {
        case "init":
            executeCommand("/home/whyknow/go/src/my-project/pkg/ritz/init.go")
        case "commit":
            executeCommand("/home/whyknow/go/src/my-project/pkg/ritz/commit.go")
        default:
            println("Invalid ritz command:", command)
            os.Exit(1)
        }
    
}

func executeCommand(filePath string) {
    cmd := exec.Command("go", "run", filePath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        println("Error executing command:", err)
        os.Exit(1)
    }
}

