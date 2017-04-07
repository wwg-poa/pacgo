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
  figura  string // emoji
}

type Fantasma struct {
  posicao Posicao
  figura  string // emoji
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
    //fmt.Println(linha)
    fmt.Print(linha)
    fmt.Print("\r\n")
  }
}

var labirinto *Labirinto
var pacgo     *PacGo
var lista_de_fantasmas []*Fantasma
var quantidade_de_fantasmas int

func construirLabirinto(nomeArquivo string) (*Labirinto, *PacGo, []*Fantasma, error) {

  var ErrMapNotFound = errors.New("Não conseguiu ler o arquivo do mapa")

  if file, err := os.Open("./data/mapa.txt"); err == nil {

    // fecha depois de ler o arquivo
    defer file.Close()

    // inicializa o mapa vazio
    var pacgo *PacGo//{ posicao: Posicao{2, 2}, figura: 'G'}
    fantasmas := []*Fantasma{}
    mapa := []string{}

    r, _ := regexp.Compile("[^ #]")

    // cria um leitor para ler linha a linha o arquivo
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      linha := scanner.Text()

      for indice , caracter := range linha {
        switch caracter {
          case 'F': {
            fantasma := &Fantasma{ posicao: Posicao{len(mapa), indice}, figura: "F"}
            fantasmas = append(fantasmas, fantasma)
          }
          //fmt.Println(caracter)
          case 'G': pacgo = &PacGo{ posicao: Posicao{len(mapa), indice}, figura: "G"}
        }
      }

      linha = r.ReplaceAllString(linha, " ")
      mapa = append(mapa, linha)
    }

    // verifica se teve erro o leitor
    if err = scanner.Err(); err != nil {
      log.Fatal(err)
      return nil, nil, nil, ErrMapNotFound
    }

    l := &Labirinto{largura: len(mapa[0]), altura: len(mapa), mapa : mapa}
    return l, pacgo, fantasmas, nil

  } else {
    log.Fatal(err)
    return nil, nil, nil, ErrMapNotFound
  }
}

const ESC = "\x1b"

func limpaTela() {
  fmt.Printf("%s[2J", ESC)
  moveCursor(Posicao{0,0})
}

func moveCursor(p Posicao) {
  fmt.Printf("%s[%d;%df", ESC, p.linha + 1, p.coluna + 1)
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
  fmt.Printf("%s", pacgo.figura)

  for _, fantasma := range lista_de_fantasmas {
    moveCursor(fantasma.posicao)
    fmt.Printf("%s", fantasma.figura)
  }

  // Move o cursor para fora do labirinto
  moveCursor(Posicao{labirinto.altura + 2, 1})
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
  array := make([]byte, 10)

  lido, _ := os.Stdin.Read(array)

  if lido == 1 && array[0] == 0x1b {
    return Sai
  } else if lido == 3 {
    if array[0] == 0x1b && array[1] == '[' {
      switch array[2] {
      case 'A': return Cima
      case 'B': return Baixo
      case 'C': return Direita
      case 'D': return Esquerda
      }
    }
  }

  return Nenhum
}

func moverPacGo(m Movimento) {

  var valorDaPosicaoAtualDaPacgo = labirinto.mapa[pacgo.posicao.linha][pacgo.posicao.coluna]
  var linhaAtualDaPacgo = pacgo.posicao.linha
  var colunaAtualDaPacgo = pacgo.posicao.coluna

  switch m {
  case Cima:
             if linhaAtualDaPacgo == 0{
                 if valorDaPosicaoAtualDaPacgo == ' '{
                   pacgo.posicao.linha = labirinto.altura - 1
                 }
             }else{
               var posicaoAcimaDaPacgo = labirinto.mapa[pacgo.posicao.linha - 1][pacgo.posicao.coluna]
               if posicaoAcimaDaPacgo != '#'{
                 pacgo.posicao.linha = pacgo.posicao.linha - 1
               }
             }
  case Baixo:
             if linhaAtualDaPacgo == labirinto.altura - 1{
                 if valorDaPosicaoAtualDaPacgo == ' '{
                   pacgo.posicao.linha = 0
                 }
             }else{
               var posicaoAbaixoDaPacgo = labirinto.mapa[pacgo.posicao.linha + 1][pacgo.posicao.coluna]
               if posicaoAbaixoDaPacgo != '#'{
                 pacgo.posicao.linha = pacgo.posicao.linha + 1
               }
             }
  case Direita:
             if colunaAtualDaPacgo == labirinto.largura-1{
                 if valorDaPosicaoAtualDaPacgo == ' '{
                   pacgo.posicao.coluna = 0
                 }
             }else{
               var posicaoDireitaDaPacgo = labirinto.mapa[pacgo.posicao.linha][pacgo.posicao.coluna + 1]
               if posicaoDireitaDaPacgo != '#'{
                 pacgo.posicao.coluna = pacgo.posicao.coluna + 1
               }
             }
  case Esquerda:
    if colunaAtualDaPacgo == 0{
      if valorDaPosicaoAtualDaPacgo == ' '{
        pacgo.posicao.coluna = labirinto.largura - 1
      }
    }else{
      var posicaoEsquerdaDaPacgo = labirinto.mapa[pacgo.posicao.linha][pacgo.posicao.coluna - 1]
      if posicaoEsquerdaDaPacgo != '#'{
        pacgo.posicao.coluna = pacgo.posicao.coluna - 1
      }
    }
  }
}

func moverFantasmas() {
  for i := 0; i < 2; i++{ // mudar para quantidade_de_fantasmas
    lista_de_fantasmas[i].posicao.coluna += 1
  }
}

func dorme() {
  time.Sleep(time.Millisecond * 100)
}

func main() {
  inicializa()
  defer finaliza()

  labirinto, pacgo, lista_de_fantasmas, _ = construirLabirinto("")
  labirinto.imprime()
  quantidade_de_fantasmas = len(lista_de_fantasmas)

  pacgo.figura = "\xF0\x9F\x98\x83"

  for _, fantasma := range lista_de_fantasmas {
    fantasma.figura = "\xF0\x9F\x91\xBB"
  }


  for  {
    atualizarLabirinto()

    moverFantasmas()

    m := entradaDoUsuario()
    if m == Sai { break }
    moverPacGo(m)

    if detectarColisao() {
      terminarJogo()
    }

    dorme()
  }
}
