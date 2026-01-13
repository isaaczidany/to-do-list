package main

import ("fmt")

func addTask(tasks []string, text string) []string {
	tasks = append(tasks, text)
	return tasks
}
func main () {
	var tasks []string
	var option int
	for {
		fmt.Println("1 | Adicionar Tarefa \n" + "0 | Sair \n"+ "Escolha: ")
		fmt.Scan(&option)
		if option == 1 {
			var text string
			fmt.Print("Digite a tarefa: ")
			fmt.Scan(&text)
			tasks = addTask(tasks, text)
			fmt.Println("Tarefa adicionada!")
		} else if option == 0 {
		 break
		}
	}
}
