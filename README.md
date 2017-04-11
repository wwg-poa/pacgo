# Tutorial Women Who Go POA

Seja bem vinda ao tutorial do Women Who Go Porto Alegre!

O objetivo deste tutorial é mostrar que programação não é uma coisa de outro mundo e pode ser muito divertido. Nosso objetivo é construir um jogo do zero, utilizando a linguagem de programação Go.

As linguagens de programação são a forma como nós nos comunicamos com os computadores e dizemos para eles o que queremos que eles façam. Além do Go existem inúmeras outras linguagens, cada uma com as suas particularidades.

Assim como uma lingua humana, você não precisa saber tudo de uma linguagem de programação para poder usar ela. Com o tempo você vai aprender recursos mais avançados e formas diferentes de escrever a mesma coisa, algumas melhores, outras piores, ou ainda, apenas diferentes.

Neste tutorial faremos o possível para explicar os principais elementos da linguagem Go, porém não se sinta mal se não entender tudo neste primeiro contato. O mais importante é entender a idéia geral. E é claro, se você estiver participando de um evento presencial, não exite em pedir ajuda para as nossas *coaches*! :)

## Passo 01: Preparar o ambiente

No primeiro passo nós vamos preparar o ambiente de desenvolvimento e criar um primeiro programa executável para testar se tudo está funcionando corretamente.

Você vai precisar de:
- Um computador com acesso a internet

Você vai terminar esta etapa com:
- Uma instalação do Go funcionando
- Um editor de texto simples com o qual você se sinta confortável
- Um programa em Go que escreve "Hello Go!" na tela

Nota: para programação nós não utilizamos editores que formatam texto (por exemplo, Microsoft Word), nós usamos os editores de texto simples (em inglês plain-text).

Sugestões de editores de texto:
- Mac OS X: atom
- Linux: gedit
- Windows: Notepad++

### Instalação do Go

#### Instalação MAC OS X

_TODO_

#### Instalação Linux

_TODO_

#### Instalação Windows

_TODO_

### Meu primeiro programa

Existe uma tradição na área da computação em que toda vez que você vai começar a aprender uma linguagem nova você comece escrevendo um programa chamado "Hello world".

Nós costumamos fazer isso porque geralmente é bem simples de fazer e nos ajuda a testar se a instalação da linguagem está ok.

Então vamos fazer o mesmo para a linguagem Go.

**_Coach_: explicar resumidamente o que é terminal e pasta, se necessário.**

Primeiro, crie uma pasta no seu computador onde você vai guardar os códigos que escrever. Você pode dar qualquer nome para ela, mas de preferência sem acentos ou espaços. Por exemplo: `tutorial`.

Digite no terminal:

```
mkdir tutorial
cd tutorial
```

Abra um editor de texto simples e copie e cole o trecho a seguir:

```
package main

import "fmt"

func main() {
  fmt.Println("Olá Go!")
}
```

Salve o arquivo como `main.go` na pasta que você criou para os seus códigos.

Vamos **compilar** o programa. Este termo pode parecer estranho, mas na verdade isso só quer dizer que vamos traduzir o texto que escrevemos na linguagem de programação para uma linguagem que o computador entende (chamada linguagem de máquina).

Para fazer isso, digite no terminal:

```
go build
```

Com este comando o Go vai "construir" o programa a partir do texto que você escreveu no arquivo `main.go`. Ele vai criar um novo arquivo na mesma pasta que é o chamado "executável".

Vamos rodar este programa. Se você estiver no MAC OS X ou Linux, o comando para executar é assim:

```
./tutorial
```

No Windows, o comando para executar é assim:

```
tutorial
```

Você deve ter visto a mensagem "Olá Go" sendo impressa na tela. Estamos prontas para começar!

## Passo 02: Estrutura de um jogo

Neste tutorial nós vamos construir um jogo chamado PacGo. O nome é uma brincadeira com o clássico PacMan.

Para quem não conhece, o seu objetivo é controlar o PacGo por um labirinto que está repleto de fantasmas. Nos corredores do labirinto existem pastilhas que contam pontos quando o PacGo come elas.

Em alguns pontos estratégicos, existem cogumelos de força que tornam o PacGo temporariamente invencivel aos fantasmas (e pode comer eles também, valendo pontos). O objetivo  do jogo é comer todas as pastilhas do labirinto antes de perder todas as vidas.

O primeiro passo no desenvolvimento de um jogo é o chamado _game design_, que é onde estabelecemos as regras e objetivos do jogo.

Como estamos emprestando a idéia do PacGo de um jogo clássico, vamos pular esta etapa e partir direto para a codificação.

**_Coaches_: explicar brevemente o jogo Pac Man observando os aspectos de _game design_**

Crie uma pasta chamada `pacgo` para separar o código do jogo dos outros arquivos. Lembra como fazer? Se não, volte na seção anterior.

Agora vamos criar o arquivo `main.go` onde vai ficar a parte principal do nosso programa. Abra o editor de textos e copie e cole o seguinte código:

```
package main

func main() {

  // Inicializar labirinto

  // Loop do jogo
  for {
    // Desenha tela

    // Processa entrada do jogador

    // Processa movimento dos fantasmas

    // Processa colisões
  }
}
```

Salve o arquivo.

**_Coach_: explicar o que são comentários, a função main e o que é um loop.**

Note que o nosso programa não faz nada por enquanto (exceto entrar em _loop_ infinito). Isto é porque ainda não inserimos nenhum código no programa, este é apenas um esqueleto com comentários onde vamos colocar os principais componentes do jogo.

Para fazer um jogo nós precisamos nos preocupar com os seguintes detalhes:

- Carregar os dados do jogo (no caso, o mapa e a posição de cada elemento no mapa)
- Criar um loop principal (pois nós queremos que o jogo continue sempre funcionado até decidirmos que ele deve parar)
- Dentro do loop:
  - Desenhar o labirinto na tela
  - Processar o movimento do jogador (também chamado de entrada do usuário)
  - Processar o movimento dos fantasmas
  - Processar colisões, o que quer dizer, verificar se o jogador bateu em algum fantasma

**_Coach_: explicar o papel de cada uma dessas etapas para a construção do jogo.**

Todos estes passos estão anotados no código do programa por meio dos comentários.

Digite no terminal o comando `go build` para criar o programa `pacgo`. Você pode executar o programa que acabou de criar com o comando `./pacgo` (em Linux ou MacOS), ou com o comando `pacgo` (em Windows).

Ao executar o `pacgo` você vai reparar que o programa parece ter **travado** o terminal, porém este é o *loop* infinito do jogo. Para sair do programa, você pode pressionar as teclas *control* (`Ctrl`) e a letra `C` simultaneamente (escrevemos `Ctrl+C` para abreviar). Lembre-se deste atalho, vamos usar ele muitas vezes ao longo do tutorial.

## Passo 03: Construindo um labirinto

A nossa primeira tarefa de codificação vai ser desenhar um labirinto na tela.

Primeiro, nós vamos criar uma representação do labirinto no programa. Para isso vamos utilizar uma `struct`. No arquivo `main.go` adicione o seguinte código entre a linha 1 e a linha 3:

```
type Labirinto struct {
  largura       int
  altura        int
  mapa          []string
}

var labirinto Labirinto
```

A `struct` chamada `Labirinto` serve para guardar as principais características do nosso labirinto: o tamanho dele (representado pela `largura` e `altura`) e o seu desenho, representado pelo `mapa`.

Vamos criar as funções para construir o labirinto e desenhá-lo na tela. Coloque o código abaixo após a linha `var labirinto Labirinto`:

```
func inicializarLabirinto() {
  labirinto = Labirinto{
    largura: 20,
    altura : 10,
    mapa   : []string{
      "####################",
      "#                 F#",
      "#                  #",
      "#                  #",
      "#                  #",
      "#                  #",
      "#                  #",
      "#                  #",
      "#G                 #",
      "####################",
    },
  }
}

func desenhaTela() {
  for _, linha := range labirinto.mapa {
    fmt.Println(linha)
  }
}
```

No mapa, o caractere `#` representa as nossas paredes. A letra `G`representa a posição inicial do nosso personagem (o PacGo) e o `F`representa a posição inicial de um fantasma.

Agora altere a função `main` com a chamada para as duas funções criadas acima colocando-as logo abaixo dos respectivos comentários. Além disso coloque a palavra `break` abaixo do comentário `// Processa entrada do jogador`. O seu código vai ficar assim:

```
func main() {

  // Inicializar labirinto
  inicializarLabirinto()

  // Loop do jogo
  for {
    // Desenha tela
    desenhaTela()

    // Processa entrada do jogador
    break

    // Processa movimento dos fantasmas

    // Processa colisões
  }
}
```

Execute agora o seu código:

```
go run main.go
```

Note que ele imprimiu o labirinto e saiu do programa. Isso é porque colocamos a palavra `break` para quebrar o _loop_ infinito. Experimente tirar esta palavra e ver o que acontece. (Lembre-se que neste caso a combinação de teclas para parar o programa é `Ctrl+C`)

O que você deve ter observado é que sem a palavra `break` dentro do _loop_ (iniciado pela palavra-chave `for`) o programa imprime infinitas vezes o mesmo mapa e a tela fica "rolando" indefinidamente.

Vamos corrigir este comportamento adicionando uma função para limpar a tela antes de imprimir o mapa. Copie e cole o código abaixo antes da função `desenhaTela()`:

```
type Posicao struct {
  linha  int,
  coluna int
}

func moveCursor(p Posicao) {
  fmt.Printf("\x1b[%d;%df", p.linha, p.coluna)
}

func limpaTela() {
  fmt.Printf("\x1b[2J")
  moveCursor( Posicao{0, 0} )
}
```
O código acima define duas funções auxiliares: `moveCursor()` e `limpaTela()`.

Pense no cursor como a "caneta" que escreve na tela. A função moveCursor diz para o computador onde é a próxima posição do terminal onde ele deve escrever.

A função limpaTela apaga todo o conteúdo do terminal e reposiciona o cursor na posição (0, 0), que é o canto superior esquerdo da tela.

**_Coaches_: explicar as coordenadas da tela.**

Não se preocupe com o código dentro das aspas na chamada de função `fmt.Printf()`. Estes são códigos de controle que têm funções especiais. Vale lembrar que pouca gente decora estes códigos - existem tabelas prontas na internet com a lista dos códigos e suas funções.

Agora altere a função `desenhaTela()` para incluir uma chamada para `limpaTela()` antes de imprimir o mapa:

```
func desenhaTela() {
  limpaTela()
  for _, linha := range labirinto.mapa {
    fmt.Println(linha)
  }
}
```
Remova a linha `break` do _loop_ principal e execute novamente o programa. Pode parecer que voltamos ao começo da lição, mas na verdade estamos prontas para fazer animações. A tela parece parada, mas está sendo atualizada várias vezes por segundo (porém sempre com a mesma imagem).

Lembre-se de sair do programa com `Ctrl+C`.

## Passo 04: Mover o PacGo

Agora que nós temos a estrutura de animação pronta, podemos começar a pensar em mover o nosso PacGo (atualmente representado pelo `G` no mapa).

## Passo 05: Mover os fantasmas

## Passo 06: Melhorar o gráfico

## Passo 08: Adicionar pastilhas e pontos

## Passo 09: Adicionar fim de jogo

## Passo 10: Verificar colisões

## Passo 11: Adicionar vidas

## Passo 12: Adicionar cogumelos de força
