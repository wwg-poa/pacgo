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

#### Instalação Linux

#### Instalação Windows

### Meu primeiro programa

Existe uma tradição na área da computação em que toda vez que você vai começar a aprender uma linguagem nova você comece escrevendo um programa chamado "Hello world".

Nós costumamos fazer isso porque geralmente é bem simples de fazer e nos ajuda a testar se a instalação da linguagem está ok.

Então vamos fazer o mesmo para a linguagem Go.

Primeiro, crie uma pasta no seu computador onde você vai guardar os códigos que escrever. Você pode dar qualquer nome para ela, mas de preferência sem acentos ou espaços. Por exemplo: `tutorial`.

Digite no terminal:

```
mkdir tutorial
cd tutorials
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

O primeiro passo no desenvolvimento de um jogo é o chamado "Game Design", que é onde estabelecemos as regras e objetivos do jogo. Como estamos emprestando a idéia do PacGo de um jogo clássico, vamos pular esta etapa e partir direto para a codificação.

Crie uma pasta chamada `pacgo` para separar o código do jogo dos outros arquivos. Lembra como fazer? Se não, volte na seção anterior.

Agora vamos criar o arquivo `main.go` onde vai ficar a parte principal do nosso programa.
