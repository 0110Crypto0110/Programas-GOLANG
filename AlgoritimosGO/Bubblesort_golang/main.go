package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

//-------------------------------------------------------------
// Bubble_sort
//-------------------------------------------------------------
// Serve para ordenar os elementos de uma lista, levando o maior valor para o final da lista
//-------------------------------------------------------------
// funcoes adicionais para mensurar tempo e uso de memoria
//-------------------------------------------------------------
func bubbleSort(arr []int) ([]int, uint64, float64) {
	inicio := time.Now()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	memoriaInicial := memStats.Alloc

	for n := len(arr) - 1; n > 0; n-- {
		troca := false
		for i := 0; i < n; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				troca = true
			}
		}
		if !troca {
			break
		}
	}

	fim := time.Now()
	tempoExecucao := fim.Sub(inicio).Seconds()

	runtime.ReadMemStats(&memStats)
	memoriaFinal := memStats.Alloc
	memoriaUsada := memoriaFinal - memoriaInicial

	return arr, memoriaUsada, tempoExecucao
}

//-------------------------------------------------------------
// lerArquivoEConverterParaArray
//-------------------------------------------------------------
// Serve para ler os valores de um arquivo de texto e convertê-los em uma lista de inteiros
//-------------------------------------------------------------
func lerArquivoEConverterParaArray(caminhoArquivoEntrada string) []int {
	file, err := os.Open(caminhoArquivoEntrada)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var arr []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			panic(err)
		}
		arr = append(arr, num)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return arr
}

//-------------------------------------------------------------
// escreverArrayEmArquivo
//-------------------------------------------------------------
// Serve para escrever os valores de uma lista ordenada em um arquivo de texto
//-------------------------------------------------------------
func escreverArrayEmArquivo(arr []int, caminhoArquivoSaida string) {
	file, err := os.Create(caminhoArquivoSaida)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("Lista ordenada: %v\n", arr))
	writer.Flush()
}

//-------------------------------------------------------------
// escreverDadosEmArquivo
//-------------------------------------------------------------
// Serve para escrever o uso de memória e tempo de execução em um arquivo de texto
//-------------------------------------------------------------
func escreverDadosEmArquivo(memoriaUsada uint64, tempoExecucao float64, caminhoArquivoSaida string) {
	file, err := os.OpenFile(caminhoArquivoSaida, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("%d\n", memoriaUsada))
	writer.WriteString(fmt.Sprintf("%f\n", tempoExecucao))
	writer.Flush()
}

//-------------------------------------------------------------
// lerLinhas
//-------------------------------------------------------------
// Serve para ler todas as linhas de um arquivo de texto e retorná-las como uma lista de strings
//-------------------------------------------------------------
func lerLinhas(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

//-------------------------------------------------------------
// obterValoresLinhasPares
//-------------------------------------------------------------
// Serve para obter os valores das linhas pares de uma lista de strings e retorná-los como uma lista de floats
//-------------------------------------------------------------
func obterValoresLinhasPares(linhas []string) []float64 {
	var valores []float64
	for i, linha := range linhas {
		if i%2 == 0 {
			valor, err := strconv.ParseFloat(strings.TrimSpace(linha), 64)
			if err != nil {
				panic(err)
			}
			valores = append(valores, valor)
		}
	}
	return valores
}

//-------------------------------------------------------------
// obterValoresLinhasImpares
//-------------------------------------------------------------
// Serve para obter os valores das linhas ímpares de uma lista de strings e retorná-los como uma lista de floats
//-------------------------------------------------------------
func obterValoresLinhasImpares(linhas []string) []float64 {
	var valores []float64
	for i, linha := range linhas {
		if i%2 != 0 {
			valor, err := strconv.ParseFloat(strings.TrimSpace(linha), 64)
			if err != nil {
				panic(err)
			}
			valores = append(valores, valor)
		}
	}
	return valores
}

//-------------------------------------------------------------
// calcularMediaEMediana
//-------------------------------------------------------------
// Serve para calcular a média e a mediana de uma lista de valores floats
//-------------------------------------------------------------
func calcularMediaEMediana(valores []float64) (float64, float64) {
	soma := 0.0
	for _, valor := range valores {
		soma += valor
	}
	media := soma / float64(len(valores))

	sort.Float64s(valores)
	var mediana float64
	if len(valores)%2 == 0 {
		mediana = (valores[len(valores)/2-1] + valores[len(valores)/2]) / 2
	} else {
		mediana = valores[len(valores)/2]
	}

	return media, mediana
}

//-------------------------------------------------------------
// gerarGraficos
//-------------------------------------------------------------
// Serve para gerar gráficos de dispersão com os valores de memória usada e tempo de execução
//-------------------------------------------------------------
func gerarGraficos(memoriaUsada, tempoExecucao []float64, mediaMemoria, medianaMemoria, mediaTempo, medianaTempo float64) {
	p := plot.New()

	p.Title.Text = "Uso de Memória e Tempo de Execução"
	p.X.Label.Text = "Execuções"
	p.Y.Label.Text = "Valores"

	pointsMemoria := make(plotter.XYs, len(memoriaUsada))
	pointsTempo := make(plotter.XYs, len(tempoExecucao))
	for i := range memoriaUsada {
		pointsMemoria[i].X = float64(i)
		pointsMemoria[i].Y = memoriaUsada[i]
		pointsTempo[i].X = float64(i)
		pointsTempo[i].Y = tempoExecucao[i]
	}

	scatterMemoria, err := plotter.NewScat
