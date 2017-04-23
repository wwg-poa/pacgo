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

Você deve ter visto a mensagem "Olá Go" sendo impressa na tela. Estamos prontas para começar!

## Passo 02: Estrutura de um jogo

Neste tutorial nós vamos construir um jogo chamado PacGo. O nome é uma brincadeira com o clássico PacMan.

Para quem não conhece, o seu objetivo é controlar o PacGo por um labirinto que está repleto de fantasmas. Nos corredores do labirinto existem pastilhas que contam pontos quando o PacGo come elas.

Em alguns pontos estratégicos, existem cogumelos de força que tornam o PacGo temporariamente invencível aos fantasmas (e pode comê-los também, valendo pontos). O objetivo do jogo é comer todas as pastilhas do labirinto antes de perder todas as vidas.

O primeiro passo no desenvolvimento de um jogo é o chamado _game design_, que é onde estabelecemos as regras e objetivos do jogo.

Como estamos emprestando a idéia do PacGo de um jogo clássico, vamos pular esta etapa e partir direto para a codificação.

**_Coach_: explicar brevemente o jogo Pac Man observando os aspectos de _game design_**

Crie uma pasta chamada `pacgo` para separar o código do jogo dos outros arquivos. Lembra como fazer? Se não, volte na seção anterior.

Agora vamos criar o arquivo `main.go` onde vai ficar a parte principal do nosso programa. Abra o editor de textos, copie e cole o seguinte código:

```
package main

func main() {
  // Inicializar terminal

  // Inicializar labirinto

  // Loop principal
  for {
    // Desenha tela

    // Processa entrada do jogador

    // Processa movimento dos fantasmas

    // Processa colisões

    // Dorme
  }
}
```

Salve o arquivo. (Lembre-se sempre de salvar o arquivo após cada alteração!)

**_Coach_: explicar o que são comentários, a função `main` e o que é um _loop_.**

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

Ao executar o `pacgo` você vai reparar que o programa parece ter **travado** o terminal, porém este é o *loop* infinito do jogo. Para sair do programa, você pode pressionar as teclas *control* (`Ctrl`) e a letra `C` simultaneamente (escrevemos `Ctrl+C` para abreviar). Lembre-se deste atalho, vamos usá-lo muitas vezes ao longo do tutorial.

## Passo 03: Construindo um labirinto

A nossa primeira tarefa de codificação vai ser desenhar um labirinto na tela.

Copie e cole o código abaixo logo após a linha `package main` (primeira linha do arquivo):

```
import "fmt"
import "os"
import "os/exec"
import "time"
```

**_Coach_: explique em poucas palavras o que é um _import_ e o que são bibliotecas.**

O nosso primeiro passo vai ser preparar o terminal para funcionar como a nossa tela. Para isto, precisamos utilizar algumas funções de baixo nível do sistema. Não se preocupe em tentar entender elas agora. Copie e cole as funções abaixo antes da definição da função `main`:

```
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
```

Na função `main`, abaixo do comentário `// Inicializa terminal`, inclua as seguintes linhas de código:

```
inicializa()
defer finaliza() // executa apenas no fim do programa
```

Agora nós vamos criar uma representação do labirinto no programa. Para isso vamos utilizar uma `struct`. As _structs_ são a nossa forma de dizer que uma coisa possui várias partes, ou "propriedades". No caso, o nosso labirinto possui uma `largura`, uma `altura` e um `mapa`.

No arquivo `main.go` adicione o seguinte código entre a linha 1 e a linha 3:

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

Agora altere a função `main` com a chamada para as duas funções criadas acima colocando-as logo abaixo dos respectivos comentários. Além disso coloque a palavra `break` abaixo do comentário `// Processa entrada do jogador`. O seu código vai ficar assim:

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
    break

    // Processa movimento dos fantasmas

    // Processa colisões

    // Dorme
  }
}
```

Execute agora o seu código:

```
go run main.go
```

Note que ele imprimiu o labirinto e saiu do programa. Isso é porque colocamos a palavra `break` para quebrar o _loop_ infinito. Experimente tirar esta palavra e ver o que acontece. (Lembre-se que neste caso a combinação de teclas para parar o programa é `Ctrl+C`)

O que você deve ter observado é que sem a palavra `break` dentro do _loop_ (iniciado pela palavra-chave `for`) o programa imprime infinitas vezes o mesmo mapa e a tela fica "rolando" indefinidamente.

Vamos corrigir este comportamento adicionando uma função para limpar a tela antes de imprimir o mapa. Copie e cole o código a seguir antes da sdeclaração da função `desenhaTela()`:

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

func dorme(milisegundos time.Duration) {
  time.Sleep(time.Millisecond * milisegundos)
}
```

O código acima define três funções auxiliares: `moveCursor()`, `limpaTela()` e `dorme()`.

Pense no cursor como a "caneta" que escreve na tela. A função `moveCursor` diz para o computador onde é a próxima posição da tela onde ele deve escrever.

A função `limpaTela` apaga todo o conteúdo do terminal e reposiciona o cursor na posição (0, 0), que é o canto superior esquerdo da tela.

A função `dorme` serve para fazer o computador ficar parado por algum tempo sem processar nada. Nós vamos utilizar esta função para evitar que a tela seja atualizada muito rapidamente, o que causa o efeito da tela ficar piscando.

**_Coach_: explicar como funciona o sistema de coordenadas da tela.**

Não se preocupe com o código dentro das aspas na chamada de função `fmt.Printf()`. Estes são códigos de controle que têm funções especiais. Vale lembrar que pouca gente decora estes códigos - existem tabelas prontas na internet com a lista dos códigos e suas funções.

Agora altere a função `desenhaTela()` para incluir uma chamada para `limpaTela()` antes de imprimir o mapa:

```
func desenhaTela() {
  limpaTela() // adicione esta linha
  for _, linha := range labirinto.mapa {
    fmt.Println(linha)
  }
}
```

Remova a palavra `break` do _loop_ principal e coloque a chamada `dorme(100)` logo após a linha com o comentário `// Dorme`. Este trecho do código vai ficar assim:

```
// Loop principal
for {
  // Desenha tela
  desenhaTela()

  // Processa entrada do jogador

  // Processa movimento dos fantasmas

  // Processa colisões

  // Dorme
  dorme(100)
}
```

Execute novamente o programa. Pode parecer que voltamos ao começo da lição, mas na verdade estamos prontas para fazer animações. A tela parece parada, mas está sendo atualizada 10 vezes por segundo, porém sempre com a mesma imagem.

Lembre-se de sair do programa com `Ctrl+C`.

## Passo 04: Mover o PacGo

Agora que nós temos a estrutura de animação pronta, podemos começar a pensar em mover o nosso PacGo (atualmente representado pelo `G` no mapa).

Para facilitar o controle do PacGo ao longo de todo o programa, vamos primeiro criar uma estrutura para representá-lo. Cole o código abaixo da definição da estrutura `Posicao`:

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

Nós criamos a função para construir o PacGo (`criarPacGo`), mas falta chamar esta função dentro do nosso programa. Vamos alterar a função `inicializarLabirinto` para construir o PacGo com a sua posição correta no mapa.

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
type Movimento int

const (
  Cima = iota
  Baixo
  Esquerda
  Direita
  Nenhum
  Sair
)

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
  }
}
```

A função `moverPacGo` recebe um sinal de movimento e tenta atualizar a posição atual do PacGo. Porém, se a nova posição cair numa parede (representada pelo caractere `#`) a função ignora o movimento.

Agora precisamos definir a função que gera este sinal. Para saber a intenção de movimento da pessoa que está jogando nós precisamos saber que tecla ela pressionou. Este processo é chamado de "entrada do usuário".

A função abaixo tem o objetivo de pegar a entrada do usuário e emitir um sinal de movimento. Copie e cole este código abaixo da definição da função `moverPacGo`:

```
func entradaDoUsuario(canal chan<- Movimento) {
  array := make([]byte, 10)

  for {
    lido, _ := os.Stdin.Read(array)

    if lido == 1 && array[0] == 0x1b {
      canal <- Sair;
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
```

O próximo passo é alterar o programa principal para chamar esta função toda vez que alguém pressionar uma tecla. Copie e cole o código abaixo na função `main`, após a chamada da função `inicializarLabirinto`:

```
  canal := make(chan Movimento, 10)
  go entradaDoUsuario(canal)

  var tecla Movimento
```

Ainda na função `main`, copie e cole o código abaixo na dentro do _loop_ principal, abaixo do comentário `// Processa entrada do jogador`:

```
// Processa entrada do jogador
select {
  case tecla = <-canal:
    moverPacGo(tecla)
  default:
}
if tecla == Sair { break }
```

O último passo vai ser alterar a função `desenhaTela` para atualizar a posição do PacGo a cada passada:

```
func desenhaTela() {
  limpaTela() // adicione esta linha
  for _, linha := range labirinto.mapa {
    fmt.Println(linha)
  }

  // Imprime PacGo
  moveCursor(pacgo.posicao)
  fmt.Printf("%s", pacgo.figura)

  // Move cursor para fora do labirinto
  moveCursor(Posicao{labirinto.altura + 2, 0})
}
``

## Passo 05: Mover os fantasmas

## Passo 06: Melhorar o gráfico

## Passo 08: Adicionar pastilhas e pontos

## Passo 09: Adicionar fim de jogo

## Passo 10: Verificar colisões

## Passo 11: Adicionar vidas

## Passo 12: Adicionar cogumelos de força
