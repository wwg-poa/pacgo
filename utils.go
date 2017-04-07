package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"
)

func inicializa() {
  // inicializaTeclado()

  rawMode := exec.Command("/bin/stty", "cbreak", "-echo")
  rawMode.Stdin = os.Stdin
  _ = rawMode.Run()
  rawMode.Wait()

}

func finaliza() {
  // restore the echoing state when exiting
  //exec.Command("stty", "-F", "/dev/tty", "echo").Run()

  rawMode := exec.Command("/bin/stty", "-cbreak", "echo")
  rawMode.Stdin = os.Stdin
  _ = rawMode.Run()
  rawMode.Wait()
}

var teclas chan byte

func inicializaTeclado() {
  // disable input buffering
  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
  // do not display entered characters on the screen
  exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

  teclas = make(chan byte)
  go rotinaTeclado(teclas)
}

func leTeclado() Movimento {
  select {
    case b := <-teclas:
      if b == 27 {
        fmt.Println("1) Leu 27")
        time.Sleep(100 * time.Millisecond)

        select {
          case b = <-teclas:
            fmt.Println("2) Leu ", b)
            b = <- teclas
            switch b {
              case 65:
                fmt.Println("Cima")
                return Cima
              case 66:
                fmt.Println("Baixo")
                return Baixo
              case 67:
                fmt.Println("Direita")
                return Direita
              case 68:
                fmt.Println("Esquerda")
                return Esquerda
              default:
                fmt.Println("3) ?: ", b)
            }
          default:
            fmt.Println("Sai")
            return Sai
        }
      } else if b == 'q' || b == 'Q' {
        return Sai
      } else {
        fmt.Println("?: ", b)
      }
    default:
  }

  return Nenhum
}

func rotinaTeclado(teclas chan byte) {
  var b []byte = make([]byte, 1)
  for {
    os.Stdin.Read(b)
    teclas <- b[0]
  }
}

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
