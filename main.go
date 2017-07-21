package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Entrada int

// Possiveis entradas do usuário
const (
	Nada = iota
	ParaCima
	ParaBaixo
	ParaEsquerda
	ParaDireita
	SairDoJogo // Tecla ESC
)

type PacGo struct {
	posicao    Posicao
	posInicial Posicao
	figura     string
	pontos     int
	vidas      int
	pilula     bool
}

type Fantasma struct {
	posicao Posicao
	figura  string
}

type Labirinto struct {
	largura      int
	altura       int
	mapa         []string
	numPastilhas int
}

var pacgo PacGo
var fantasmas []*Fantasma
var labirinto Labirinto

func criarPacGo(pos Posicao, fig string) {
	pacgo = PacGo{
		posicao:    pos,
		posInicial: pos,
		figura:     fig,
		vidas:      2,
	}
}

func criarFantasma(posicao Posicao, figura string) {
	fantasma := &Fantasma{posicao: posicao, figura: figura}
	fantasmas = append(fantasmas, fantasma)
}

func leEntradaDoUsuario(canal chan<- Entrada) {
	array := make([]byte, 10)

	for {
		lido, _ := os.Stdin.Read(array)

		if lido == 1 && array[0] == 0x1b {
			canal <- SairDoJogo
		} else if lido == 3 {
			if array[0] == 0x1b && array[1] == '[' {
				switch array[2] {
				case 'A':
					canal <- ParaCima
				case 'B':
					canal <- ParaBaixo
				case 'C':
					canal <- ParaDireita
				case 'D':
					canal <- ParaEsquerda
				}
			}
		}
	}
}

func inicializarLabirinto(arquivo string) error {
	var ErrMapNotFound = errors.New("Não conseguiu ler o arquivo do mapa")

	// aplica o valor default caso não seja passado um arquivo
	var tmpArquivo string
	tmpArquivo = arquivo
	if tmpArquivo == "" {
		tmpArquivo = "./mapas/mapa01.txt"
	}

	// abre arquivo
	file, err := os.Open(tmpArquivo)
	if err != nil {
		log.Fatal(err)
		return ErrMapNotFound
	}
	// fecha depois de ler o arquivo
	defer file.Close()

	// inicializa o mapa vazio
	mapa := []string{}

	// cria um leitor para ler linha a linha o arquivo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := scanner.Text()
		mapa = append(mapa, linha)
	}

	// determina o tamanho do mapa baseado no arquivo lido
	largura := len(mapa[0])
	altura := len(mapa)

	labirinto = Labirinto{
		largura: largura,
		altura:  altura,
		mapa:    mapa,
	}

	// Processa caracteres especiais
	for linha, linhaMapa := range labirinto.mapa {
		for coluna, caractere := range linhaMapa {
			switch caractere {
			case 'G':
				criarPacGo(Posicao{linha, coluna}, "\xF0\x9F\x98\x83")
			case 'F':
				criarFantasma(Posicao{linha, coluna}, "\xF0\x9F\x91\xBB")
			case '.':
				labirinto.numPastilhas++
			}
		}
	}

	return nil
}

func desenhaTela() {
	LimpaTela()

	// Imprime placar
	MoveCursor(Posicao{0, 0})
	placar := fmt.Sprintf("Pontos: %d Vidas: %d", pacgo.pontos, pacgo.vidas)
	fmt.Println(Vermelho(Intenso(placar)))

	// Ajuste para desenhar o mapa embaixo do placar
	deslocamento := Posicao{2, 0}
	MoveCursor(deslocamento)

	// Imprime mapa
	for _, linha := range labirinto.mapa {
		for _, char := range linha {
			switch char {
			case '#':
				if pacgo.pilula {
					fmt.Print(FundoVermelho(" "))
				} else {
					fmt.Print(FundoAzul(" "))
				}
			case '.':
				fmt.Print(".")
			case 'P':
				fmt.Print("\xF0\x9F\x8D\x84")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

	// Imprime PacGo
	MoveCursor(pacgo.posicao.Soma(deslocamento))
	fmt.Printf("%s", pacgo.figura)

	// Imprime fantasmas
	for _, fantasma := range fantasmas {
		MoveCursor(fantasma.posicao.Soma(deslocamento))
		fmt.Printf("%s", fantasma.figura)
	}

	// Move cursor para fora do labirinto
	MoveCursor(deslocamento.Soma(Posicao{labirinto.altura + 2, 0}))
}

func moverPacGo(m Entrada) {
	var novaLinha = pacgo.posicao.linha
	var novaColuna = pacgo.posicao.coluna

	switch m {
	case ParaCima:
		novaLinha--
		if novaLinha < 0 {
			novaLinha = labirinto.altura - 1
		}
	case ParaBaixo:
		novaLinha++
		if novaLinha >= labirinto.altura {
			novaLinha = 0
		}
	case ParaDireita:
		novaColuna++
		if novaColuna >= labirinto.largura {
			novaColuna = 0
		}
	case ParaEsquerda:
		novaColuna--
		if novaColuna < 0 {
			novaColuna = labirinto.largura - 1
		}
	}

	conteudoDoMapa := labirinto.mapa[novaLinha][novaColuna]
	if conteudoDoMapa != '#' {
		pacgo.posicao.linha = novaLinha
		pacgo.posicao.coluna = novaColuna

		if conteudoDoMapa == '.' || conteudoDoMapa == 'P' {
			switch conteudoDoMapa {
			case '.':
				pacgo.pontos += 10
				labirinto.numPastilhas--
			case 'P':
				pacgo.pontos += 100
				ativarPilula()
			}

			// Remove item do mapa
			linha := labirinto.mapa[novaLinha]
			linha = linha[:novaColuna] + " " + linha[novaColuna+1:]
			labirinto.mapa[novaLinha] = linha
		}
	}
}

func moverFantasmas() {
	for _, fantasma := range fantasmas {
		// gera um número aleatório entre 0 e 4 (Direita = 3)
		var direcao = rand.Intn(ParaDireita + 1)

		var novaPosicao = fantasma.posicao

		// Atualiza posição testando os limites do mapa
		switch direcao {
		case ParaCima:
			novaPosicao.linha--
			if novaPosicao.linha < 0 {
				novaPosicao.linha = labirinto.altura - 1
			}
		case ParaBaixo:
			novaPosicao.linha++
			if novaPosicao.linha > labirinto.altura-1 {
				novaPosicao.linha = 0
			}
		case ParaEsquerda:
			novaPosicao.coluna--
			if novaPosicao.coluna < 0 {
				novaPosicao.coluna = labirinto.largura - 1
			}
		case ParaDireita:
			novaPosicao.coluna++
			if novaPosicao.coluna > labirinto.largura-1 {
				novaPosicao.coluna = 0
			}
		}

		// Verifica se a posição nova do mapa é válida
		conteudoMapa := labirinto.mapa[novaPosicao.linha][novaPosicao.coluna]
		if conteudoMapa != '#' {
			fantasma.posicao = novaPosicao
		}
	}
}

func dorme(milisegundos time.Duration) {
	time.Sleep(time.Millisecond * milisegundos)
}

func detectarColisao() *Fantasma {
	for _, fantasma := range fantasmas {
		if fantasma.posicao == pacgo.posicao {
			return fantasma
		}
	}
	return nil
}

func main() {
	// Inicializar terminal
	Inicializa()
	defer Finaliza()

	// Lê parâmetros do sistema operacional
	args := os.Args[1:]
	var arquivo string
	if len(args) >= 1 {
		arquivo = args[0]
	} else {
		arquivo = ""
	}

	// Inicializar labirinto
	inicializarLabirinto(arquivo)

	// Cria rotina para ler entradas
	canal := make(chan Entrada, 10)
	go leEntradaDoUsuario(canal)

	// Loop principal
	for {
		// Desenha tela
		desenhaTela()

		// Fim de jogo
		if labirinto.numPastilhas == 0 {
			MoveCursor(Posicao{labirinto.altura + 3, 0})
			fmt.Print("Fim de jogo! Você venceu! \xF0\x9F\x98\x84\n\n")
			break
		}

		// Processa entrada do jogador
		var tecla Entrada
		select {
		case tecla = <-canal:
		default:
		}

		if tecla == SairDoJogo {
			break
		} else {
			moverPacGo(tecla)
		}

		// Processa movimento dos fantasmas
		moverFantasmas()

		// Processa colisões
		if fantasma := detectarColisao(); fantasma != nil {
			if pacgo.pilula {
				go ressucitarFantasma(fantasma, 10000) // 10 segundos
				matarFantasma(fantasma)
				pacgo.pontos += 500
			} else {
				pacgo.vidas--

				// Reseta posição do PacGopher para a posição inicial
				pacgo.posicao = pacgo.posInicial
			}

			if pacgo.vidas < 0 {
				MoveCursor(Posicao{labirinto.altura + 3, 0})
				fmt.Print("Fim de jogo! Os fantasmas venceram... \xF0\x9F\x98\xAD\n\n")
				break
			}
		}

		// Dorme
		dorme(100)
	}
}

func ativarPilula() {
	pacgo.pilula = true
	go desativarPilula(10000) // 10 segundos
}

func desativarPilula(milisegundos time.Duration) {
	dorme(milisegundos)
	pacgo.pilula = false
}

func ressucitarFantasma(fantasma *Fantasma, milisegundos time.Duration) {
	dorme(milisegundos)
	criarFantasma(fantasma.posicao, fantasma.figura)
}

func matarFantasma(fantasma *Fantasma) {
	var indice int
	for idx, f := range fantasmas {
		if f.posicao == fantasma.posicao {
			indice = idx
			break
		}
	}
	fantasmas = append(fantasmas[:indice], fantasmas[indice+1:]...)
}
