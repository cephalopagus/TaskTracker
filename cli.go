package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StartApp() {
	fmt.Print("\n=== Welcome to CLI app - Task Tracker ===\n\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command >> ")
		scanner.Scan()
		text := scanner.Text()
		args := strings.Fields(text) /*принимаем команды и делаем из строки слайсы*/
		command := args[0]
		new_text := args[1:]
		switch {
		case command == "help":
			HelpCommand()
		case command == "exit":
			fmt.Print("\n=== APP SHUTING DOWN ===")
			return
		case command == "add":
			AddCommand(new_text)
		case command == "list":
			ListCommand()
		case command == "delete":

			idTask, err := strconv.Atoi(string(args[1]))
			if err != nil {
				fmt.Print("\n!!! WRONG ID !!!\n\n")
			} else {
				DeleteCommand(idTask)
			}

		}
	}

}
