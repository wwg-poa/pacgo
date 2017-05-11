# Tutorial Women Who Go POA

Seja bem-vinda ao tutorial do Women Who Go Porto Alegre!

O objetivo deste tutorial é mostrar que programação não é uma coisa de outro mundo e pode ser muito divertido. Nosso objetivo é construir um jogo do zero, utilizando a linguagem de programação Go.

As linguagens de programação são a forma como nós nos comunicamos com os computadores e dizemos para eles o que queremos que eles façam. Além do Go existem inúmeras outras linguagens, cada uma com as suas particularidades.

Assim como uma língua humana, você não precisa saber tudo de uma linguagem de programação para usá-la. Com o tempo você vai aprender recursos mais avançados e formas diferentes de escrever a mesma coisa, algumas melhores, outras piores, ou ainda, apenas diferentes.

Neste tutorial faremos o possível para explicar os principais elementos da linguagem Go, porém não se sinta mal se não entender tudo neste primeiro contato. O mais importante é entender a idéia geral e é claro, se você estiver participando de um evento presencial, não hesite em pedir ajuda para as nossas *coaches*! :)

## Passo 01: Preparar o ambiente

No primeiro passo nós vamos preparar o ambiente de desenvolvimento e criar um primeiro programa executável para testar se tudo está funcionando corretamente.

Você vai precisar de:
- Um computador com acesso à internet

Você vai terminar esta etapa com:
- Uma instalação do Go funcionando
- Um editor de texto simples com o qual você se sinta confortável
- Um programa em Go que escreve "Hello Go!" na tela

Nota: para programação nós não utilizamos editores que formatam texto (por exemplo, Microsoft Word), nós usamos os editores de texto simples (em inglês *plain-text*).

Sugestões de editores de texto:
- MacOS: atom
- Linux: gedit
- Windows: Notepad++

### Instalação do Go

#### Instalação MacOS

Temos duas opções de como instalar o Go no MacOS: usando Homebrew ou o instalador de pacotes do MacOS.

##### Instalando Homebrew (Opcional)

O jeito mais simples de instalar o Go é usando o [Homebrew](https://brew.sh/), que é um gerenciador de pacotes para o MacOS. Para instalar o Homebrew basta rodar o seguinte comando no terminal:

```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

O próprio script de intalação do Homebrew explica o que esta fazendo e pausa quando necessário.

##### Instalando Go

Caso você tenha instalado o Homebrew, basta rodar:

```
brew install go
```

Caso você tenha optado pelo instalador de pacotes do MacOS, abra a [página de downloads](https://golang.org/dl/) do Go procure pelo arquivo do instalador para MacOS e clique no link. Os arquivos serão instalados por padrão em `/usr/local/go`.

Em ambos os casos, após a instalação a sua variável de ambiente `PATH` já deve conter o binário do go. Para que o comando possa ser utilizado é possível que você tenha que reiniciar as sessões do terminal que estão abertas.

Para garantir que a instalação foi bem sucedida e que o Go foi instalado corretamente, basta rodar o comando `go version` e observar uma saída parecida com a seguinte:

```
❯ go version
go version go1.8.1 darwin/amd64
```

#### Instalação Linux

##### Instalando Go

Se você usa uma distribuição do Linux baseada no Debian, como o Ubuntu, basta rodar o seguinte comando:

```
sudo apt-get install golang-1.8-go
```

Para outras distribuições de linux o primeiro passo para instalar o Go é baixar o arquivo .tar.gz, para isso abra a [página de downloads](https://golang.org/dl/) do Go procure pelo link correto.

Em seguida será necessário extrair o conteúdo do arquivo baixado em `/usr/local`, para criar uma árvore dos arquivos do Go em `/usr/local/go`.

```
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

Por exemplo, se você baixou o arquivo `go1.8.1.linux-amd64.tar.gz`:
```
sudo tar -C /usr/local -xzf go1.8.1.linux-amd64.tar.gz
```

Finalmente, tanto para a instalação com `apt-get` quanto com o arquivo .tar.gz é necessário adicionar o binário do Go à sua variável de ambiente `PATH`. Para fazer isso adicione ao bash profile `$HOME/.profile` a linha seguinte:

```
export PATH=$PATH:/usr/local/go/bin
```

Para que o comando possa ser utilizado é possível que você tenha que reiniciar as sessões do terminal que estão abertas.

Para garantir que a instalação foi bem sucedida e que o Go foi instalado corretamente, basta rodar o comando `go version` e observar uma saída parecida com a seguinte:

```
ubuntu@svartir-sandar:~$ go version
go version go1.8.1 linux/amd64
```

#### Instalação Windows

##### Instalando Git

O primeiro passo é instalar o Git. O jeito mais fácil é fazer o download do [Git For Windows](https://git-for-windows.github.io/). A vantagem de utilizar o Git For Windows é que também será instalado o Git Bash. No caso desse tutorial, faremos os personagens do jogo usando caracteres Unicode. Infelizmente, o prompt de comando do Windows não consegue exibir os caracteres Unicode, então para esse tutorial usaremos o Git Bash que suporta esses caracteres.

Para instalar o Git For Windows basta clicar no link acima e então em Download. Siga os passos da instalação normalmente até chegar na tela de seleção de componentes. Nessa tela é muito importante selecionar a opção "Use a TrueType font in all console windows". É essa opção que vai fazer com que os caracteres Unicode sejam exibidos corretamente.

![Tela de seleção de componentes do instalador do Git For Windows.](./img/git_bash_setup_unicode.png)

Continue seguindo as instruções de instalação até a tela de configuração do emulador de terminal que será utilizado pelo Git Bash. Nessa tela selecione a opção "Use MinTTY (the default terminal of MSYS2)". Juntamente com a opção selecionada anteriormente, essa opção também garantirá a exibição correta dos carcteres Unicode.

![Tela de seleção do emulador de terminal do instalador do Git For Windows.](./img/git_bash_setup_unicode_2.png)

Após essa tela continue seguindo as instruções do instalador até que o Git For Windows seja instalado. Ao fim da instalação abra o Git Bash.

##### Instalando Go

O jeito mais simples de instalar o Go no Windows é utilizar o instalador MSI. Na [página de downloads](https://golang.org/dl/) do Go procure pelo arquivo do instalador para Windows e clique no link. Os arquivos serão instalados por padrão em `C:\Go`.

Após a instalação a sua variável de ambiente `PATH` já deve conter o binário do go (`C:\Go\bin`). Para que o comando possa ser utilizado é possível que você tenha que reiniciar o Git Bash que está aberto.

Para garantir que a instalação foi bem sucedida e que o Go foi instalado corretamente, basta rodar o comando `go version` e observar uma saída parecida com a seguinte:

```
C:\Users\Camila\Documents\GitHub> go version
go version go1.8.1 windows/amd64
```

### Configurando o Workspace do Go

A linguagem Go necessita que todo o código Go esteja localizado em um único workspace. O workspace é uma pasta que contém três sub-pastas: uma pasta `src`, que contém todos os arquivos fonte em Go. Uma pasta `pkg`, que contém os objetos dos pacotes e uma pasta `bin` que contém os comandos executáveis. Para mais detalhes leia a [documentação](https://golang.org/doc/code.html#Workspaces) da linguagem.

A localização do workspace é definido por uma varável de ambiente chamada `GOPATH`. A localização padrão dessa variável é `%USERPROFILE%\go` (geralmente C:\Users\SeuNome\go) no Windows ou `$HOME/go` no caso do MacOS e Linux. Nesse tutorial usaremos a localização padrão do workspace. Caso você deseje utilizar outro local será necessário alterar o valor do `GOPATH` para que seu programa funcione corretamente.

Para criar o workspace na localização padrão vá até a pasta home e crie a pasta `go`. Dentre as três sub-pastas do workspace, apenas a pasta `src` precisa ser criada:

#### MacOS / Linux / Windows (usando Git Bash)
```
cd ~
mkdir go
cd go
mkdir src
cd src
```

#### Windows (usando linha de comando do Windows)
```
cd %USERPROFILE%
mkdir go
cd go
mkdir src
cd src
```

Pronto! É dentro desta pasta `src` onde você pode começar a construir seu primeiro programa em Go :)

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

Vamos rodar este programa. Se você estiver no MacOS ou Linux, o comando para executar é assim:

```
./tutorial
```

No Windows, o comando para executar é assim:

```
tutorial.exe
```

Você deve ter visto a mensagem "Olá Go!" sendo impressa na tela. Estamos prontas para começar!

## Passo 02: Estrutura de um jogo

Neste tutorial nós vamos construir um jogo chamado PacGo. O nome é uma brincadeira com o clássico PacMan.

Para quem não conhece, o seu objetivo é controlar o PacGo por um labirinto que está repleto de fantasmas. Nos corredores do labirinto existem pastilhas que contam pontos quando o PacGo come elas.

Em alguns pontos estratégicos, existem cogumelos de força que tornam o PacGo temporariamente invencível aos fantasmas (e pode comê-los também, valendo pontos). O objetivo do jogo é comer todas as pastilhas do labirinto antes de perder todas as vidas.

O primeiro passo no desenvolvimento de um jogo é o chamado _game design_, que é onde estabelecemos as regras e objetivos do jogo.

Como estamos emprestando a idéia do PacGo de um jogo clássico, vamos pular esta etapa e partir direto para a codificação.

**_Coach_: explicar brevemente o jogo Pac Man observando os aspectos de _game design_**

Digite no seu terminal o seguinte comando:

```
go get github.com/wwg-poa/tutorial
```

Ele vai baixar automaticamente para você os arquivos iniciais deste projeto.

Vá para a pasta `$GOHOME/src/github.com/wwg-poa/tutorial` e abra o arquivo `main.go` no seu editor de textos. Você deve ver o código abaixo:


```
package main

import "fmt"

func main() {
  // Inicializar terminal
  Inicializa()
  defer Finaliza()

  // Inicializar labirinto

  // Loop principal
  for {
    // Desenha tela

    // Processa entrada do jogador

    // Processa movimento dos fantasmas

    // Processa colisões

    // Dorme

    fmt.Println("Olá Go!")
    break // Temporário: quebra o loop infinito
  }
}
```

Salve o arquivo. (Lembre-se sempre de salvar o arquivo após cada alteração!)

**_Coach_: explicar o que são comentários, a função `main` e o que é um _loop_.**

Note que o nosso programa não faz nada diferente do programa anterior. Porém, nós incluimos alguns comentários com o objetivo de preparar o terreno para as próximas etapas e um _loop_ `for` para ser o nosso _loop_ principal do jogo.

Além disso, logo no começo a função `main` incluímos as chamadas para as funções `Inicializa` e `Finaliza`. O objetivo destas funções é preparar o terminal para que ele entenda corretamente as instruções de impressão e os comandos do teclado, e restaurar ele para o modo anterior quando acabarmos (a palavra chave `defer` diz para o Go executar a função `Finaliza` por último).

O código que faz isto é este aqui:

```
// Inicializar terminal
Inicializa()
defer Finaliza() // executa no final da função
```

Note que não definimos estas funções neste arquivo, elas foram definidas para você no arquivo `utils.go`. O entendimento destas funções não é necessário para este tutorial, mas caso fique curiosa fique a vontade para explorar este arquivo.

Resumidamente, para fazer um jogo nós precisamos nos preocupar com os seguintes detalhes:

- Preparar os recursos do computador (tela, teclado, etc)
- Carregar os dados do jogo (no caso, o mapa e a posição de cada elemento no mapa)
- Criar um _loop_ principal (pois nós queremos que o jogo continue sempre funcionado até decidirmos que ele deve parar)
- Dentro do _loop_:
  - Desenhar o labirinto na tela
  - Processar o movimento do jogador (também chamado de entrada do usuário)
  - Processar o movimento dos fantasmas
  - Processar colisões, o que quer dizer, verificar se o jogador bateu em algum fantasma

**_Coach_: explicar o papel de cada uma dessas etapas para a construção do jogo.**

Todos estes passos estão anotados no código do programa por meio dos comentários.

Digite no terminal o comando `go build` para criar o programa `pacgo`. Você pode executar o programa que acabou de criar com o comando `./pacgo` (em Linux ou MacOS), ou com o comando `pacgo` (em Windows).

Ao executar o `pacgo` você vai reparar que o programa parece ter **travado** o terminal, porém este é o *loop* infinito do jogo. Para sair do programa, você pode pressionar as teclas *control* (`Ctrl`) e a letra `C` simultaneamente (escrevemos `Ctrl+C` para abreviar). Lembre-se deste atalho, vamos usá-lo muitas vezes ao longo do tutorial.

## Passo 03: Construindo um labirinto

A nossa primeira tarefa de codificação vai ser desenhar um labirinto na tela.

**_Coach_: explique em poucas palavras o que é um _import_ e o que são bibliotecas.**

Nós vamos criar uma representação do labirinto no programa. Para isso vamos utilizar uma `struct`. As _structs_ são a nossa forma de dizer que uma coisa possui várias partes, ou "propriedades". No caso, o nosso labirinto possui uma `largura`, uma `altura` e um `mapa`.

No arquivo `main.go` adicione o seguinte código abaixo de `import "fmt"`:

```
type Labirinto struct {
  largura       int
  altura        int
  mapa          []string
}

var labirinto Labirinto
```

**_Coach_: explicar a diferença entre declaração e definição.**

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

**_Coach_: explicar a diferença entre declaração de função e chamada de função.**

No mapa, o caractere `#` representa as nossas paredes. A letra `G` representa a posição inicial do nosso personagem (o PacGo) e o `F` representa a posição inicial de um fantasma.

Agora altere a função `main` para incluir a chamada para a função `inicializarLabirinto` antes do _loop_ principal. Dentro do _loop_ adicione a chamada para `desenhaTela`. Finalmente, remova a linha que imprime "Olá Go!".

O código da função `main`deve ficar assim:

```
func main() {
  // Inicializar terminal
  inicializa()
  defer finaliza() // executa apenas no fim do programa

  // Inicializar labirinto
  inicializarLabirinto()

  // Loop principal
  for {
    // Desenha tela
    desenhaTela()

    // Processa entrada do jogador

    // Processa movimento dos fantasmas

    // Processa colisões

    // Dorme

    break
  }
}
```

Vamos executá-lo, mas não esqueça de compilar o programa primeiro. No terminal:

```
go build
./tutorial
```

Note que ele imprimiu o labirinto e saiu do programa. Isso é porque colocamos a palavra `break` para quebrar o _loop_ infinito.

## Passo 04: Adicionando a entrada do teclado

Por enquanto nosso programa só imprime o labirinto e sai da tela. Nada muito emocionante, certo? Mas antes de começarmos a ver as animações, precisamos preparar um pouco mais o terreno e incluir uma forma do usuário interagir com o programa. Para isso precisamos que o nosso jogo reconheça os comandos do teclado.

Nós estamos particularmente interessadas em 5 teclas: a tecla ESC, que vai ser usada para sair do jogo, e as setas, que vão ser usadas para controlar o PacGo.

Para o computador, cada tecla pressionada no teclado tem um valor númerico especial. Nós lemos qual tecla o usuário pressinou com a função `os.Stdin.Read`. O código abaixo faz a leitura apenas das teclas que nos interessam e dá nomes mais amigáveis para elas através do tipo `Entrada`.

Copie e cole o código abaixo logo após a linha 3 (`import "fmt"`):

```
import "os"

type Entrada int

// Possiveis entradas do usuário
const (
  Nada = iota  
  ParaCima
  ParaBaixo
  ParaEsquerda
  ParaDireita
  SairDoJogo // Tecla ESC
)

func leEntradaDoUsuario() Entrada {
  var m Entrada
  array := make([]byte, 10)

  lido, _ := os.Stdin.Read(array)

  if lido == 1 && array[0] == 0x1b {
    m = SairDoJogo;
  } else if lido == 3 {
    if array[0] == 0x1b && array[1] == '[' {
      switch array[2] {
        case 'A': m = ParaCima
        case 'B': m = ParaBaixo
        case 'C': m = ParaDireita
        case 'D': m = ParaEsquerda
      }
    }
  }
  return m
}
```

Agora que nós sabemos quando o usuário pressionou a tecla `ESC`, podemos nos livrar do comando `break` no _loop_ principal e deixar o usuário decidir quando quer encerrar o programa.

Altere o _loop_ principal para incluir a chamada para `leEntradaDoUsuario` conforme o código abaixo:

```
// Loop principal
for {
  // Desenha tela
  desenhaTela()

  // Processa entrada do jogador
  m := leEntradaDoUsuario()
  if m == SairDoJogo { break }

  // Processa movimento dos fantasmas

  // Processa colisões

  // Dorme
}
```

A linha `if m == SairDoJogo { break }` interrompe o jogo toda vez que você pressionar `ESC`. Experimente:

```
go build
./tutorial
```

Você vai reparar que o jogo fica parado até você pressionar a tecla `ESC` sair. Porém, você também deve ter percebido que ao pressionar qualquer outra tecla o jogo imprime novamente o mapa logo abaixo do anterior.

Isto acontece porque a cada passo do _loop_ o computador fica esperando você pressionar uma tecla, parando a execução na chamada da função `leEntradaDoUsuario`. Quando a tecla chega, ele "desprende" o programa e executa novamente o _loop_, passando por `desenhaTela`.

Primeiro, vamos fazer a tela ser impressa corretamente.

## Passo 05: Corrigindo a animação

**_Coach_: explicar como funciona o sistema de coordenadas da tela.**

Altere a função `desenhaTela()` para incluir uma chamada para `LimpaTela()` antes de imprimir o mapa:

```
func desenhaTela() {
  LimpaTela() // adicione esta linha
  for _, linha := range labirinto.mapa {
    fmt.Println(linha)
  }
}
```

A função limpa tela garante que a tela remove todo o conteúdo do terminal e retorna o cursor para a posição (0, 0) para que o desenho seja feito sempre no mesmo lugar.

Experimente executar o programa novamente. Lembre-se que a tecla para sair é `ESC`.

Você deve reparar que agora o programa parece não responder a nenhuma tecla exceto a `ESC`... isto acontece na verdade porque nós ainda não programamos as outras teclas para fazer nada.

## Passo 06: Mover o PacGo

Agora que nós temos a estrutura de animação pronta, podemos começar a pensar em mover o nosso PacGo (atualmente representado pelo `G` no mapa).

Para facilitar o controle do PacGo ao longo de todo o programa, vamos primeiro criar uma estrutura para representá-lo. Cole o código antes da definição da função `leEntradaDoUsuario`:

```
type PacGo struct {
  posicao    Posicao
  figura     string
}

var pacgo PacGo

func criarPacGo(posicao Posicao, figura string) {
  pacgo = PacGo{
    posicao: posicao,
    figura: "G",
  }
}
```

Agora vamos alterar a função `inicializarLabirinto` para construir o PacGo com a sua posição correta no mapa:

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

  // O código novo começa aqui
  for linha, linhaMapa := range labirinto.mapa {
    for coluna, caractere := range linhaMapa {
      switch( caractere ) {
        case 'G': { criarPacGo( Posicao{linha, coluna}, "G") }
      }
    }
  }
}
```

Com o PacGo criado podemos movimentá-lo. Copie e cole o código abaixo depois da definição da função `desenhaTela`:

```
func moverPacGo(m Movimento) {
  var novaLinha = pacgo.posicao.linha
  var novaColuna = pacgo.posicao.coluna

  switch m {
    case ParaCima:
      novaLinha--
      if novaLinha < 0 {
        novaLinha = labirinto.altura - 1
      }
    case ParaBaixo:
      novaLinha++
      if novaLinha >= labirinto.altura {
        novaLinha = 0
      }
    case ParaDireita:
      novaColuna++
      if novaColuna >= labirinto.largura {
        novaColuna = 0
      }
    case ParaEsquerda:
      novaColuna--
      if novaColuna < 0 {
        novaColuna = labirinto.largura - 1
      }
  }

  conteudoDoMapa := labirinto.mapa[novaLinha][novaColuna]
  if conteudoDoMapa != '#' {
    pacgo.posicao.linha = novaLinha
    pacgo.posicao.coluna = novaColuna
  }
}
```

A função `moverPacGo` recebe um sinal de movimento e tenta atualizar a posição atual do PacGo. Porém, se a nova posição cair numa parede (representada pelo caractere `#`) a função ignora o movimento.

O próximo passo é alterar o programa principal para chamar esta função toda vez que alguém pressionar uma tecla.

Altere o código abaixo do comentário `// Processa entrada do jogador` para o código a seguir:

```
// Processa entrada do jogador
m := leEntradaDoUsuario()

if m == SairDoJogo {
  break
} else {
  moverPacGo(m)
}
```

O último passo vai ser alterar a função `desenhaTela` para atualizar a posição do PacGo a cada passada. Modifique o código desta função para que fique igual a função abaixo:

```
func desenhaTela() {
  LimpaTela()

  // Imprime mapa
  for _, linha := range labirinto.mapa {
    for _, char := range linha {
      switch char {
        case '#': fmt.Print("#")
        default:  fmt.Print(" ")
      }
    }
    fmt.Println("")
  }

  // Imprime PacGo
  MoveCursor(pacgo.posicao)
  fmt.Printf("%s", pacgo.figura)

  // Move cursor para fora do labirinto
  MoveCursor(Posicao{labirinto.altura + 2, 0})
}
```
**_Coach_: comentar o impacto das mudanças na função desenhaTela.**

Compile o programa e execute-o. Você deve reparar que as setas movem o `G` na tela. Estamos fazendo progresso!

Pressione `ESC` para sair.

## Passo 07: Movendo os fantasmas

Agora que o nosso PacGo é capaz de se mexer está na hora de animar os fantasmas. Vamos começar definindo uma estrutura para representar os fantasmas no código. Copie e cole a definição da `struct` Fantasma após a definição da `struct` PacGo:

```
type Fantasma struct {
	posicao Posicao
	figura  string
}
```

Assim como para o PacGo a estrutura acima só define a "receita" para construir o fantasma. Precisamos também criar os fantasmas propriamente ditos. Como podem existir mais de um fantasma, ao invés de declarar um único objeto fantasma vamos declarar um _array_ de fantasmas.

**_Coach_: explicar o que é um _array_.**

Copie e cole o código abaixo da definição do PacGo:

```
var fantasmas []*Fantasma
```

Também precisamos de uma função para criar fantasmas. Copie e cole o código abaixo da função `criarPacGo`:

```
func criarFantasma(posicao Posicao, figura string) {
	fantasma := &Fantasma{posicao: posicao, figura: figura}
	fantasmas = append(fantasmas, fantasma)
}
```

Precisamos alterar a função que inicializa o mapa (`inicializarLabirinto`) para chamar a função `criarFantasma` toda vez que encontrar um caractere `F`. Faça a modificação abaixo:

```
// Processa caracteres especiais
for linha, linhaMapa := range labirinto.mapa {
  for coluna, caractere := range linhaMapa {
    switch( caractere ) {
      case 'G': { criarPacGo(Posicao{linha, coluna}, "G") }
      case 'F': { criarFantasma(Posicao{linha, coluna}, "F") }
    }
  }
}
```

A estrutura para criar fantasmas está completa, mas ainda não temos o código que exibe eles. Para isso precisamos alterar a função `desenhaTela`:

```
// Imprime PacGo
MoveCursor(pacgo.posicao)
fmt.Printf("%s", pacgo.figura)

// Imprime fantasmas
for _, fantasma := range fantasmas {
  MoveCursor(fantasma.posicao)
  fmt.Printf("%s", fantasma.figura)
}

// Move cursor para fora do labirinto
MoveCursor(Posicao{labirinto.altura + 2, 0})
```

O último passo é o código que move os fantasmas.  Crie a função abaixo depois da definição da função `moverPacGo`:

```
func moverFantasmas() {
  for _, fantasma := range fantasmas {
    // gera um número aleatório entre 0 e 4 (ParaDireita = 3)
    var direcao = rand.Intn(ParaDireita+1)

    var novaPosicao = fantasma.posicao

    // Atualiza posição testando os limites do mapa
    switch direcao {
    case ParaCima:
      novaPosicao.linha--
      if novaPosicao.linha < 0 { novaPosicao.linha = labirinto.altura - 1 }
    case ParaBaixo:
      novaPosicao.linha++
      if novaPosicao.linha > labirinto.altura - 1 { novaPosicao.linha = 0 }
    case ParaEsquerda:
      novaPosicao.coluna--
      if novaPosicao.coluna < 0 { novaPosicao.coluna = labirinto.largura - 1 }
    case ParaDireita:
      novaPosicao.coluna++
      if novaPosicao.coluna > labirinto.largura - 1 { novaPosicao.coluna = 0 }
    }

    // Verifica se a posição nova do mapa é válida
    conteudoMapa := labirinto.mapa[novaPosicao.linha][novaPosicao.coluna]
    if conteudoMapa != '#' { fantasma.posicao = novaPosicao }
  }
}
```

Como nós estamos utilizando a função `rand.Intn` do pacote `rand` precisamos adicionar o respectivo `import`:

```
import "math/rand"
```

E adicione a chamada para esta função abaixo do comentário `// Processa movimento dos fantasmas` no _loop_ principal:

```
// Processa movimento dos fantasmas
moverFantasmas()
```

Compile o programa e utilize as setas para mover o PacGo. Você vai reparar que o `F` que representa o nosso fantasma vai se mover toda vez que você pressionar uma tecla.

Experimente adicionar mais um fantasma no mapa para ver o que acontece.

## Passo 08: Corrigindo o movimento

**_Coach_: Explicar porque o movimento dos fantasmas só ocorre após pressionar uma tecla.**

Você deve ter reparado na seção anterior que o nosso jogo "trava" esperando o usuário pressionar uma tecla. Num jogo de verdade é esperado que o movimento dos inimigos seja independente do movimento do jogador. Nós precisamos separar o código que lê as teclas pressionadas pelo usuário do código do _loop_ principal.

**_Coach_: Explicar brevemente os conceitos de _goroutine_ e _canais_.**

Para conseguir este objetivo, vamos utilizar o conceito de `goroutines` e canais (`channels`). A função de uma _goroutine_ é justamente executar um código separado do código principal.

Porém, como ele vai estar separado, é preciso ter uma maneira de comunicar com o código principal (por exemplo, para informar qual tecla foi pressionada). É aí que entram os canais: são formas de comunicação entre dois códigos que estão executando em paralelo.

Altere a função `leEntradaDoUsuario` para ter a seguinte forma:

```
func leEntradaDoUsuario(canal chan<- Entrada) {
	array := make([]byte, 10)

	for {
		lido, _ := os.Stdin.Read(array)

		if lido == 1 && array[0] == 0x1b {
			canal <- SairDoJogo
		} else if lido == 3 {
			if array[0] == 0x1b && array[1] == '[' {
				switch array[2] {
				case 'A': canal <- ParaCima
				case 'B':	canal <- ParaBaixo
				case 'C':	canal <- ParaDireita
				case 'D': canal <- ParaEsquerda
				}
			}
		}
	}
}
```

Repare nas seguintes mudanças:

1) a função agora recebe um parâmetro do tipo `chan<- Entrada` e não possui retorno.
2) ao invés de gravar na variável `m`, a função grava na variável `canal` com o operador `<-` ao invés de `=`
3) todo o código está envolto por um _loop_ infinito, o que quer dizer que uma vez chamada esta função vai executar repetidamente até o término do programa

Agora vamos alterar a função `main` para chamar a função `leEntradaDoUsuario` como uma _goroutine_:

```
func main() {
  // Inicializar terminal
  Inicializa()
  defer Finaliza()

  // Inicializar labirinto
  inicializarLabirinto()

  // Cria rotina para ler entradas
  canal := make(chan Movimento, 10)
	go leEntradaDoUsuario(canal)

  // Loop principal
  for {
    // Desenha tela
    desenhaTela()

    // Processa entrada do jogador
    var tecla Entrada
		select {
		case tecla = <-canal:
		default:
		}

    if tecla == SairDoJogo {
      break
    } else {
      moverPacGo(tecla)
    }

    // Processa movimento dos fantasmas
    moverFantasmas()

    // Processa colisões

    // Dorme
  }
}

```

Se você compilar e executar o programa agora, vai reparar que o movimento dos fantasmas não depende mais do movimento do PacGo. Porém, temos um efeito indesejado que é a tela ficar piscando rapidamente.

Isto acontece porque o jogo não está mais "travando" na entrada do usuário antes de atualizar a tela e está atualizando ela o mais rápido possível.

Tipicamente, para dar a ilusão de movimento, os jogos atualizam a tela do jogador várias vezes por segundo, onde cada uma destas telas apresenta uma imagem (também chamada de quadro ou _frame_) com uma pequena diferença em relação a anterior.

Se atualizarmos a tela poucas vezes por segundo a animação vai parecer travada, mas se atualizarmos rápido demais ocorre o fenômeno de _flicker_ que é o que vocês devem ter percebido agora.

O nosso jogo não possui uma animação muito complexa, então é suficiente atualizar a tela 10 vezes por segundo. O truque para fazer isso é a função `dorme`.

Esta função faz com que o programa fique parado pelo número de milisegundos que passarmos como parâmetro. Passando o valor de 100 milisegundos nós conseguimos fazer com que o _loop_ principal seja executado 10 vezes por segundo.

Adicione a declaração de `dorme` antes da função `main`:

```
func dorme(milisegundos time.Duration) {
  time.Sleep(time.Millisecond * milisegundos)
}
```

E adicione a sua chamada abaixo do comentário `// Dorme` dentro do _loop_ principal:

```
// Dorme
dorme(100)
```

Você também vai precisar do _import_ que define `time`:

```
import "time"
```

Teste novamente o programa. O efeito de _flicker_ deve ter sumido.

## Passo 09: Melhorar o gráfico

_TODO_

## Passo 10: Adicionar pastilhas e pontos

_TODO_

## Passo 11: Adicionar fim de jogo

_TODO_

## Passo 12: Verificar colisões

_TODO_

## Passo 13: Adicionar vidas

_TODO_

## Passo 14: Adicionar cogumelos de força

_TODO_

## Passo 15: Adicionar suporte a novos mapas

_TODO_
