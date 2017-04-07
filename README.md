# pacgo

## Estrutura do jogo

- Inicialização de dados
- Criação do labirinto, a partir de um arquivo de texto
- Espera por movimento do usuário (tecla pressionada)
- Movimentação do pacgo
- Movimetação dos fantasmas
- Detecção de colisão
- Fim do jogo


### Inicialização de dados
 Esta parte você pode abstrair. É o momento em que inicilizamos os dados necessários para o jogo, isto é, configuramos o terminal e executamos alguns comandos.

### Criação do labirinto
 O labirinto é criado a partir de um arquivo de texto. Isto significa que vamos fazer o nosso programa ler um arquivo de texto para saber como imprimir cada item. Abaixo, temos uma configuração de exemplo do labirinto.

```
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
```
Sabendo que esses caracteres estarão em um arquivo txt externo, o nosso trabalho é dizer para o nosso programa o que significam cada um dos símbolos #, ., -, F, e G. Além disso, dizemos o que deve ser impresso na tela, em vez do caractere em questão - como você pode reparar, o nosso jogo possui paredes coloridas, pastilhas e emojis.

O trecho de código responsável por esse controle, é mostrado abaixo.

``` go

if char == '#' {
  fmt.Print(labirinto.figura) /* Quando encontrar o caractere '#', imprime a figura do labirinto, ou seja, a parede. */
} else if char == '.'{
  fmt.Print(".") /* Quando encontrar o caractere '.', imprime '.', que corresponde a uma pastilha do jogo. */
} else {
  fmt.Print(" ") /* Caso o caractere encontrado não seja '#' ou '.', imprime um espaço em branco. */
}

```


### Esperando pelo movimento do usuário
O controle do clique das teclas é um pouco mais complexo, mas você só precisa entender a lógica por trás disso. A ideia é diferenciar as teclas pressionadas pelo usuário: se é uma tecla para sair, para fazer o pacgo ir para cima, para baixo ou para os lados. De acordo com a tecla pressionada, determinamos qual é a ação correspondente.

```go
unc entradaDoUsuario(canal chan<- Movimento) {
  array := make([]byte, 10)

  for {
    lido, _ := os.Stdin.Read(array)

    if lido == 1 && array[0] == 0x1b { /* Se a tecla pressionada é o ESC, sai do jogo. */
      canal <- Sai;
    } else if lido == 3 {
      if array[0] == 0x1b && array[1] == '[' {
        switch array[2] {
        case 'A': canal <- Cima /* tecla = seta para cima */
        case 'B': canal <- Baixo /* tecla = seta para baixo */
        case 'C': canal <- Direita /* tecla = seta para direita */
        case 'D': canal <- Esquerda /* tecla = seta para esquerda */
        }
      }
    }
  }
}

```


### Movimentando o pacgo

Para movimentar todos os personagens do jogo, utilizamos uma função chamada *move()*. Essa função é responsável por verificar em qual seta o usuário clicou e o que deve ser feito. Por exemplo, se usuário clica na setinha que aponta para cima, devemos movimentar o pacgo para a linha de cima - caso não tenha uma parede ou algum outro objeto atrapalhando. Isso significa que se o pacgo está na linha 1, deve ir para a linha 0; se está na linha 10, deve ir para a linha 9, e assim por diante.

O comportamento é análogo para qualquer direção que o usuário selecionar: para movimentar o pacgo para baixo, subtraimos 1 unidade da linha atual do pacgo; para direita, somamos 1; e para esquerda, subtraímos 1. O código abaixo é extenso, mas implementa este comportamento.

Um detalhe importante é que se o pacgo estiver na linha 0 (linha na extremidade de cima) e o usuário mandar ele ir para cima, o pacgo aparece na última linha do labirinto, como se fosse um portal! Por isso, temos testes para verificar em qual linha o pacgo está, antes de movimentar o pacgo.

```go

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

```

### Movimentando os fantasmas

 Os fantasmas se movimentam de forma semelhante ao pacgo. A diferença é que eles devem se mover sozinhos, independente de uma tecla que o usuário clicar. Para isso, utilizamos a mesma função *move()*.
 Iteramos sobre a lista de fantasmas, para obter a posição de cada um deles e, assim, movimentar para onde quisermos - isto , movimentá-los de acordo com um algoritmo especial que criamos!


```go

func moverFantasmas() {

  for {
    for i := 0; i < len(lista_de_fantasmas); i++{
        var valorDaPosicaoAtualDoFantasma = labirinto.mapa[lista_de_fantasmas[i].posicao.linha][lista_de_fantasmas[i].posicao.coluna]
        var linhaAtualDoFantasma = lista_de_fantasmas[i].posicao.linha
        var colunaAtualDoFantasma = lista_de_fantasmas[i].posicao.coluna
        move(lista_de_fantasmas[i], valorDaPosicaoAtualDoFantasma, linhaAtualDoFantasma, colunaAtualDoFantasma)
    }
    dorme(200)
  }
}

```

Após movimentar todos os fantasmas do jogo, pedimos para o programa esperar por 200ms, chamando a função *dorme()*.


### Detectando uma colisão
Para detectar se o pacgo encostou em um fantasma, precisamos verificar se a posição do pacgo é a mesma do fantasma. Para isso, o código abaixo é suficiente.

```go

func detectarColisao() bool {
  for _, fantasma := range lista_de_fantasmas {
    if fantasma.posicao == pacgo.posicao {
      return true /* Se algum fantasma estiver na mesma posição que o pacgo, houve uma colisão. */
    }
  }
  return false /* Senão, não houve colisão. */
}


```


### Fim do jogo
O jogo se passa num loop que espera por uma tecla do usuário. Se a tecla for tecla de movimento, movimentamos o pacgo; se for a tecla para sair, saímos do jogo. Caso o pacgo morra (colida com um fantasma), o jogo termina.

```go
for  {
    atualizarLabirinto() /* Imprime na tela as novas posições do pacgo e dos fantasmas. */

    select {
    case tecla = <-canal:
        moverPacGo(tecla) /* Caso o usuário clique em uma tecla de movimento, movimenta o pacgo na direção desejada. */
    default:
    }
    if tecla == Sai { break } /* Caso o usuário clique na tecla para sair, sai do jogo. */

    if detectarColisao() {
      terminarJogo() /* Se houve uma colisão do pacgo com um fantasma, termina o jogo. */
      break;
    }
```

A função que termina o jogo está abaixo.
```go
func terminarJogo() {
  moveCursor( Posicao{labirinto.altura + 2, 0} )
  fmt.Println("Fim de jogo! Os fantasmas venceram... \xF0\x9F\x98\xAD")
}
```
