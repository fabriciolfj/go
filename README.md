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