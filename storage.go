package main

import (
	"encoding/json"
	"os"
)

const FileName = "tasks.json"

func LoadTask() (TaskList, error) {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		return TaskList{}, nil
	}

	data, err := os.ReadFile(FileName)
	if err != nil {
		return nil, err
	}

	var tasks TaskList

	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(tasks TaskList) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(FileName, data, 0644)
}
