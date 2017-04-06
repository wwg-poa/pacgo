package main

import (
  "fmt"
  "time"
)

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

var labirinto Labirinto
var pacgo     PacGo
var fantasmas []Fantasma

const ESC = "\x1b"

func construirLabirinto(nomeArquivo string) {
  // TODO: carregar arquivo de mapa
  // TODO: determinar dimensao da tela
  // Julia
}

func limpaTela() {
  fmt.Printf("%s[2J", ESC)
  moveCursor(Posicao{1,1})
}

func moveCursor(p Posicao) {
  fmt.Printf("%s[%d;%df", ESC, p.linha, p.coluna)
}

func escondeCursor() {
  fmt.Printf("%s?25l", ESC)
}

func mostraCursor() {
  fmt.Printf("%s?25h", ESC)
}

func atualizarLabirinto() {
  limpaTela()
  for _, linha := range labirinto.mapa {
    fmt.Println(linha)
  }

  // Atualiza PacGo
  moveCursor(pacgo.posicao)
  fmt.Printf("%c", pacgo.figura)

  // TODO: imprime fantasmas
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
  time.Sleep(time.Millisecond * 500) // 1s
}

func main() {

//  defer mostraCursor()
  escondeCursor()

  pacgo     = PacGo{ posicao: Posicao{2, 2}, figura: 'G'}

  fantasmas = []Fantasma{
    { posicao: Posicao{2, 4}, figura:'F'},
    { posicao: Posicao{1, 6}, figura:'F'},
  }

  labirinto =  Labirinto{ largura:10, altura:4, mapa:[]string {"#### #####",
                                                               "#        #",
                                                               "          ",
                                                               "#### #####"}}

  construirLabirinto("")

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
