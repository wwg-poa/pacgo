package main

import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
)

/* Coordenadas na tela */
type Posicao struct {
  linha int
  coluna int
}

/* Aplica um deslocamento na posição. */
func (pos1 *Posicao) Soma(pos2 Posicao) Posicao {
	return Posicao{pos1.linha + pos2.linha, pos1.coluna + pos2.coluna}
}

/* Inicializa prepara o terminal para funcionar como base do jogo. */
func Inicializa() {
  rawMode := exec.Command("/bin/stty", "cbreak", "-echo")
  rawMode.Stdin = os.Stdin
  _ = rawMode.Run()
  rawMode.Wait()
}

/* Finaliza restaura o terminal para o estado original. */
func Finaliza() {
  rawMode := exec.Command("/bin/stty", "-cbreak", "echo")
  rawMode.Stdin = os.Stdin
  _ = rawMode.Run()
  rawMode.Wait()
}

const ESC = "\x1b"

/* limpaTela Limpa a tela e volta o cursor para a posição inicial. */
func LimpaTela() {
  fmt.Printf("%s[2J", ESC)
  MoveCursor(Posicao{0,0})
}

/* MoveCursor Move o cursor para as coordendas definidas pela posição p */
func MoveCursor(p Posicao) {
  fmt.Printf("%s[%d;%df", ESC, p.linha + 1, p.coluna + 1)
}

func Vermelho(s string) string { return ansi(31, s) }
func Verde(s string) string { return ansi(32, s) }
func Azul(s string) string { return ansi(34, s) }
func FundoVermelho(s string) string { return ansi(41, s) }
func FundoVerde(s string) string { return ansi(42, s) }
func FundoAzul(s string) string { return ansi(44, s) }
func Intenso(s string) string { return ansi(1, s) }

var ansiRE *regexp.Regexp

func ansi(code int, s string) string {
  if ansiRE == nil {
    ansiRE = regexp.MustCompile(`^` + ESC + `\[(\d+(?:;\d+)*m.*` + ESC + `\[0m)$`)
  }

  parts := ansiRE.FindStringSubmatch(s)
  if parts == nil {
    return fmt.Sprintf("%s[%dm%s%s[0m", ESC, code, s, ESC)
  } else {
    return fmt.Sprintf("%s[%d;%s", ESC, code, parts[1])
  }

  return ""
}
