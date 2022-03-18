# anotações

base do estudo em Go:
https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg

	1. aprenda o básico
		1. sintaxe básica
		2. variáveis e declarações
		3. tipos de dados
		4. bool
		5. byte
		6. rune
		7. complex64/128
		8. int, int8/16/32/64
		9. uint, uint8/16/32/64
		10. float32, float64
		11. uintptr
		12. for loop
		13. range
		14. if, switch statements
		15. errors, panic, recover
		16. conditionals
		17. functions, multiple/named returns
		18. packages, imports e exports
		19. type casting
		20. type inference
		21. arrays
		22. slices
		23. maps
		24. make()
		25. structs
	2. going deeper
		1. go modules
		2. types, type assertions, switches
		3. interfaces
		4. context
		5. goroutines
		6. channels
		7. buffer
		8. select
		9. mutex
		10. marshaling & unmarshaling JSON
	
	3. building CLIs
		1. cobra

	4. ORMs
		1. GORM

	5. web frameworks
		1. beego
		2. gin
		3. revel
		4. echo
		5. gorilla
		6. buffalo
	
	6. logging
		1. logrus
		2. zap
	
	7. real time communitation
		1. melody
		2. centrifugo
	
	8. API clients
		1. REST
			1. heimdall
			2. GRequests
		2. graphQL
			1. graphql-go
			2. gqlgen

	8. testando suas aplicações
	9. ferramentas para microsserviços
		1. watermill
		2. rpcx
		3. go-kit
		4. micro
		5. go-zero
		6. protocol buffers
		7. gRPC-Go
		8. gRPC-gateway
		9. twirp

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

---

### fluxo de controle

é a forma como o computador lê e executa um programa:

- computadores geralmente lêem em sequência - de cima para baixo
- além do fluxo sequencial existem outras duas declarações que podem afetar a maneira do computador ler código:
    - fluxo de controle de repetição - loop | iterativo
    - fluxo de controle condicional - seleção

> existem 3 tipos de fluxo de controle: sequencial, repetição e condicional
> 

### loop

estágios do for:

- inicialização
- condição
- pós

break - quebra o loop

continue - pula a iteração atual

### if

só executa se for true