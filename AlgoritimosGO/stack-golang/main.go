package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type (// define a estrutura de uma pilha
	stack struct {
		TOPO    *No 
		TAMANHO int
	}
	No struct { // define a estrutura do no, atributos (valor e anterior)
		valor interface{}
		anterior  *No
	}
)

// Cria uma nova pilha
func nova() *stack {
	return &stack{nil, 0}
}
//____________________________________________________________
//metodo tamanho
//____________________________________________________________
//-Este método retorna o tamanho da pilha.
//-É útil quando você precisa verificar o tamanho da pilha antes de fazer operações.
//-A complexidade de tempo é O(1).
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) tamanho() int {
	return pilha.TAMANHO
}

//____________________________________________________________
//metodo insere
//____________________________________________________________
//-Este método insere um valor no topo da pilha.
//-É util quando você precisa adicionar um novo elemento no topo da pilha.
//-A complexidade de tempo é O(1).
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) insere(valor interface{}) {
	n := &No{valor, pilha.TOPO}
	pilha.TOPO = n
	pilha.TAMANHO++
}
//____________________________________________________________
//metodo apaga pilha
//____________________________________________________________
//-Este método apaga a pilha, liberando a memória ocupada por seus nos.
//-É útil quando você precisa liberar a memória ocupada pela pilha.
//-A complexidade de tempo é O(n), onde n é o tamanho da pilha.
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) apaga(){
	for pilha.tamanho() > 0{
		pilha.remove()
	}
}

//____________________________________________________________
//metodo remover valor da pilha
//____________________________________________________________
//-Este método remove um valor em uma posicao especifica, removendo os valores acima dele. apagando o valor e reinserindo os valores acima.
//-É útil quando você precisa remover um valor em uma posição específica da pilha, mantendo a logica de remover todos os valores anteriores.
//-A complexidade de tempo é O(n), onde n é o tamanho da pilha.
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) desempilhaREEMPILHA(valor interface{}) {
	// Encontra o índice do valor na pilha
	indice := pilha.recuperaINDICE(valor)
	if indice == -1 {
		fmt.Println("Valor não encontrado na pilha.")
		return
	}

	novaPilha := nova()
	// Desempilha até o valor desejado e insere os remanescentes na nova pilha
	for pilha.tamanho() > indice+1 {
		novaPilha.insere(pilha.remove())
	}
	pilha.remove()
	// Reempilha os valores restantes na pilha original
	for novaPilha.tamanho() > 0 {
		pilha.insere(novaPilha.remove())
	}
}

//____________________________________________________________
//metodo imprime topo
//____________________________________________________________
//-Este método imprime o valor do topo da pilha.
//-É útil quando você precisa verificar o valor do topo da pilha.
//-A complexidade de tempo é O(1).
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) topo() interface{} {
	if pilha.tamanho() == 0 {
		return nil
	}
	return pilha.TOPO.valor
}

//____________________________________________________________
//metodo remove topo
//____________________________________________________________
//-Este método remove o valor do topo da pilha.
//-É útil quando você precisa remover o valor do topo da pilha.
//-A complexidade de tempo é O(1).
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) remove() interface{} {
	if pilha.tamanho() == 0 {
		return nil 
	}

	valorRemovido := pilha.TOPO.valor 
	pilha.TOPO = pilha.TOPO.anterior 
	pilha.TAMANHO-- 
	return valorRemovido 
}

//____________________________________________________________
//metodo imprime todos os valores da pilha
//____________________________________________________________
//-Este método imprime todos os valores da pilha.
//-É útil quando você precisa verificar todos os valores da pilha.
//-A complexidade de tempo é O(n), onde n é o tamanho da pilha.
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) imprime() {
	for i := pilha.TOPO; i != nil; i = i.anterior {
		fmt.Println(i.valor)
	}
}

//____________________________________________________________
//metodo insere valor numa posicao especifica
//____________________________________________________________
//-Este método insere um valor em uma posição específica da pilha.
//-É útil quando você precisa inserir um valor em uma posição específica da pilha.
//-A complexidade de tempo é O(n), onde n é o tamanho da pilha
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) inserePosicao(valor interface{}, posicao int) {
	if posicao < 0 || posicao > pilha.tamanho() {
		fmt.Println("Posição inválida.")
		return
	}

	// Cria uma nova pilha para armazenar os valores desempilhados
	novaPilha := nova()

	// Desempilha os valores até a posição desejada
	for i := 0; i < posicao; i++ {
		novaPilha.insere(pilha.remove())
	}

	// Insere o novo valor na posição especificada
	pilha.insere(valor)

	// Reempilha os valores desempilhados
	for novaPilha.tamanho() > 0 {
		pilha.insere(novaPilha.remove())
	}
}

//____________________________________________________________
//metodo que a partir de um indice retorna o valor
//____________________________________________________________
//-Este método retorna o valor armazenado em uma posição específica da pilha.
//-É útil quando você precisa obter o valor armazenado em uma posição específica da pilha.
//-A complexidade de tempo é O(n), onde n é o tamanho da pilha.
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) recuperaVALOR(indice int) interface{} {
	if indice < 0 || indice >= pilha.tamanho() {
		return nil
	}

	noAtual := pilha.TOPO
	for i := 0; i < indice; i++ {
		noAtual = noAtual.anterior
	}
	return noAtual.valor
}

//____________________________________________________________
//metodo que a partir do valor retorna o indice
//____________________________________________________________
//-Este método retorna o índice da primeira ocorrência do valor na pilha.
//-É útil quando você precisa encontrar o índice da primeira ocorrência de um valor na pilha.
//-A complexidade de tempo é O(n), onde n é o tamanho da pilha.
//-A complexidade de espaço é O(1).
//____________________________________________________________
func (pilha *stack) recuperaINDICE(valor interface{}) int {
	noAtual := pilha.TOPO
	indice := 0
	for noAtual != nil {
		if noAtual.valor == valor {
			return indice
		}
		noAtual = noAtual.anterior
		indice++
	}
	return -1 // Retorna -1 se o valor não for encontrado na pilha
}
//____________________________________________________________

//____________________________________________________________
//exemplo de funcionamento por meio de interface, menu no terminal.
//____________________________________________________________

//____________________________________________________________


func main() {
	//defnindo as opcoes do menu
	stack := nova()//criando a pilha
	fmt.Println("-------------------------------")
	fmt.Println("0: para lista de metodos ")
	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")
	fmt.Println("selecione o método: ")
	fmt.Println("1: insere valor ")
	fmt.Println("2: remove o topo ")
	fmt.Println("3: tamanho ")
	fmt.Println("4: topo da pilha ")
	fmt.Println("5: imprime ")
	fmt.Println("6: inserePosicao ")
	fmt.Println("7: recupera INDICE")
	fmt.Println("8: recupera VALOR  ")
	fmt.Println("9: remove valor(desempilha)  ")
	fmt.Println("10: remove toda a pilha)  ")
	fmt.Println("qualquer outro valor: encerra o programa ")
	fmt.Println("-------------------------------")

	scanner := bufio.NewScanner(os.Stdin)//criando o scanner
	for scanner.Scan() {
		input := scanner.Text()//pegando o input do usuario
		num, err := strconv.Atoi(input)//converte o input para int
		if err != nil {
			fmt.Println("Erro ao converter o valor para inteiro:", err)
			continue
		}
		fmt.Println("Você digitou:", num)
		if num == 0 {
			fmt.Println("-------------------------------")
			fmt.Println("selecione o método: ")
			fmt.Println("1: insere valor ")
			fmt.Println("2: remove o topo ")
			fmt.Println("3: tamanho ")
			fmt.Println("4: topo da pilha ")
			fmt.Println("5: imprime ")
			fmt.Println("6: inserePosicao ")
			fmt.Println("7: recupera INDICE")
			fmt.Println("8: recupera VALOR  ")
			fmt.Println("9: remove valor(desempilha)  ")
			fmt.Println("10: remove toda a pilha)  ")
			fmt.Println("qualquer outro valor: encerra o programa ")
			fmt.Println("-------------------------------")
		} else if num == 1 {
				fmt.Println("Digite o valor a ser inserido na pilha:")
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					input1 := scanner.Text()
					num1, err := strconv.Atoi(input1)
					if err != nil {
						fmt.Println("Erro ao converter o valor para inteiro:", err)
						continue
					}
					(*stack).insere(num1)//chamando o metodo insere
					break 
				}
			fmt.Println("-------------------------------")
			fmt.Println("selecione o método: ")

		} else if num == 2 {
			fmt.Println("removendo topo:")
				
				(*stack).remove()//chamando o metodo remove topo
			
			fmt.Println("Selecione o método:")

		}else if num == 3 {
			fmt.Println("-------------------------------")
			fmt.Println("tamanho da pilha")
			fmt.Println((*stack).tamanho())//chamando o metodo tamanho
			fmt.Println("-------------------------------")
			fmt.Println("selecione o método: ")	

		}else if num == 4 {
			fmt.Println("-------------------------------")
			topo1 := (*stack).topo()//chamando o metodo topo
			fmt.Println(topo1)
			fmt.Println("-------------------------------")
			fmt.Println("Selecione o método:")	

		}else if num == 5 {
			fmt.Println("-------------------------------")
			fmt.Println("pilha:")
			(*stack).imprime()//chamando o metodo imprime
			fmt.Println("-------------------------------")
			fmt.Println("selecione o método: ")	

		}else if num == 6 {
				scanner1 := bufio.NewScanner(os.Stdin)
				fmt.Println("Digite o valor ")
				if scanner1.Scan() {
						input1 := scanner1.Text() 
						num1, err1 := strconv.Atoi(input1)
						if err1 != nil {
								fmt.Println("Erro ao converter o valor para inteiro:", err1)
								continue
						}

						scanner2 := bufio.NewScanner(os.Stdin) 
						fmt.Println("Digite a posicao:")
						if scanner2.Scan() {
								input2 := scanner2.Text() 
								pos, err2 := strconv.Atoi(input2)
								if err2 != nil {
										fmt.Println("Erro ao converter a posicao para inteiro:", err2)
										continue
								}

								(*stack).inserePosicao(num1, pos)//chamando o metodo inserePosicao
								fmt.Println("-------------------------------")
								fmt.Println("Selecione o método:")
						}
				}
		}else if num == 7 {
				fmt.Println("Digite o valor para saber seu índice:")
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
						input1 := scanner.Text()
						num1, err := strconv.Atoi(input1)
						if err != nil {
								fmt.Println("Erro ao converter o valor para inteiro:", err)
								return
						}
						indice := stack.recuperaINDICE(num1) // Chamando o método recuperaINDICE
						if indice == -1 {
								fmt.Println("Erro: valor não encontrado na pilha")
						} else {
								fmt.Println("O valor", num1, "está no índice", indice)
						}
						fmt.Println("-------------------------------")
						fmt.Println("Selecione o método:")
				}
		}else if num == 8 {
					fmt.Println("Digite o indice para saber seu valor:")
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						input1 := scanner.Text()
						num1, err := strconv.Atoi(input1)
						if err != nil {
							fmt.Println("Erro ao converter o valor para inteiro:", err)
							continue
						}
						valor := stack.recuperaVALOR(num1)//chamando o metodo recuperaVALOR
						if valor == nil {
								fmt.Println("Erro: indice não encontrado na pilha")
						} else {
						fmt.Println("Valor no índice", num1, ":", valor)
						}
					
					fmt.Println("-------------------------------")
					fmt.Println("selecione o método: ")
					break
					}
				

			}else if num == 9 {
				fmt.Println("Digite o valor a ser removido na pilha:")
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						input1 := scanner.Text()
						num1, err := strconv.Atoi(input1)
						if err != nil {
							fmt.Println("Erro ao converter o valor para inteiro:", err)
							continue
						}
					(*stack).desempilhaREEMPILHA(num1)	//chamando o metodo desempilhaREEMPILHA
					fmt.Println("-------------------------------")
					fmt.Println("selecione o método: ")
					break
					}
						
			}else if num == 10 {
			fmt.Println("removendo todos os valores:")

				(*stack).apaga()//chamando o metodo apaga
			fmt.Println("-------------------------------")
			fmt.Println("Selecione o método:")

		}else {
				fmt.Println("metodo invalido, fim do processo")
				break
			}
		}
	}

