package main

import (
	"bufio"
	"fmt"
	"os"
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
	db.nextID++
	return t
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
			fmt.Println("Erro ao ler a entrada:", err)
			continue
		}
		cleanInput := strings.TrimSpace(input)
		parts := strings.SplitN(cleanInput, " ", 2)
		comand := parts[0]

		switch comand {

		case "add":
			db.addTask(Task{description: parts[1]})
		case "update":
		case "delete":
		case "mark-in-progress":
		case "mark-done":
		case "list":
		case "list done":
		case "list todo":
		case "list in-progress":

		case "exit":
			fmt.Println("Saindo da aplicação...")
			return
		}

	}
}
