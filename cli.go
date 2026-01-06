package main

import (
	"fmt"
)

func StartApp() {
	fmt.Print("Welcome to CLI app - Task Tracker\n")
	str := Todo{ID: 2, Title: "12fesfesfsef313"}
	SaveTodo("file.json", []Todo{str})
	//scanner := bufio.NewScanner(os.Stdin)
	//todonew := Todo{ID: 1, Title: "сделать дз"}
	// for {
	// 	fmt.Print("Enter command >> ")
	// 	scanner.Scan()
	// 	command := scanner.Text()
	// 	args := strings.Fields(command) /*принимаем команды и делаем из строки слайсы*/

	// 	switch args[0] {
	// 	case "add":
	// 		todo_title := strings.Join(args[1:], " ")
	// 	case "update":
	// 	case "delete":

	// 	}
	// }

}
