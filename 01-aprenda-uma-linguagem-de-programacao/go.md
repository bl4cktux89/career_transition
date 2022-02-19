# anotações

package main | function main
:bulb:
> *aqui você começa, aqui você termina*
> 

---

### usando funções

package.identificator

---

### scope

package level

function level "code block"

---

### tipos

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

### package fmt

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

### tipo boolean

- é um tipo binário que só pode conter dois valores:
    - true
    - false

> sempre que você ver operadores relacionais, o resultado da operação será um valor booleano
> 

é utilizado em: lógica condicional, declarações switch, declarações if, fluxo de controle

---

### tipos numéricos

- INT - integer | inteiros
- FLOAT - floating point | frações

> INTx - tamanho da palavra suportada pelo processador, podendo ser 32 ou 64 bits
> 

> regra geral: utilizar INT, para inteiros, e FLOAT64 para frações
> 

---

### caracteres

- caracteres são feitos de runes - conjunto de bits - apelido para INT32

---

### constantes

- valores imutáveis
- podem ser tipadas ou não
- as constantes não tipadas só vão ter um tipo associado quando forem utilizadas - flexibilidade conveniente

> IOTA - cria uma sequência de constantes integer - inicia com 0 - _ descarta valores
> 

---

### deslocamento de bits

- '<<' desloca os bits à esquerda
- '>>' desloca os bits à direita