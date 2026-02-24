package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("pulseforge pure-go baseline")
	fmt.Println("usage:")
	fmt.Println("  go run . --scripted [--variant core|go_native]")
}

func main() {
	variant := "go_native"
	scripted := false

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--scripted":
			scripted = true
		case "--variant":
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "missing value for --variant")
				os.Exit(2)
			}
			variant = args[i+1]
			i++
		case "scripted":
			scripted = true
		case "help", "--help", "-h":
			printUsage()
			return
		default:
			fmt.Fprintf(os.Stderr, "unknown argument: %s\n", arg)
			os.Exit(2)
		}
	}

	if !scripted {
		printUsage()
		return
	}

	runtime, err := newRuntime("pure_go", variant)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
	fmt.Println(runScripted(runtime))
}
