# 4
---
## introdução
---
## aprenda o básico
### sintaxe básica
### variáveis e declarações
### tipos de dados
### bool
### byte
### rune
### complex64/128
### int, int8/16/32/64
### uint, uint8/16/32/64
### float32, float64
### uintptr
### for loop
### range
### if, switch statements
### errors, panic, recover
### conditionals
### functions, multiple/named returns
### packages, imports e exports
### type casting
### type inference
### arrays
### slices
### maps
### make()
### structs
---
## going deeper
### go modules
### types, type assertions, switches
### interfaces
### context
### goroutines
### channels
### buffer
### select
### mutex
### marshaling & unmarshaling JSON
---
## building CLIs
### cobra
---
## ORMs
### GORM
---
## web frameworks
### beego
### gin
### revel
### echo
### gorilla
### buffalo
---
## logging
### logrus
### zap
---
## real time communitation
### melody
### centrifugo
---
## API clients
### REST
#### heimdall
#### GRequests
### graphQL
#### graphql-go
#### gqlgen
---
## testando suas aplicações
---
## ferramentas para microsserviços
### watermill
### rpcx
### go-kit
### micro
### go-zero
### protocol buffers
### gRPC-Go
### gRPC-gateway
### twirp
---

## referências

Documentação da Linguagem: <a href="https://go.dev/doc/">GO</a></br>
playlist da Ellen Korbes: <a href="https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg">Aprenda Go</a></br>
playlist do Tiago Temporin: <a href="https://www.youtube.com/playlist?list=PLHPgIIn9ls6-1l7h8RUClMKPHi4NoKeQF">Aprenda Golang</a></br>

---
package main | function main
> *aqui você começa, aqui você termina*
> 

---

# usando funções

package.identificator

---

# scope

package level

function level "code block"

---

# tipos

tipos são estáticos

static type language

valor 0 "zero":

int: 0

float: 0.0

boolean: false

strings: ""

pointers, functions, interfaces, slices, channels, maps: nil

usar sempre que possível o operador curto de declaração (marmota) ':='

---

# package fmt

strings - interpreted = "" - aspas | raw = `` - crases

print

- group #1:
    
    print -> standard out
    
    func Print
    
- group #2:
    
    print -> string, pode ser usado como variável
    
    func Sprint
    
- group #3:
    
    print -> file, writer interface, e.g., arquivo ou resposta de servidor
    
    func Fprint
    

---

# tipo boolean

- é um tipo binário que só pode conter dois valores:
    - true
    - false

> sempre que você ver operadores relacionais, o resultado da operação será um valor booleano
> 

é utilizado em: lógica condicional, declarações switch, declarações if, fluxo de controle

---

# tipos numéricos

- INT - integer | inteiros
- FLOAT - floating point | frações

> INTx - tamanho da palavra suportada pelo processador, podendo ser 32 ou 64 bits
> 

> regra geral: utilizar INT, para inteiros, e FLOAT64 para frações
> 

---

# caracteres

- caracteres são feitos de runes - conjunto de bits - apelido para INT32

---

# constantes

- valores imutáveis
- podem ser tipadas ou não
- as constantes não tipadas só vão ter um tipo associado quando forem utilizadas - flexibilidade conveniente

> IOTA - cria uma sequência de constantes integer - inicia com 0 - _ descarta valores
> 

---

# deslocamento de bits

- '<<' desloca os bits à esquerda
- '>>' desloca os bits à direita

---

# fluxo de controle

é a forma como o computador lê e executa um programa:

- computadores geralmente lêem em sequência - de cima para baixo
- além do fluxo sequencial existem outras duas declarações que podem afetar a maneira do computador ler código:
    - fluxo de controle de repetição - loop | iterativo
    - fluxo de controle condicional - seleção

> existem 3 tipos de fluxo de controle: sequencial, repetição e condicional
> 

# loop

estágios do for:

- inicialização
- condição
- pós

break - quebra o loop

continue - pula a iteração atual

# if

só executa se for true