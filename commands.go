package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func AddCommand(args []string) {

	description := strings.Join(args, " ")

	oldTask, err := LoadTask()

	if err != nil {
		fmt.Println(err)
		return
	}
	id := 1

	if len(oldTask) > 0 {
		id = oldTask[len(oldTask)-1].ID + 1

	}
	newTask := Task{
		ID:          id,
		Description: description,
		Status:      "To do",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	oldTask = append(oldTask, newTask)

	if err := SaveTasks(oldTask); err != nil {
		return
	}
	fmt.Print("\n=== TASK ADDED ===\n\n")
}

func ListCommand() {
	list, err := LoadTask()
	if err != nil {
		fmt.Println(err)
		return
	}
	w := tabwriter.NewWriter(
		os.Stdout,
		0, 0, 5, ' ',
		tabwriter.Debug,
	)
	fmt.Println()
	fmt.Fprintln(w, "ID\tDESCRIPTION\tSTATUS\tCREATED\tUPDATED")
	fmt.Fprintln(w, "--\t------\t-----------\t-------\t-------")

	for _, t := range list {
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%s\n",
			t.ID,
			t.Description,
			t.Status,
			t.CreatedAt.Format(time.RFC822),
			t.UpdatedAt.Format(time.RFC822),
		)
	}

	w.Flush()
	fmt.Println()

}

func DeleteCommand(idTask int) {
	oldTask, err := LoadTask()

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, v := range oldTask {
		if v.ID == idTask {
			oldTask = append(oldTask[:i], oldTask[i+1:]...)
			SaveTasks(oldTask)
			fmt.Print("\n=== TASK DELETED ===\n\n")
			return
		}

	}

	fmt.Print("\n!!! ID NOT FOUND !!!\n\n")

}

func HelpCommand() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)

	fmt.Fprintln(w, "\nCommands:")

	fmt.Fprintln(w, " >> add `description`\tAdd new task")

	fmt.Fprintln(w, " >> list\tShow all tasks")

	fmt.Fprintln(w, " >> list todo\tShow todo tasks")

	fmt.Fprintln(w, " >> list in-progress\tShow in progress tasks")

	fmt.Fprintln(w, " >> list done\tShow done tasks")

	fmt.Fprintln(w, " >> delete `id`\tDelete task by id")

	fmt.Fprintln(w, " >> update `id` `description`\tUpdate description by id")

	fmt.Fprintln(w, " >> mark-in-progress `id`\tMarking a task as in progress by id")

	fmt.Fprintln(w, " >> mark-done `id`\tMarking a task as in done by id")

	w.Flush()
	fmt.Println()
}

func MarkInProgress(idTask int) {
	oldTask, err := LoadTask()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range oldTask {
		if v.ID == idTask {
			oldTask[idTask-1].Status, oldTask[idTask-1].UpdatedAt = "In progress", time.Now()

			SaveTasks(oldTask)
			fmt.Print("\n=== STATUS CHANGED ===\n\n")
			return
		}
	}
	fmt.Print("\n!!! ID NOT FOUND !!!\n\n")
}

func MarkDone(idTask int) {
	oldTask, err := LoadTask()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range oldTask {
		if v.ID == idTask {
			oldTask[idTask-1].Status, oldTask[idTask-1].UpdatedAt = "Done", time.Now()

			SaveTasks(oldTask)
			fmt.Print("\n=== STATUS CHANGED ===\n\n")
			return
		}
	}
	fmt.Print("\n!!! ID NOT FOUND !!!\n\n")
}

func ListByStatus(status string) {
	list, err := LoadTask()
	if err != nil {
		fmt.Println(err)
		return
	}
	w := tabwriter.NewWriter(
		os.Stdout,
		0, 0, 5, ' ',
		tabwriter.Debug,
	)
	fmt.Println()
	fmt.Fprintln(w, "ID\tDESCRIPTION\tSTATUS\tCREATED\tUPDATED")
	fmt.Fprintln(w, "--\t------\t-----------\t-------\t-------")

	switch status {
	case "done":
		for _, t := range list {
			if t.Status == "Done" {
				fmt.Fprintf(
					w,
					"%d\t%s\t%s\t%s\t%s\n",
					t.ID,
					t.Description,
					t.Status,
					t.CreatedAt.Format(time.RFC822),
					t.UpdatedAt.Format(time.RFC822),
				)
			}

		}
	case "in-progress":
		for _, t := range list {
			if t.Status == "In progress" {
				fmt.Fprintf(
					w,
					"%d\t%s\t%s\t%s\t%s\n",
					t.ID,
					t.Description,
					t.Status,
					t.CreatedAt.Format(time.RFC822),
					t.UpdatedAt.Format(time.RFC822),
				)
			}

		}
	case "todo":
		for _, t := range list {
			if t.Status == "To do" {
				fmt.Fprintf(
					w,
					"%d\t%s\t%s\t%s\t%s\n",
					t.ID,
					t.Description,
					t.Status,
					t.CreatedAt.Format(time.RFC822),
					t.UpdatedAt.Format(time.RFC822),
				)
			}

		}

	}

	w.Flush()
	fmt.Println()
}

func UpdateTask(idTask int, args []string) {

	description := strings.Join(args, " ")

	oldTask, err := LoadTask()

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range oldTask {
		if v.ID == idTask {
			oldTask[idTask-1].Description, oldTask[idTask-1].UpdatedAt = description, time.Now()
			SaveTasks(oldTask)
			fmt.Print("\n=== DESCRIPTION CHANGED ===\n\n")
			return
		}
	}

	fmt.Print("\n!!! ID NOT FOUND !!!\n\n")
}
