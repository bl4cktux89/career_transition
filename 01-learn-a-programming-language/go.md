package main
function main
> aqui você começa
> aqui você termina

usando funções
package.identificator

scope
package level
function level "code block"

tipos são estáticos
static type language

valor 0 "zero":
int: 0
float: 0.0
boolean: false
strings: ""
pointers, functions, interfaces, slices, channels, maps: nil

usar sempre que possível o operador curto de declaração (marmota) ':='

package fmt

strings - interpreted = "" | raw = ''

print

group #1:
print -> standard out
func Print

group #2:
print -> string, pode ser usado como variável
func Sprint

group #3:
print -> file, writer interface, e.g., arquivo ou resposta de servidor
func Fprint