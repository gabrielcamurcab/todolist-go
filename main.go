package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Description string
	Completed   bool
}

func main() {
	tasks := make([]Task, 0)

	for {
		fmt.Println("")
		fmt.Println("1. Adicionar tarefa")
		fmt.Println("2. Marcar tarefa como concluída")
		fmt.Println("3. Exibir tarefas pendentes")
		fmt.Println("4. Exibir tarefas concluídas")
		fmt.Println("5. Sair")

		var choice int
		fmt.Print("Escolha uma opção:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(&tasks)
		case 2:
			markTaskCompleted(&tasks)
		case 3:
			displayPendingTasks(tasks)
		case 4:
			displayCompletedTasks(tasks)
		case 5:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func addTask(tasks *[]Task) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite a descrição da tarefa: ")
	if scanner.Scan() {
		description := scanner.Text()
		*tasks = append(*tasks, Task{Description: description, Completed: false})
		fmt.Println("Tarefa adicionada com sucesso!")
	}
}

func markTaskCompleted(tasks *[]Task) {
	fmt.Print("DIgite o número da tarefa que deseja marcar como concluída: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum >= 1 && taskNum <= len(*tasks) {
		(*tasks)[taskNum-1].Completed = true
		fmt.Println("Tarefa marcada como concluída!")
	} else {
		fmt.Println("Número de tarefa inválido!")
	}
}

func displayPendingTasks(tasks []Task) {
	fmt.Println("\nTarefas pendentes:")
	for i, task := range tasks {
		if !task.Completed {
			fmt.Printf("%d. %s\n", i+1, task.Description)
		}
	}
}

func displayCompletedTasks(tasks []Task) {
	fmt.Println("\nTarefas concluídas:")
	for i, task := range tasks {
		if task.Completed {
			fmt.Printf("%d. %s\n", i+1, task.Description)
		}
	}
}
