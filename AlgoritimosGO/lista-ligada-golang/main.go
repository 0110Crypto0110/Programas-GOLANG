package main
// programa de um lista ligada em python/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ____________________________________________________________
//define a estrura da NO
type No struct {
	valor   int
	proximo *No
}

//____________________________________________________________
//define a estrura da lista ligada
type ListaLigada struct {
	inicio *No
}

//____________________________________________________________
//metodo inserir_meio
//____________________________________________________________
//-Este método insere um valor em uma posição específica da lista.
//-É útil quando você precisa inserir um elemento em uma posição específica da lista.
//-O método recebe três parâmetros: a lista, o valor a ser inserido e
//a posição específica onde o valor deve ser inserido.
//-A posição é um número inteiro que indica a posição na lista onde o valor
//____________________________________________________________

func (list *ListaLigada) inserirmeio(indice, valor int) {
	if list.inicio == nil {
		fmt.Println("A lista está vazia. Não é possível inserir.")
		return
	}

	if indice == 0 {
		novoNo := &No{valor: valor, proximo: list.inicio}
		list.inicio = novoNo
		return
	}

	posicao := list.inicio
	for i := 0; i < indice-1 && posicao.proximo != nil; i++ {
		posicao = posicao.proximo
	}

	if posicao == nil || posicao.proximo == nil {
		fmt.Println("Erro: índice fora dos limites da lista.")
		return
	}

	novoNo := &No{valor: valor, proximo: posicao.proximo}
	posicao.proximo = novoNo
}

//____________________________________________________________
//metodo inserir_fim 
//____________________________________________________________
//-Este método insere um valor no final da lista.
//-É útil quando você precisa adicionar um elemento ao final da lista sem se preocupar com a ordem dos elementos.
//-O método recebe dois parâmetros: a lista e o valor a ser inserido.
//-O método percorre a lista até o último nó e adiciona um novo nó com o valor
//____________________________________________________________

func (list *ListaLigada) inserirFim(valor int) {
	novoNo := &No{valor: valor, proximo: nil}

	if list.inicio == nil {
		list.inicio = novoNo
		return
	}

	posicao := list.inicio
	for posicao.proximo != nil {
		posicao = posicao.proximo
	}

	posicao.proximo = novoNo
}

//____________________________________________________________
//metodo imprime
//____________________________________________________________
//-Este método imprime todos os elementos da lista.
//-É útil quando você precisa verificar se a lista está corretamente preenchida.
//-O método percorre a lista e imprime o valor de cada nó.
//____________________________________________________________

func (list *ListaLigada) imprime() {
	posicao := list.inicio
	for posicao != nil {
		fmt.Printf("%d ->", posicao.valor)
		posicao = posicao.proximo
	}
	fmt.Println()
}

//____________________________________________________________
//metodo esta_vazia
//____________________________________________________________
//-Este método verifica se a lista está vazia.
//-É útil quando você precisa verificar se a lista está vazia antes de realizar algo
//____________________________________________________________

func (list *ListaLigada) esta_vazio() {
	if list.inicio == nil {
		fmt.Print("esta vazio")
		fmt.Println()
		return
	}

}

//____________________________________________________________
//metodo remove_valor
//____________________________________________________________
//-Este método remove um valor da lista(o primeior a ser ecnotrado)
//-É útil quando você precisa remover um elemento específico da lista.
//-O método recebe dois parâmetros: a lista e o valor a ser removido.
//-O método percorre a lista até encontrar o valor a ser removido e remove o nó correspondente
//____________________________________________________________

func (list *ListaLigada) remove_valor(PROCURADO int) {
	posicao := list.inicio
	if posicao.valor == PROCURADO {
		list.inicio = posicao.proximo
	} else {
		for posicao.proximo != nil {
			if posicao.proximo.valor == PROCURADO {
				posicao.proximo = posicao.proximo.proximo
			}
			posicao = posicao.proximo
		}
	}
}

//____________________________________________________________
//metodo tamanho
//____________________________________________________________
//-Este método retorna o tamanho da lista.
//-É útil quando você precisa saber o número de elementos na lista.
//-O método percorre a lista contando o número de nós até o final e retorna o número de nos.
//____________________________________________________________

func (list *ListaLigada) Tamanho() (cont int) {
	posicao := list.inicio
	for posicao != nil {
		posicao = posicao.proximo
		cont++
	}
	return
}

//____________________________________________________________
//metodo limpa
//____________________________________________________________
//-Este método remove todos os elementos da lista.
//-É útil quando você precisa limpar a lista para reiniciá-la.
//-O método percorre a lista e remove cada nó, liberando a memória alocada para cada no.
//____________________________________________________________
func (list *ListaLigada) limpa() {
	for list.inicio != nil {
		proximo := list.inicio.proximo
		list.inicio = nil
		list.inicio = proximo
	}
}

//____________________________________________________________
//metodo procura
//____________________________________________________________
//-Este método procura um valor na lista e retorna se ele esta presente ou nao, true or false
//-É útil quando você precisa verificar se um valor está presente na lista.
//-O método recebe dois parâmetros: a lista e o valor a ser procurado.
//-O método percorre a lista até encontrar o valor procurado e retorna true se encontrar.
//____________________________________________________________

func (list *ListaLigada) procura(PROCURADO int) bool {
	posicao := list.inicio
	for posicao != nil {
		if posicao.valor == PROCURADO {
			return true
		}
		posicao = posicao.proximo
	}
	return false
}

//____________________________________________________________
//metodo transforma em lista
//____________________________________________________________
//-Este método converte uma string em uma lista de inteiros.
//-É útil quando você precisa converter uma string de entrada em uma lista de inteiros.
//-O método recebe uma string como entrada e retorna uma lista de inteiros.
//-A string de entrada é dividida em substrings usando o caractere espaço como delimitador.
//-Cada substring é convertida em um inteiro e adicionada à lista de inteiros.
//____________________________________________________________

func (list *ListaLigada) transformaEmLista() []int {
	var lista []int
	pos := list.inicio
	for pos != nil {
		lista = append(lista, pos.valor)
		pos = pos.proximo
	}
	return lista
}

//____________________________________________________________
//
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
//exemplo de execucao(leitura de arquivo externo)
//
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
//____________________________________________________________


func main() {
	// Registra o tempo de início
	start := time.Now()

	// Aqui vai o código da sua função main()

	// Exemplo: um loop simples, usado para calcular o tempo de execucao
	for i := 0; i < 1; i++ {
	// Lista ligada para armazenar os valores do arquivo
	lista := &ListaLigada{}

	fmt.Println("________________________________________________")
	fmt.Println("lendo um arquivo de texto externo")
	fmt.Println("________________________________________________")

	arquivo, err := os.Open("arq.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer arquivo.Close()

	linhas := []string{}
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	fmt.Println("________________________________________________")
	fmt.Println("quantidade de linhas do arquivo:", len(linhas))
	fmt.Println("________________________________________________")

	// Lendo o número de métodos a serem chamados
	segundaLinha := linhas[1]
	parte := strings.Split(segundaLinha, " ")
	numero, _ := strconv.Atoi(parte[0])

	// Inserindo os valores da primeira linha do arquivo
	primeiraLinha := linhas[0]
	valores := strings.Split(primeiraLinha, " ")
	for _, strValor := range valores {
			valor, err := strconv.Atoi(strValor)
			if err != nil {
					fmt.Println("Erro ao converter valor para inteiro:", err)
					continue
			}
			lista.inserirFim(valor)
	}

	// Imprimindo a lista ligada
	fmt.Println("Lista:")
	lista.imprime()

	fmt.Println("________________________________________________")
	fmt.Println("numero de metodos a serem chamados =", numero)
	fmt.Println("________________________________________________")

	// Imprimindo as linhas do arquivo indexadas
	for indice, linha := range linhas {
		fmt.Printf("Índice: %d, Linha: %s\n", indice, linha)
	}

	fmt.Println("________________________________________________")
	
	// Definindo o intervalo dos métodos
inicioIndice := 2
ultimoIndice := inicioIndice + numero

for j := inicioIndice; j < ultimoIndice; j++ {//iterando em um intervalo especifico
		parte := strings.Fields(linhas[j])//dividindo a linha em partes


		for _, valorStr := range parte {
			indice := strings.Index(valorStr, "A") // essa funcao indexa uma string 

				if strings.Contains(valorStr, "A") {//checando se a string contem a letra"A"
												
						valor1, err := strconv.Atoi(parte[indice+1])
						fmt.Println("erro:",err)

						posi1, err1 := strconv.Atoi(parte[indice+2])
						fmt.Println("erro:",err1)

						lista.inserirmeio(posi1, valor1)
						
						}
				if strings.Contains(valorStr, "P") {// se contem a letra "P"
						fmt.Println("-----------------------------------------")
						lista.imprime()
						fmt.Println("-----------------------------------------")


				}

				if strings.Contains(valorStr, "R") {//se contem a letra "R"

						valor2, err3 := strconv.Atoi(parte[indice+2])
						fmt.Println("erro:",err3)

						lista.remove_valor(valor2)

				}
				
			}
	
		}
		}

				// Registra o tempo de término
				end := time.Now()

				// Calcula a duração da execução
				duration := end.Sub(start)

				// Imprime a duração
				fmt.Println("Tempo de execução:", duration)

	
//----------------------------------------------------------
//Essa parte e responsavel por escrever no arquivo de saida.(teste rapido)
//----------------------------------------------------------
	// Abre o arquivo
	arquivo, err := os.Open("arquivoteste.txt")
	if err != nil {
			fmt.Println("Erro ao abrir o arquivo:", err)
			return
	}
	defer arquivo.Close()

	// Verificar se a primeira linha está vazia
	info, err := arquivo.Stat()
	if err != nil {
			fmt.Println("Erro ao obter informações do arquivo:", err)
			return
	}
	if info.Size() == 0 {
		
		arquivo, err := os.OpenFile("arquivoteste.txt", os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("Erro ao abrir o arquivo para escrita:", err)
				return
		}
		defer arquivo.Close()
//adiciona na primeira linha dur
		_, err = arquivo.WriteString(duration.String()+ "\n" )
		if err != nil {
				fmt.Println("Erro ao escrever no arquivo:", err)
				return
		}
			fmt.Println("Duração adicionada à primeira linha.")
			return
	}

	// Verificar a próxima linha
	linha, err4 := bufio.NewReader(arquivo).ReadString('\n')
	if err4 != nil {
			fmt.Println("Erro ao ler a próxima linha:", err4)
			return
	}
	if linha == "\n" {
			// Adicionar a duração à próxima linha
			arquivo, err := os.OpenFile("arquivoteste.txt", os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
					fmt.Println("Erro ao abrir o arquivo para escrita:", err)
					return
			}
			defer arquivo.Close()
			//adiciona o valor na proxima linha
			_, err = arquivo.WriteString(duration.String() + "\n")
			if err != nil {
					fmt.Println("Erro ao escrever no arquivo:", err)
					return
			}
			fmt.Println("Duração adicionada à próxima linha.")
	}
			
}

