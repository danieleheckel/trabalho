package main

import (
	"bufio"   // leitura de entrada do usuário
	"fmt"     // impressão no console
	"math"    // funções matemáticas (raiz quadrada)
	"os"      // acesso ao sistema (entrada padrão)
	"strconv" // conversão de string para número
	"strings" // manipulação de texto
)

// Struct para armazenar os coeficientes da equação (ax² + bx + c = 0)
type Equacao struct {
	A float64
	B float64
	C float64
}

// Função para ler números decimais com validação
func lerFloat(reader *bufio.Reader, mensagem string) float64 {
	for { // loop até o usuário digitar corretamente
		fmt.Print(mensagem)

		// Lê o valor digitado até o ENTER
		input, _ := reader.ReadString('\n')

		// Remove espaços e quebra de linha
		input = strings.TrimSpace(input)

		// Verifica se tem ponto decimal (ex: 2.0)
		if !strings.Contains(input, ".") {
			fmt.Println("Erro: digite um número decimal (ex: 2.0, 3.5)")
			continue // volta para o início do loop
		}

		// Converte string para float64
		valor, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Erro: valor inválido")
			continue
		}

		// Retorna o valor válido
		return valor
	}
}

func main1() {
	// Cria leitor para entrada do teclado
	reader := bufio.NewReader(os.Stdin)

	// Variável para quantidade de equações
	var repeticoes int

	// Loop para garantir que o usuário digite um número inteiro válido
	for {
		fmt.Print("Quantas equações deseja calcular? ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Converte string para inteiro
		valor, err := strconv.Atoi(input)

		// Valida se é número inteiro positivo
		if err != nil || valor <= 0 {
			fmt.Println("Erro: digite um número inteiro válido")
			continue
		}

		repeticoes = valor
		break
	}

	// Map para armazenar os resultados (chave = número da equação)
	resultados := make(map[int]string)

	// Loop para calcular várias equações
	for i := 1; i <= repeticoes; i++ {
		fmt.Printf("\n--- Equação %d ---\n", i)

		// Lê os coeficientes usando a função criada
		eq := Equacao{
			A: lerFloat(reader, "Digite o valor de A: "),
			B: lerFloat(reader, "Digite o valor de B: "),
			C: lerFloat(reader, "Digite o valor de C: "),
		}

		// Cálculo do delta: Δ = b² - 4ac
		delta := (eq.B * eq.B) - (4 * eq.A * eq.C)

		// Verifica se é equação do 2º grau
		if eq.A == 0 {
			resultados[i] = "Não é equação do 2º grau (A = 0)"
			continue
		}

		// Se delta for negativo, não há raízes reais
		if delta < 0 {
			resultados[i] = "Não possui raízes reais"
			continue
		}

		// Fórmula de Bhaskara:
		// x = (-b ± √delta) / (2a)
		x1 := (-eq.B + math.Sqrt(delta)) / (2 * eq.A)
		x2 := (-eq.B - math.Sqrt(delta)) / (2 * eq.A)

		// Armazena o resultado formatado (2 casas decimais)
		resultados[i] = fmt.Sprintf("x1 = %.2f | x2 = %.2f", x1, x2)
	}

	// Exibe todos os resultados ao final
	fmt.Println("\n===== RESULTADOS =====")

	for i := 1; i <= repeticoes; i++ {
		fmt.Printf("Equação %d: %s\n", i, resultados[i])
	}
}

/*
Observações:

math.Sqrt(delta) -> calcula a raiz quadrada do delta

%.2f -> formata número com 2 casas decimais

%s -> formata string
*/