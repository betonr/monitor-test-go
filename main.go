package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 2
const delay = 5

func main() {

	exibeIntro()
	getSiteFile()

	for {
		exibeMenu()

		comando := getComando()
		switch comando {
		case 1:
			startMonitamento()

		case 2:
			fmt.Println("exibindo logs")
			imprimeLogs()

		case 0:
			fmt.Println("Você saiu, até breve!")
			os.Exit(0)

		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}

func exibeIntro() {
	nome := "Beto Noronha"
	versao := 1.2

	fmt.Println("=============================================")
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Seja Bem-vindo ao monitorador de sites!!!")
	fmt.Println("Esse programa está na versão", versao)
	fmt.Println("=============================================")
	fmt.Println()
}

func exibeMenu() {
	fmt.Println("1- iniciar monitoramento")
	fmt.Println("2- exibir logs")
	fmt.Println("0- sair do programa")
	fmt.Println()
}

func getComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func startMonitamento() {
	fmt.Println("iniciou monitoramento ...")
	fmt.Println()

	sites := getSiteFile()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("testando site", i)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}

	fmt.Println()
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site: ", site, "foi carregado com SUCESSO!")
		writeLog(site, true)
	} else {
		fmt.Println("O site: ", site, "está com problemas, status =", resp.StatusCode)
		writeLog(site, false)
	}
}

func getSiteFile() []string {

	file, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Erro:", err)
	}

	return strings.Split(string(file), "\n")
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

// func getSiteFile() {

// 	var sites []string

//     arquivo, err := os.Open("sites.txt")
//     if err != nil {
//         fmt.Println("Ocorreu um erro:", err)
//     }

//     leitor := bufio.NewReader(arquivo)

//     for {
//         linha, err := leitor.ReadString('\n')
//         linha = strings.TrimSpace(linha)
//         sites = append(sites, linha)
//         if err == io.EOF {
//             break
//         }
//     }

//     arquivo.Close()

// 	return sites
// }

func writeLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - O site:" + site + " está online:" + strconv.FormatBool(status) + "\n")

	file.Close()
}
