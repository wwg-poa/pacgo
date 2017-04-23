package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

// Tela é a struct que define uma tela
type Tela struct{}

var tela Tela

func inicializa() {
	rawMode := exec.Command("/bin/stty", "cbreak", "-echo")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()

	tela = Tela{}
}

func finaliza() {
	rawMode := exec.Command("/bin/stty", "-cbreak", "echo")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()
}

func (posicao1 *Posicao) adiciona(posicao2 *Posicao) Posicao {
	return Posicao{posicao1.linha + posicao2.linha, posicao1.coluna + posicao2.coluna}
}

func (pacgo *PacGo) incrementaIndice() {
	if pacgo.contadorFig.contador == pacgo.contadorFig.max {
		if pacgo.indiceFig == len(pacgo.figuras)-1 {
			pacgo.indiceFig = 0
		} else {
			pacgo.indiceFig++
		}
		pacgo.contadorFig.contador = 0
	} else {
		pacgo.contadorFig.contador++
	}

}

// ESC é a string que define o comando escape
const ESC = "\x1b"

func (tela *Tela) limpa() {
	fmt.Printf("%s[2J", ESC)
	tela.moveCursor(Posicao{0, 0})
}

func (tela *Tela) moveCursor(p Posicao) {
	fmt.Printf("%s[%d;%df", ESC, p.linha+1, p.coluna+1)
}

func vermelho(s string) string      { return ansi(31, s) }
func verde(s string) string         { return ansi(32, s) }
func azul(s string) string          { return ansi(34, s) }
func fundoVermelho(s string) string { return ansi(41, s) }
func fundoVerde(s string) string    { return ansi(42, s) }
func fundoAzul(s string) string     { return ansi(44, s) }
func intenso(s string) string       { return ansi(1, s) }

var ansiRE *regexp.Regexp

func ansi(code int, s string) string {
	if ansiRE == nil {
		ansiRE = regexp.MustCompile(`^` + ESC + `\[(\d+(?:;\d+)*m.*` + ESC + `\[0m)$`)
	}

	parts := ansiRE.FindStringSubmatch(s)
	if parts == nil {
		return fmt.Sprintf("%s[%dm%s%s[0m", ESC, code, s, ESC)
	}
	return fmt.Sprintf("%s[%d;%s", ESC, code, parts[1])
}
