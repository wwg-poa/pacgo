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

/*
var labirinto *Labirinto
var pacgo     PacGo
var lista_de_fantasmas []Fantasma
*/

func (pacgo *PacGo) imprime() {
  fmt.Println("PacGo")
  fmt.Println(pacgo.posicao.linha)
  fmt.Println(pacgo.posicao.coluna)
}

func imprimeFantasma(fantasmas []*Fantasma) {
  for indice, fantasma := range fantasmas{
    fmt.Println("Fantasma : ",indice)
    fmt.Println(fantasma.posicao.linha)
    fmt.Println(fantasma.posicao.coluna)
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
            fantasma := &Fantasma{ posicao: Posicao{len(mapa), indice}, figura: 'F'}
            fantasmas = append(fantasmas, fantasma)
          }
          //fmt.Println(caracter)
          case 'G': pacgo = &PacGo{ posicao: Posicao{len(mapa), indice}, figura: 'G'}
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
  fmt.Printf("%c", pacgo.figura)

  // TODO: imprime fantasmas
  for _, fantasma := range lista_de_fantasmas {
    moveCursor(fantasma.posicao)
    fmt.Printf("%c", fantasma.figura)
  }

  // Move o cursor para fora do labirinto
  moveCursor(Posicao{labirinto.altura + 2, 1})
}

// func criarFantasmas() {
//   /* Valor inicial utilizado para posicionar um fantasma ao lado do outro */
//   var contador = 10 // hard coded
//   lista_de_fantasmas = make([]Fantasma, 2) // mudar para quantidade_de_fantasmas

//   for i := 0; i < 2; i++ { // mudar para quantidade_de_fantasmas
//       fantasma := new(Fantasma)
//       fantasma.posicao.linha = 5 // hard coded
//       fantasma.posicao.coluna = contador

//       lista_de_fantasmas = append(lista_de_fantasmas, *fantasma)
//       contador += 1
//       fmt.Printf("Fantasma %d: (%d, %d)\n", i, fantasma.posicao.linha, fantasma.posicao.coluna)
//   }
// }

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
  return leTeclado()
}

func moverPacGo() bool {

  var sinal = entradaDoUsuario()

  var valorDaPosicaoAtualDaPacgo = labirinto.mapa[pacgo.posicao.linha][pacgo.posicao.coluna]
  var linhaAtualDaPacgo = pacgo.posicao.linha
  var colunaAtualDaPacgo = pacgo.posicao.coluna

  switch sinal {
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
  time.Sleep(time.Millisecond * 500)
}

func main() {
  defer mostraCursor()
  escondeCursor()

  inicializa()
  defer finaliza()

  labirinto, pacgo, lista_de_fantasmas, _ = construirLabirinto("")
  labirinto.imprime()
  quantidade_de_fantasmas = len(lista_de_fantasmas)

  // pacgo.imprime()
  // imprimeFantasma(lista_de_fantasmas)

  //criarFantasmas()

  // TODO: Loop do jogo
  for  {
    atualizarLabirinto()

    moverFantasmas()

    if (moverPacGo()) {
      break
    }

    if detectarColisao() {
      terminarJogo()
    }

    dorme()
  }
}
