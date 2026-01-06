package main

import (
	"encoding/json"
	"os"
)

func CheckJson() {
	const file_name = "tasks.json"
	var todos []Todo
	data, err := os.ReadFile(file_name)
	if os.IsNotExist(err) {
		todos = []Todo{}
		SaveTodo(file_name, todos)

	} else if err != nil {
		panic(err)
	} else {
		if err := json.Unmarshal(data, &todos); err != nil {
			todos = []Todo{}
			SaveTodo(file_name, todos)
		}
	}

}
func SaveTodo(file_name string, todos []Todo) {
	file, err := os.Create(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(todos); err != nil {
		panic(err)
	}
}
