package main

import (
	"strings"
	"io"
	"fmt"
	"net/http"
	"os"
	"time"
	"bufio"
	"strconv"
	"io/ioutil"
)

const monitoring = 3
const delay = 5

//import "reflect"

func main() {

	//fmt.Scanf("%d", &command)

	//fmt.Printf("O valor é ", command)

	/*
		if command == 1 {
			fmt.Print("\nMonitorando...\n")
		} else if command == 2 {
			fmt.Print("\nExibindo logs...\n")
		} else if command == 0 {
			fmt.Print("\nSaindo do programa...\n")
		} else {
			fmt.Print("\nComando inválido!\n")
		}
	*/

	name, age := returnNameAndAge()
	fmt.Print("Hello, ", name, " your age is ", age, "\n")

	for {
		menu()
		command := readCommand()

		switch command {
		case 1:
			toMonitor()
		case 2:
			printLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido!")
			os.Exit(-1)
		}
	}

}

func returnNameAndAge() (string, int) {
	//var name = "Gleydson"
	name := "Gleydson"
	age := 22

	return name, age
	//fmt.Printf("The type of the age variable is ", reflect.TypeOf(version))
}

func menu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func toMonitor() {
	fmt.Println("Monitorando...")

	//sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}
	//sites = append(sites, "https://demo.greenmile.com/#/")
	/*
		for i := 0; i < len(sites); i++ {
			fmt.Println(sites[i])
		}*/

	sites := readFileSites()

	for i := 0; i < monitoring; i++ {
		fmt.Print("Executando o ", i+1, "º teste...\n")
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		logRegister(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		logRegister(site, false)
	}
}

func readFileSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.Close()

	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") +  " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLog() {
	fmt.Println("Exibindo logs...")
	
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}
