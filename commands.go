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
		Status:      "In process",
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
	defer w.Flush()
	fmt.Println()
	fmt.Fprintln(w, "ID\tSTATUS\tDESCRIPTION\tCREATED\tUPDATED")
	fmt.Fprintln(w, "--\t------\t-----------\t-------\t-------")

	for _, t := range list {
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%s\n",
			t.ID,
			t.Status,
			t.Description,
			t.CreatedAt.Format(time.RFC822),
			t.UpdatedAt.Format(time.RFC822),
		)
	}

}
