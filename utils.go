package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"
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
