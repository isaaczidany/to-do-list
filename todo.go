package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addTask(tasks []string, text string) []string {
	tasks = append(tasks, text)
	return tasks
}
func listTask(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa encontrada.")
	}
	fmt.Println("Lista de Tarefas:")
	for i, tasks := range tasks {
		fmt.Printf("%d - %s\n", i+1, tasks)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tasks []string
	var running bool = true
	for running == true {
		fmt.Println("--------------- \n" + "1 | Adicionar Tarefa \n" + "2 | Listar Tarefas \n" + "0 | Sair \n" + "Escolha: ")
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
		case "0":
			running = false
		default:
			running = false
		}
	}
}
