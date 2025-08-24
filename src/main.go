package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   time.Time
	updateAt    time.Time
}

type DataBase struct {
	tasks  []Task
	nextID int
}

func newDataBase() *DataBase {
	return &DataBase{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func (db *DataBase) addTask(t Task) Task {
	t.id = db.nextID
	t.createdAt = time.Now()
	db.tasks = append(db.tasks, t)
	fmt.Printf("Task added successfully (ID: %v) \n", db.nextID)
	db.nextID++
	return t
}

func (db *DataBase) updateTask(t Task) bool {
	for i, task := range db.tasks {
		if task.id == t.id {
			db.tasks[i] = t
			fmt.Printf("New Task description: %v \n", db.tasks[i].description)
			return true
		}
	}

	return false
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("")
	fmt.Println(strings.Repeat("-", 28) + " TASK TRACKER IN GOLANG " + strings.Repeat("-", 28))
	fmt.Println("")
	fmt.Println(strings.Repeat("-", 80))

	fmt.Println("* Digit 'help' for list of commands")

	db := newDataBase()

	for {
		fmt.Print(">> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		cleanInput := strings.TrimSpace(input)
		parts := strings.SplitN(cleanInput, " ", 3)
		comand := parts[0]

		switch comand {

		case "add":
			db.addTask(Task{description: parts[1]})

		case "update":

			id, err1 := strconv.Atoi(parts[1])
			if err1 != nil {
				fmt.Printf("Error converting to int %v \n", err1)
				continue
			}
			db.updateTask(Task{id: id, description: parts[2]})

		case "delete":
		case "mark-in-progress":
		case "mark-done":
		case "list":
		case "list done":
		case "list todo":
		case "list in-progress":

		case "exit":
			fmt.Println("Quitting an application...")
			return
		default:
			fmt.Println("Unknown command, type help to see the list of supported commands.")
		}

	}
}
