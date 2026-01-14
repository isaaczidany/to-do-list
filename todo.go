package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Text string
	Done bool
}

func addTask(tasks []Task, text string) []Task {
	tasks = append(tasks, Task{
		Text: text,
		Done: false,
	})
	return tasks
}
func listTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa encontrada.")
	}
	fmt.Println("Lista de Tarefas:")
	for i, tasks := range tasks {
		status := "[Pendente]" //pendente
		if tasks.Done {
			status = "[Concluída]" //concluída
		}
		fmt.Printf("%s | %d - %s \n", status, i+1, tasks.Text)
	}
}
func removeTask(tasks []Task, index string) []Task {
	num, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("digite um número válido")
	}
	num = num - 1
	if num < 0 || num >= len(tasks) {
		fmt.Println("Índice Inválido")
		return tasks
	}
	return append(tasks[:num], tasks[num+1:]...)
}
func completeTask(tasks []Task, index string) []Task {
	num, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("digite um número válido")
	}
	num = num - 1
	if num < 0 || num >= len(tasks) {
		fmt.Println("Índice Inválido")
		return tasks
	}
	tasks[num].Done = true
	return tasks
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var tasks []Task
	var running bool = true
	for running == true {
		fmt.Println("\033[1;34m============== TODO LIST ==============\033[0m")
		fmt.Println("1 | Adicionar Tarefa")
		fmt.Println("2 | Listar Tarefas")
		fmt.Println("3 | Marcar Tarefa como Concluída")
		fmt.Println("4 | Remover Tarefa da Lista")
		fmt.Println("0 | Sair")
		fmt.Print("Escolha uma opção: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			fmt.Print("Digite a tarefa: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			tasks = addTask(tasks, text)
			fmt.Println("Tarefa adicionada!")
		case "2":
			listTask(tasks)
		case "3":
			fmt.Println("Qual tarefa será concluída?")
			index, _ := reader.ReadString('\n')
			index = strings.TrimSpace(index)
			tasks = completeTask(tasks, index)

		case "4":
			fmt.Println("Qual item será removido?")
			index, _ := reader.ReadString('\n')
			index = strings.TrimSpace(index)

			tasks = removeTask(tasks, index)
			fmt.Println("Tarefa Removida")
		case "0":
			running = false
		default:
			running = false
		}
	}
}
