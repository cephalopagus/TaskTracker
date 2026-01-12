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

		switch command {

		case "help":
			HelpCommand()

		case "exit":
			fmt.Print("\n=== APP SHUTING DOWN ===")
			return

		case "add":
			new_text := args[1:]
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

		case "mark":
			if args[1] == "in-progress" {
				idTask, err := strconv.Atoi(string(args[2]))
				if err != nil {
					fmt.Print("\n!!! WRONG ID !!!\n\n")
				} else {
					MarkInProgress(idTask)
				}
			} else if args[1] == "done" {
				idTask, err := strconv.Atoi(string(args[2]))
				if err != nil {
					fmt.Print("\n!!! WRONG ID !!!\n\n")
				} else {
					MarkDone(idTask)
				}
			}

		case "update":
			updateText := args[2:]
			idTask, err := strconv.Atoi(string(args[1]))
			if err != nil {
				fmt.Print("\n!!! WRONG ID !!!\n\n")
			} else {
				UpdateTask(idTask, updateText)
			}

		}

	}

}
