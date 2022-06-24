package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
)


const monitoramento = 3
const delay = 5

func main(){
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		//Para apresentar qual o valor do endereço da variável
		//fmt.Println("O endereço da variável comando é:", &comando)
	
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do progr...")
			os.Exit(0)
		default:
			fmt.Println("Opção informada não existe...")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao(){
	versao := 1.2
	fmt.Println("Este programa está na versão", versao)	
}

func exibeMenu(){
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i:= 0; i < monitoramento; i++ {
		for i, site := range sites{
			fmt.Println("Testando site", i, ":", site)
			testeSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testeSite(site string) {
	resp, err := http.Get(site)

	if err != nil{
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	}else{
		fmt.Println("Site:", site, "está com probelmas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, erro := os.Open("sites.txt")
	
	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)	
			
		if erro == io.EOF {
            break
        }
	}

	arquivo.Close()

	return sites
}

func imprimeLogs() {

    arquivo, err := ioutil.ReadFile("log.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(string(arquivo))
}

func registraLog(site string, status bool){
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }
    arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

    arquivo.Close()
}