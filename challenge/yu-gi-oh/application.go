package main

import (
	"bufio"
	"fmt"
	"github.com/ttacon/chalk"
	"os"
	"strconv"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(chalk.Cyan, "1. New Game")
		fmt.Println(chalk.Cyan, "2. Exit")
		scanner.Scan()
		choiceInput := scanner.Text()
		choice, err := strconv.Atoi(choiceInput)
		if err != nil {
			fmt.Println("Invalid option, please enter a number.")
			continue
		}

		switch choice {
		case 1:
		case 2:
			fmt.Println(chalk.Yellow, "See you later...")
			os.Exit(0)
		}
	}
}
