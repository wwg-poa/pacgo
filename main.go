package main

import (
  "fmt"
  "time"
  "os"
  "bufio"
  "log"
  "regexp"
  "errors"
  "math/rand"
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

type Movimento int

const (
        Cima = iota
        Baixo
        Esquerda
        Direita
        Nenhum
        Sai
)

var labirinto *Labirinto
var pacgo     *PacGo
var lista_de_fantasmas []*Fantasma
var mapaSinais map[int]string

func construirLabirinto(nomeArquivo string) (*Labirinto, *PacGo, []*Fantasma, error) {

  var ErrMapNotFound = errors.New("Não conseguiu ler o arquivo do mapa")

  var arquivo string
  if nomeArquivo == "" {
    arquivo = "./data/mapa.txt"
  } else {
    arquivo = nomeArquivo
  }

  if file, err := os.Open(arquivo); err == nil {

    // fecha depois de ler o arquivo
    defer file.Close()

    // inicializa o mapa vazio
    var pacgo *PacGo
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
            fantasma := &Fantasma{ posicao: Posicao{len(mapa), indice}, figura: "\xF0\x9F\x91\xBB"}
            fantasmas = append(fantasmas, fantasma)
          }
          //fmt.Println(caracter)
          case 'G': pacgo = &PacGo{ posicao: Posicao{len(mapa), indice}, figura: "\xF0\x9F\x98\x83"}
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

func atualizarLabirinto() {
  limpaTela()

  for _, linha := range labirinto.mapa {
      fmt.Println(linha)
  }

  // Imprime PacGo
  moveCursor(pacgo.posicao)
  fmt.Printf("%s", pacgo.figura)

  // Imprime fantasmas
  for _, fantasma := range lista_de_fantasmas {
    moveCursor(fantasma.posicao)
    fmt.Printf("%s", fantasma.figura)
  }

  // Move o cursor para fora do labirinto
  moveCursor(Posicao{labirinto.altura + 2, 1})
}

func detectarColisao() bool {
  for _, fantasma := range lista_de_fantasmas {
    if fantasma.posicao == pacgo.posicao {
      return true
    }
  }
  return false
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

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func move(fantasma *Fantasma, valorDaPosicaoAtualDoFantasma byte, linhaAtualDoFantasma int, colunaAtualDoFantasma int){

  var direcao = random(0, 4)
  var sinal = mapaSinais[direcao]
  //fmt.Println(sinal)
  switch sinal {
  case "Cima":
              if linhaAtualDoFantasma == 0{
                if valorDaPosicaoAtualDoFantasma == ' '{
                   fantasma.posicao.linha = labirinto.altura - 1
                 }
             }else{
               var posicaoAcimaDoFantasma = labirinto.mapa[fantasma.posicao.linha - 1][fantasma.posicao.coluna]
               if posicaoAcimaDoFantasma != '#'{
                 fantasma.posicao.linha = fantasma.posicao.linha - 1
               }
             }
  case "Baixo":
              if linhaAtualDoFantasma == labirinto.altura - 1{
                 if valorDaPosicaoAtualDoFantasma == ' '{
                   fantasma.posicao.linha = 0
                 }
              }else{
                var posicaoAbaixoDoFantasma = labirinto.mapa[fantasma.posicao.linha + 1][fantasma.posicao.coluna]
                if posicaoAbaixoDoFantasma != '#'{
                  fantasma.posicao.linha = fantasma.posicao.linha + 1
                }
              }
  case "Direita":
                if colunaAtualDoFantasma == labirinto.largura-1{
                  if valorDaPosicaoAtualDoFantasma == ' '{
                    fantasma.posicao.coluna = 0
                  }
                }else{
                  var posicaoDireitaDofantasma = labirinto.mapa[fantasma.posicao.linha][fantasma.posicao.coluna + 1]
                  if posicaoDireitaDofantasma != '#'{
                    fantasma.posicao.coluna = fantasma.posicao.coluna + 1
                  }
                }
  case "Esquerda":
                 if colunaAtualDoFantasma == 0{
                   if valorDaPosicaoAtualDoFantasma == ' '{
                     fantasma.posicao.coluna = labirinto.largura - 1
                   }
                 }else{
                   var posicaoEsquerdaDoFantasma = labirinto.mapa[fantasma.posicao.linha][fantasma.posicao.coluna - 1]
                   if posicaoEsquerdaDoFantasma != '#'{
                     fantasma.posicao.coluna = fantasma.posicao.coluna - 1
                   }
                 }
  }
}

func moverFantasmas() {

  for {
    for i := 0; i < len(lista_de_fantasmas); i++{
        var valorDaPosicaoAtualDoFantasma = labirinto.mapa[lista_de_fantasmas[i].posicao.linha][lista_de_fantasmas[i].posicao.coluna]
        var linhaAtualDoFantasma = lista_de_fantasmas[i].posicao.linha
        var colunaAtualDoFantasma = lista_de_fantasmas[i].posicao.coluna
        //fmt.Println(valorDaPosicaoAtualDoFantasma, linhaAtualDoFantasma, colunaAtualDoFantasma)
        move(lista_de_fantasmas[i], valorDaPosicaoAtualDoFantasma, linhaAtualDoFantasma, colunaAtualDoFantasma)
    }
    dorme(200)
  }
}

func dorme(mili time.Duration) {
  time.Sleep(time.Millisecond * mili)
}

func entradaDoUsuario(canal chan<- Movimento) {
  array := make([]byte, 10)

  for {
    lido, _ := os.Stdin.Read(array)

    if lido == 1 && array[0] == 0x1b {
      canal <- Sai;
    } else if lido == 3 {
      if array[0] == 0x1b && array[1] == '[' {
        switch array[2] {
        case 'A': canal <- Cima
        case 'B': canal <- Baixo
        case 'C': canal <- Direita
        case 'D': canal <- Esquerda
        }
      }
    }
  }
}

func terminarJogo() {
  // pacgo morreu :(
  moveCursor( Posicao{labirinto.altura + 2, 0} )
  fmt.Println("Fim de jogo! Os fantasmas venceram... \xF0\x9F\x98\xAD")
}

func main() {
  inicializa()
  defer finaliza()

  mapaSinais = make(map[int]string)
  mapaSinais[0] = "Cima"
  mapaSinais[1] = "Baixo"
  mapaSinais[2] = "Direita"
  mapaSinais[3] = "Esquerda"

  args    := os.Args[1:]
  var arquivo string
  if len(args) >= 1 {
    arquivo = args[0]
  } else {
    arquivo = ""
  }

  labirinto, pacgo, lista_de_fantasmas, _ = construirLabirinto(arquivo)

  canal := make(chan Movimento, 10)

  // Processos assincronos
  go entradaDoUsuario(canal)
  go moverFantasmas()

  var tecla Movimento
  for  {
    atualizarLabirinto()

    // canal não-bloqueador
    select {
    case tecla = <-canal:
        moverPacGo(tecla)
    default:
    }
    if tecla == Sai { break }

    if detectarColisao() {
      terminarJogo()
      break;
    }

    dorme(100)
  }
}
