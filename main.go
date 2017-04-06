package main

import (
  "fmt"
  "time"
  "os"
  "bufio"
  "log"
  "regexp"
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
var lista_de_fantasmas []Fantasma
var quantidade_de_fantasmas int


func construirLabirinto(nomeArquivo string) (*Labirinto, error) {

  var ErrMapNotFound = errors.New("Não conseguiu ler o arquivo do mapa")

  if file, err := os.Open("./data/mapa.txt"); err == nil {

    // fecha depois de ler o arquivo
    defer file.Close()

    // inicializa o mapa vazio
    mapa := []string{}

    r, _ := regexp.Compile("[^ #]")

    // cria um leitor para ler linha a linha o arquivo
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      linha := scanner.Text()
      linha = r.ReplaceAllString(linha, " ")
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
const ESC = "\x1b"

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

func criarFantasmas() {
  /* Valor inicial utilizado para posicionar um fantasma ao lado do outro */
  var contador = 10 // hard coded
  lista_de_fantasmas = make([]Fantasma, 2) // mudar para quantidade_de_fantasmas

  for i := 0; i < 2; i++ { // mudar para quantidade_de_fantasmas
      fantasma := new(Fantasma)
      fantasma.posicao.linha = 5 // hard coded
      fantasma.posicao.coluna = contador

      lista_de_fantasmas = append(lista_de_fantasmas, *fantasma)
      contador += 1
      fmt.Printf("Fantasma %d: (%d, %d)\n", i, fantasma.posicao.linha, fantasma.posicao.coluna)
  }
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
        Sai
)

func entradaDoUsuario() Movimento {
  // Eduardo
  return leTeclado()
}

func moverPacGo() bool {
  // Atualiza posição do usuário
  switch entradaDoUsuario() {
  case Cima:
  case Baixo:
  case Direita:
  case Esquerda:
  case Sai: return true
  default:
  }
  return false
}

func moverFantasmas() {
  for i := 0; i < 2; i++{ // mudar para quantidade_de_fantasmas
    lista_de_fantasmas[i].posicao.coluna += 1
  }
}

func dorme() {
  time.Sleep(time.Millisecond * 500) // 1s
}

func main() {
//  defer mostraCursor()
  escondeCursor()

  inicializa()

  pacgo     = PacGo{ posicao: Posicao{2, 2}, figura: 'G'}

  quantidade_de_fantasmas = 2

  lista_de_fantasmas = []Fantasma{
    { posicao: Posicao{2, 4}, figura:'F'},
    { posicao: Posicao{1, 6}, figura:'F'},
  }

  labirinto, _ = construirLabirinto("")
  labirinto.imprime()

  criarFantasmas()

  // TODO: Loop do jogo
  for  {
    moverFantasmas()

    dorme()

    atualizarLabirinto()

    if (moverPacGo()) {
      break
    }

    if detectarColisao() {
      terminarJogo()
    }
  }
}
