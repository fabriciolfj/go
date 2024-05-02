# GO

## Slice
```
package main

import "fmt"

func main() {
	//compartilha a mesma memoria
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//pega o valor 1 2
	y := x[:2:5]

	fmt.Println(y)

	//coloca no final do tamanho, mas modifica o valor de x por é compartilhado, substituindo os valores nas respectivas posicoes
	y = append(y, 11, 12, 13)

	fmt.Println(y)
	fmt.Println(x) //como foi alterado o resultado agora e [1 2 11 12 13 6 7 8 9 10]
}
```

- funcao de limpeza, chama-se defer
```
import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
```
- nesse caso estamos adiando o fechamento do arquivo até o final da func main, lembrando que o defer funciona como LILO
- ultimo que foi declarado para adiar, será o primeiro a ser executado

## mudança de valores
- se passarmos primitivos ou structs para uma funcao, e essa modifica-os, não replete nos valores originais, pois o go passa uma cópia dos mesmos.
- se passarmos um map ou um slice para uma função, e esta modifica-os, eles são alterados, pois go passa um ponteiro.
- para alterar o parâmetro passado, usamos pontiero como parâmetro:
```

func update(px *int) {
	*px = 20 //referencia ao ponteiro da variavel externa passada para essa funcao
}
```
- para slice, quando modifico valores dentro da capacidade da variavel original (tamanho 3, mudo index 0 1 e 2), são refletidos na mesma
- quando adiciono valores, deixando slice maior que a original, esses não são repletidos no original
```
func addValues(list []int) {
	list = append(list, 4, 5, 6)
	list[0] = 9
	list[1] = 8
	list[2] = 7
}

func main() {
	list := make([]int, 3, 6)

	fmt.Println(list)

	addValues(list)
	fmt.Println(list)
}

```

# Ponteiros via parâmetro
- ideal quando usamos interface
```
func failUpdate(px *int) {
	fmt.Println(px)
	x2 := 20
	px = &x2
}
```
- se temos struct pequenos, é mais rápido passar o valor via parâmetro
- no caso de struct com valores grandes, como 10mb, considere passar o ponteiro em vez do valor por questão de performance


# garbage collection
- em go gc e otimizado para baixa latência, ou seja, rodar rápido e limpar a maior quantidade de memória possível
- é mais rapido limpar a memoria sem o uso de ponteiros, por isso go não recomenda o uso excessivo do mesmo
- ponteiros são armazenados no heap, exigindo mais do gc, 
- variaveis ficam no heap, facilitando o trabalho do gc, pois são armazenadas na memória de forma sequencial e não espalhada como os ponteiros
- em java usa-se muito ponteiro, (ex: cada valor de uma list é um ponteiro), o que torna inificiente esse ponto em relação ao go

# methods e funcoes
- metodos estão atrelados  a um tipo, e este é invocuado apenas por ele, e se limita sua declarao a nivel de pacote
- funcoes sao mais genericas e não estão atrelados a um tipo

# ponteiros em methods
- quando a funcao recepciona um ponteiro e e este possui um método que também recepciona um ponteiro, o valor será alterado quando invoca-lo
- quando a funcao receptiona valor, e este possui um método que recepciona ponteiro, seu valor nao será alterado caso invoque-o
````
func executeIncrement(c Counter) {
	c.Increment()
}

func executeIncrementModify(c *Counter) {
	c.Increment()
}

func main() {
	fmt.Println("inicio")
	c := Counter{}
	executeIncrement(c)
	fmt.Println("executeIncrement", c.String())
	executeIncrementModify(&c)
	fmt.Println("executeIncrementModify", c.String())
}
````
- um ponto importante, tanto para funcoes e métodos, para checar um valor nulo, deve-se receber um ponteiro e o valor passao também
`````
	var z *Counter //consegui verificar nulo
	testNuloOrZeo(z)

	var p Counter
	testNuloOrZeo(&p) //nao consegui

	func testNuloOrZeo(c *Counter) {
	if c == nil {
		fmt.Println("veio nulo")
		return
	}

	c.Increment() //se chegar aqui como valor, nao faz nada se nao for instanciado
	fmt.Println("executou")
}
`````