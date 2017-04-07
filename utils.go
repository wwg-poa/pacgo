package main

import (
  "fmt"
  "os"
  "os/exec"
)

func inicializa() {
  rawMode := exec.Command("/bin/stty", "cbreak", "-echo")
  rawMode.Stdin = os.Stdin
  _ = rawMode.Run()
  rawMode.Wait()
}

func finaliza() {
  rawMode := exec.Command("/bin/stty", "-cbreak", "echo")
  rawMode.Stdin = os.Stdin
  _ = rawMode.Run()
  rawMode.Wait()
}

func (posicao1 *Posicao) adiciona(posicao2 *Posicao) (Posicao){
  return Posicao{posicao1.linha + posicao2.linha, posicao1.coluna + posicao2.coluna}
}

func (pacgo *PacGo) incrementaIndice() {
  if (pacgo.contadorFig.contador == pacgo.contadorFig.max){
    if (pacgo.indiceFig == len(pacgo.figuras)-1) {
      pacgo.indiceFig = 0
    } else {
      pacgo.indiceFig += 1
    }
    pacgo.contadorFig.contador = 0 
  } else {
    pacgo.contadorFig.contador += 1
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
