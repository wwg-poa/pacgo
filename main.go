package main

import "fmt"

type pacgo struct {
  linha  int
  coluna int
}

type labirinto struct {
  largura int
  altura  int
  mapa    []string
}

func construirLabirinto(nomeArquivo string) {
  // TODO: carregar arquivo de mapa
  // TODO: determinar dimensao da tela
}

func atualizarLabirinto() {
  // TODO: imprime o labirinto
  // TODO: imprime pacgo na posição x,y
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
  // Lê teclado
  return Nenhum
}

func moverPacgo() {
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

}

func main() {

  fmt.Println("Hello pac go!")

  construirLabirinto()

  // TODO: Loop do jogo
  for  {
    // TODO: timer (sleep)

    atualizarLabirinto()

    moverPacgo()

    moverFantasmas()

    if detectarColisao() {
      terminarJogo()
    }
  }

}
