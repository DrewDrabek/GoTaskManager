package tasks

import (
	"errors"
	"fmt"
)

// This it the struct for the task - creates a new type for it - hold name and date

type Task struct {
	Name       string
	ExpireDate int
}

// Craeting a function to add tasks to a list -  it takes in a pointer to a list and then the attributes of what gets added

func AddTask(l *[]Task, name string, expireDate int) error {

	if name == "" {
		return errors.New("missing the name input")
	}

	if expireDate <= 0 {
		return errors.New("missing the date input")
	}

	*l = append(*l, Task{Name: name, ExpireDate: expireDate})

	return nil
}

// This is a function to view tasks

func ViewTasks(l *[]Task) error {
	if len(*l) == 0 {
		return errors.New("there are no tasks in the task list given")
	}

	// this loops through everything to print out the tasks - really i should use the range function here but that is fine for now

	for i := 0; i < len(*l); i++ {
		n := (*l)[i]
		fmt.Printf("%d: %s (ExpireDate: %d)\n", i, n.Name, n.ExpireDate)
	}

	return nil

}

// This is a function to remove items from a task

func RemoveTask(l *[]Task, t string) error {
	if t == "" {
		return errors.New("missing task input")
	}

	if len(*l) == 0 {
		return errors.New("there are no tasks in the list")
	}

	// this loops through the slice and then compares them to the item that is passed and then removes them

	for i := 0; i <= len(*l); i++ { // should use the range here instead
		r := (*l)[i]
		if t == r.Name {

			// go does not have a built in remove function there might be some out there but this is all default
			// This removes by creating a new slice from the old one effectively replacing it
			// it add the items up till the index we are ona nd then everything after it

			*l = append((*l)[:i], (*l)[i+1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}
