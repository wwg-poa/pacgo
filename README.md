# pacgo

## Estrutura do jogo

- Inicialização de dados
- Criação do labirinto, a partir de um arquivo de texto
- Espera por movimento do usuário (aguarda por uma tecla pressionada)
- Movimentação do pacgo
- Movimetação dos fantasmas
- Detecção de colisão
- Fim do jogo


### Inicialização de dados
 Esta parte você pode abstrair. É o momento em que inicilizamos os dados necessários para o jogo, isto é, configuramos o terminal e executamos alguns comandos.

### Criação do labirinto
 O labirinto é criado a partir de um arquivo de texto. Isto significa que vamos fazer o nosso programa ler um arquivo de texto para saber como imprimir cada item. Abaixo, temos uma configuração de exemplo do labirinto.

 ############################
 #............##............#
 #.####.#####.##.#####.####.#
 #P####.#####.##.#####.####P#
 #..........................#
 #.####.##.########.##.####.#
 #......##....##....##......#
 ######.##### ## #####.######
      #.##          ##.#
      #.## ###--### ##.#
 ######.## # FFFF # ##.######
       .   # FFFF #   .      
 ######.## # FFFF # ##.######
      #.## ######## ##.#
      #.##    G     ##.#     
 ######.## ######## ##.######
 #............##............#
 #.####.#####.##.#####.####.#
 #P..##................##..P#
 ###.##.##.########.##.##.###
 #......##....##....##......#
 #.##########.##.##########.#
 #..........................#
 ############################

Sabendo que esses caracteres estarão em um arquivo txt externo, o nosso trabalho é dizer para o nosso programa o que significam cada um dos símbolos #, ., -, F, e G. Além disso, dizemos o que deve ser impresso na tela, em vez do caractere em questão - como você pode reparar, o nosso jogo possui paredes coloridas, pastilhas e emojis.

O trecho de código responsável por esse controle, é mostrado abaixo.

``` go

if char == '#' {
  fmt.Print(labirinto.figura)
} else if char == '.'{
  fmt.Print(".")
} else {
  fmt.Print(" ")
}

```
