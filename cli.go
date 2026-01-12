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
		switch command {
		case "help":
			HelpCommand()
		case "exit":
			fmt.Print("\n=== APP SHUTING DOWN ===")
			return
		case "add":
			AddCommand(new_text)

		case "list":
			if len(args) == 2 {
				ListByStatus(args[1])
			} else {
				ListCommand()
			}

		case "delete":

			idTask, err := strconv.Atoi(string(args[1]))
			if err != nil {
				fmt.Print("\n!!! WRONG ID !!!\n\n")
			} else {
				DeleteCommand(idTask)
			}
		case "mark-in-progress":
			idTask, err := strconv.Atoi(string(args[1]))
			if err != nil {
				fmt.Print("\n!!! WRONG ID !!!\n\n")
			} else {
				MarkInProgress(idTask)
			}
		case "mark-done":
			idTask, err := strconv.Atoi(string(args[1]))
			if err != nil {
				fmt.Print("\n!!! WRONG ID !!!\n\n")
			} else {
				MarkDone(idTask)
			}
		}
	}

}
