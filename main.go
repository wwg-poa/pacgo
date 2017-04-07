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
  posicao    Posicao
  figura     string // emoji
  pilula     bool
  vidas      int
  pontos     int
  invencivel bool
  figuras []string
  indiceFig int
  figuraBravo string
  contadorFig Contador
}

type Contador struct {
  max       int
  contador  int
}

type Fantasma struct {
  posicao Posicao
  figura  string // emoji
}

type Labirinto struct {
  largura       int
  altura        int
  mapa          []string
  figMuro       string
  figMuroSuper  string
  figSP         string
  quantiaPastilhas int
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

var labirinto  *Labirinto
var pacgo      *PacGo
var fantasmas  []*Fantasma
var mapaSinais map[int]string

func criarFantasma(posicao Posicao, figura string) {
  fantasma := &Fantasma{ posicao:posicao, figura: figura}
  fantasmas = append(fantasmas, fantasma)
}

func criarPacGo(posicao Posicao, figura string, pilula bool, vidas int) {
  pacgo = &PacGo{ posicao:posicao, figura: "\xF0\x9F\x98\x83", pilula: false, vidas:3 ,
    figuras: []string {"\xF0\x9F\x98\x83", "\xF0\x9F\x98\x8C"}, indiceFig : 0 , contadorFig: Contador{3, 0},
    figuraBravo: "\xF0\x9F\x98\xA1"}
}

func construirLabirinto(nomeArquivo string) error {

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
    mapa := []string{}

    r, _ := regexp.Compile("[^ #.P]")

    // cria um leitor para ler linha a linha o arquivo
    scanner := bufio.NewScanner(file)
    quantiaPastilhas := 0
    for scanner.Scan() {
      linha := scanner.Text()

      for indice , caracter := range linha {
        switch caracter {
          case 'F': criarFantasma( Posicao{len(mapa), indice}, "\xF0\x9F\x91\xBB" )
          case 'G': criarPacGo( Posicao{len(mapa), indice}, "\xF0\x9F\x98\x83", false, 3 )
          case '.': quantiaPastilhas += 1
        }
      }

      linha = r.ReplaceAllString(linha, " ")
      mapa = append(mapa, linha)
    }

    // verifica se teve erro o leitor
    if err = scanner.Err(); err != nil {
      log.Fatal(err)
      return ErrMapNotFound
    }

    labirinto = &Labirinto{
      largura: len(mapa[0]),
      altura: len(mapa),
      mapa: mapa,
      figMuro: "\x1b[44m \x1b[0m",
      figMuroSuper: "\x1b[41m \x1b[0m",
      figSP: "\xF0\x9F\x8D\x84",
      quantiaPastilhas: quantiaPastilhas,
    }
    return nil

  } else {
    log.Fatal(err)
    return ErrMapNotFound
  }
}

func atualizarLabirinto() {
  limpaTela()

  // Imprime os pontos
  moveCursor(Posicao{0,0})
  fmt.Printf("%sPontos: %d Vidas: %d%s\n", "\x1b[31;1m", pacgo.pontos, pacgo.vidas, "\x1b[0m")

  posicaoInicial := Posicao{2,0}
  moveCursor(posicaoInicial)

  var muro = labirinto.figMuro
  if pacgo.pilula == true{
    muro = labirinto.figMuroSuper
  }

  for _, linha := range labirinto.mapa {
    for _, char := range linha {
      switch char {
        case '#': fmt.Print(muro)
        case '.': fmt.Print(".")
        case 'P': fmt.Print(labirinto.figSP)
        default:  fmt.Print(" ")
      }
    }
    fmt.Println("")
  }

  // Imprime PacGo
  moveCursor(posicaoInicial.adiciona(&pacgo.posicao))
  if (pacgo.pilula) {
    fmt.Printf("%s", pacgo.figuraBravo)
  } else {
    fmt.Printf("%s", pacgo.figuras[pacgo.indiceFig])
    pacgo.incrementaIndice()
  }



  // Imprime fantasmas
  for _, fantasma := range fantasmas {
    moveCursor(posicaoInicial.adiciona(&fantasma.posicao))
    fmt.Printf("%s", fantasma.figura)
  }

  // Move o cursor para fora do labirinto
  moveCursor(posicaoInicial.adiciona(&Posicao{labirinto.altura + 2, 0}))
}

func detectarColisao() bool {
  for _, fantasma := range fantasmas {
    if fantasma.posicao == pacgo.posicao {
      return true
    }
  }
  return false
}

func moverPacGo(m Movimento) {
  var novaLinha = pacgo.posicao.linha
  var novaColuna = pacgo.posicao.coluna

  switch m {
    case Cima:
      novaLinha--
      if novaLinha < 0 {
        novaLinha = labirinto.altura - 1
      }
    case Baixo:
      novaLinha++
      if novaLinha >= labirinto.altura {
        novaLinha = 0
      }
    case Direita:
      novaColuna++
      if novaColuna >= labirinto.largura {
        novaColuna = 0
      }
    case Esquerda:
      novaColuna--
      if novaColuna < 0 {
        novaColuna = labirinto.largura - 1
      }
  }

  conteudoDoMapa := labirinto.mapa[novaLinha][novaColuna]
  if conteudoDoMapa != '#' {
    pacgo.posicao.linha = novaLinha
    pacgo.posicao.coluna = novaColuna

    if (conteudoDoMapa == '.') || (conteudoDoMapa == 'P') {
      if (conteudoDoMapa == '.') {
        pacgo.pontos += 10
        fmt.Print("\x07")
        labirinto.quantiaPastilhas--
      } else {
        pacgo.pontos += 100
        ativarPilula()
      }

      linha := labirinto.mapa[novaLinha]
      linha = linha[:novaColuna] + " " + linha[novaColuna+1:]
      labirinto.mapa[novaLinha] = linha
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
    for i := 0; i < len(fantasmas); i++{
        var valorDaPosicaoAtualDoFantasma = labirinto.mapa[fantasmas[i].posicao.linha][fantasmas[i].posicao.coluna]
        var linhaAtualDoFantasma = fantasmas[i].posicao.linha
        var colunaAtualDoFantasma = fantasmas[i].posicao.coluna
        //fmt.Println(valorDaPosicaoAtualDoFantasma, linhaAtualDoFantasma, colunaAtualDoFantasma)
        move(fantasmas[i], valorDaPosicaoAtualDoFantasma, linhaAtualDoFantasma, colunaAtualDoFantasma)
    }
    dorme(300)
  }
}

func dorme(milisegundos time.Duration) {
  time.Sleep(time.Millisecond * milisegundos)
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

func ativarPilula() {
  pacgo.pilula = true
  go desativarPilula(10000)
}

func desativarPilula(milisegundos time.Duration) {
  dorme(milisegundos)
  pacgo.pilula = false
}

func terminarJogo() {
  // pacgo morreu :(
  atualizarLabirinto()
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

  construirLabirinto(arquivo)

  canal := make(chan Movimento, 10)

  // Processos assincronos
  go entradaDoUsuario(canal)
  go moverFantasmas()

  var tecla Movimento
  for {
    atualizarLabirinto()
    if labirinto.quantiaPastilhas == 0 {
      moveCursor( Posicao{labirinto.altura + 2, 0} )
      fmt.Println("Fim de jogo! Você venceu! \xF0\x9F\x98\x84")
      break
    }

    // canal não-bloqueador
    select {
      case tecla = <-canal:
        moverPacGo(tecla)
      default:
    }
    if tecla == Sai { break }

    if detectarColisao() {
      if pacgo.pilula {
        _, f := buscaFantasma(pacgo.posicao)
        go criarFantasmaTemporizado(f, 5000)
        matarFantasma(f)
        pacgo.pontos = pacgo.pontos + 500
      } else {
        // pacgo perde vidas
        if !pacgo.invencivel {
          pacgo.vidas--
          if pacgo.vidas == 0 {
            terminarJogo()
            break
          }
          ativarInvencibilidade(3000)
        }
      }
    }

    dorme(100)
  }
}

func ativarInvencibilidade(milisegundos time.Duration) {
  pacgo.invencivel = true
  go func() {
    dorme(milisegundos)
    pacgo.invencivel = false
  }()
}

func buscaFantasma(posicao Posicao) (int, *Fantasma) {
  for i, fantasma := range fantasmas {
    if fantasma.posicao == posicao {
      return i, fantasma
    }
  }
  return -1, nil
}

func criarFantasmaTemporizado(fantasma *Fantasma, milisegundos time.Duration) {
  dorme(milisegundos)
  criarFantasma(fantasma.posicao, fantasma.figura)
}

func matarFantasma(fantasma *Fantasma) {
  pos, _ := buscaFantasma(fantasma.posicao)
  fantasmas = append(fantasmas[:pos], fantasmas[pos+1:]...)
  fmt.Print("\x07")
}
