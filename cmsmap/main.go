package main

import (
	"bufio"
	"cmsmap/cmd/superficial"
	"cmsmap/internal/commands"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Print ASCII art
	random_ascii_art := superficial.Select_ascii_art()
	fmt.Println(random_ascii_art)
	fmt.Println("cmsmap: A modern exploit CMS exploitation framework\n\n\n\n\n\n\n\n------------------------------------------------------\n")

	// Main loop
	for {
		fmt.Print("cmsmap )>>>")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			break
		}

		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		if handler, ok := commands.GetHandler(args[0]); ok {
			output := handler.Handle(args)
			fmt.Println(output)
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of available commands.")
		}

	}
}
