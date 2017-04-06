package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"

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

var ErrMapNotFound = errors.New("Não conseguiu ler o arquivo do mapa")

func construirLabirinto(nomeArquivo string) (*Labirinto, error) {
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
  for i := 0; i < 2; i++{ // mudar para quantidade_de_fantasmas
    lista_de_fantasmas[i].posicao.coluna += 1
  }
}

func dorme() {
  time.Sleep(time.Second) // 1s
}

func main() {
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

    moverPacGo()


    if detectarColisao() {
      terminarJogo()
    }
  }

}
