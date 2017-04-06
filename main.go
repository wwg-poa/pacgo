package main

import "fmt"

type Posicao struct {
  linha  int
  coluna int
}

type PacGo struct {
  posicao Posicao
  figura  rune // emoji
}

type Fantasma struct {
  posicao Posicao
  figura  rune // emoji
}

type Labirinto struct {
  largura int
  altura  int
  mapa    []string
}

func construirLabirinto(nomeArquivo string) {
  // TODO: carregar arquivo de mapa
  // TODO: determinar dimensao da tela
  // Julia
}

func atualizarLabirinto() {
  // TODO: imprime o labirinto
  // TODO: imprime pacgo na posição x,y
  // TODO: imprime fantasmas
  // Dani
}

func detectarColisao() bool {
  // TODO: posição do pacgo == posição de algum fantasma?
  // Ação ?
  return false
}

func terminarJogo() {
  // pacgo morreu :(
}

type Movimento int

const (
        Cima = iota
        Baixo
        Esquerda
        Direita
        Nenhum
)

func entradaDoUsuario() Movimento {
  // Eduardo
  // Lê teclado
  return Nenhum
}

func moverPacGo() {
  // Atualiza posição do usuário
  switch entradaDoUsuario() {
  case Cima:
  case Baixo:
  case Direita:
  case Esquerda:
  default:
  }
}

func moverFantasmas() {
  // Isa
}

func dorme() {
  // TODO: sleep
}

func main() {

  fmt.Println("Hello pac go!")

  construirLabirinto()

  // TODO: Loop do jogo
  for  {
    dorme()

    atualizarLabirinto()

    moverPacGo()

    moverFantasmas()

    if detectarColisao() {
      terminarJogo()
    }
  }

}
