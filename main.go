package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"

  "bufio"
  "log"
  "errors"
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

func (labirinto *Labirinto) imprime() {
  fmt.Println(labirinto.largura)
  fmt.Println(labirinto.altura)

  for _, linha := range labirinto.mapa {
    fmt.Println(linha)
  }
}

var labirinto *Labirinto
var pacgo     PacGo
var fantasmas []Fantasma

var ErrMapNotFound = errors.New("Não conseguiu ler o arquivo do mapa")

func construirLabirinto(nomeArquivo string) (*Labirinto, error) {
  // TODO: carregar arquivo de mapa
  // TODO: determinar dimensao da tela
  // Julia

  if file, err := os.Open("./data/mapa.txt"); err == nil {

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

    // verifica se teve erro o leitor
    if err = scanner.Err(); err != nil {
      log.Fatal(err)
      return nil, ErrMapNotFound
    }

    l := &Labirinto{largura: len(mapa[0]), altura: len(mapa), mapa : mapa}
    return l, nil

  } else {
    log.Fatal(err)
    return nil, ErrMapNotFound
  }
}

func limpaTela() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func atualizarLabirinto() {
  limpaTela()
  for _, linha := range labirinto.mapa {
    fmt.Println(linha)
  }


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
  time.Sleep(time.Second) // 1s
}

func main() {

  pacgo     = PacGo{ posicao: Posicao{2, 2}, figura: 'G'}

  fantasmas = []Fantasma{
    { posicao: Posicao{2, 4}, figura:'F'},
    { posicao: Posicao{1, 6}, figura:'F'},
  }

  labirinto, _ = construirLabirinto("")
  labirinto.imprime()

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
