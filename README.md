# GO

## Slice
```
package main

import "fmt"

func main() {
	//compartilha a mesma memoria
	x := make([]int, 0, 5) //cria um slice de inteiros, com comprimento 0, e capacidade inicial de 5
	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//pega o valor 1 2
	y := x[:2:5]

	var test int //por default insere 0
	

	fmt.Println(y)

	//coloca no final do tamanho, mas modifica o valor de x por é compartilhado, substituindo os valores nas respectivas posicoes
	y = append(y, 11, 12, 13)

	fmt.Println(y)
	fmt.Println(x) //como foi alterado o resultado agora e [1 2 11 12 13 6 7 8 9 10]
}
```

- explicacao
````
A sintaxe x[:2:5] cria um novo slice a partir do slice x, com os seguintes parâmetros:

x: é o slice base a partir do qual o novo slice será criado.
2: é o índice (não inclusivo) onde o novo slice deve terminar. Neste caso, o novo slice terminará no índice 2 (mas não incluirá o elemento nessa posição).
5: é o comprimento máximo do novo slice. Isso determina a capacidade do novo slice, que é o tamanho máximo que ele pode crescer antes de precisar alocar um novo bloco de memória.

Dessa forma, y := x[:2:5] cria um novo slice y com os seguintes atributos:

Os elementos do slice y serão os elementos de x desde o índice 0 (início implícito) até o índice 2 (não inclusivo).
````

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
- um ponto importante, tanto para funcoes e métodos, para checar um valor nulo, deve-se receber um ponteiro e o valor passado também
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

# interfaces
- as interfaces são dinâmicas em go, ou seja, não precisamos deixa-las explicítas nas implementações
- desde que as implementações respeitem seu contrato
- um ponto que interfaces nao sao comparaveis, quando utiliza-se slides em suas implementações ou ponteiros.
- um ponto importante que a interface e considerada nula, tanto o tipo subjacente quanto o valor, for nulo.

# composição
- quando tenho um struct dentro de outro, e esse interno possui um método, consigo chamar ele pelo struct "pai"
```
func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func main() {
	o := Outer{
		Inner: Inner{A: 10}, S: "hello",
	}

	fmt.Println(o.Double())
}
```

# asserção
- e para verificar o tipo de um any / interface
- é executada em tempo de execução
- diferente da conversão (que muda o valor), que é realizada em tempo de compilação (menos para ponteiros e slide)

# injecao de dependencia
- em go podemos usar o wine

# tipo subjacente
- para usar um tipo, que tenha no seu interior um tipo subjacente, utilize ~
- no exemplo abaixo tenho MyInt, mas uma interface que usa int, posso declarar que ela aceita tipos subjacentes com ˜int
```
type Integer interface {
	˜int
} 

type MyInt int

func metodo[T Integer](num T) {
	--
}

var my MyInt = 10
metodo(10)
```

# enum em go
- go não possui enum, um exemplo seria utilizar o recurso iota
```
package main

import "fmt"

var Status int

const (
	VALID = iota
	INVALID
	ERROR
)

func main() {
	fmt.Println(VALID, INVALID, ERROR)
}

```
- iota preenche as demais constante seguindo o valor, como zero é default para int, invalid será 1 e error será 2 

# error
- em go por padrão temos o error de uma variavel de erro
- quando não corre falha, esta permanece como nil
- o codigo chamador da função pode verificar a variavel error != nil ou ignorar ela com _ (que deixa explícito que estamos ignorando a variavel)
- a mensagem de erro deve ser tudo minuscula e sem pontos no final.
- exemplo de criação de constantes de erro em go
```
package main

import (
    "errors"
    "fmt"
)

var (
    // ErrInvalidInput é retornado quando a entrada é inválida
    ErrInvalidInput = errors.New("entrada inválida")

    // ErrNotFound é retornado quando o item não foi encontrado
    ErrNotFound = errors.New("item não encontrado")
)

func processInput(input string) error {
    if input == "" {
        return ErrInvalidInput
    }
    // Processa a entrada válida
    return nil
}

func main() {
    err := processInput("")
    if err != nil {
        fmt.Println(err) // Imprime "entrada inválida"
        if errors.Is(err, ErrInvalidInput) {
            fmt.Println("Erro de entrada inválida")
        }
    }
}
```

# error sentinela
```
Em Go, um erro sentinela é um valor especial de erro definido em um pacote que indica uma condição específica. É uma prática comum em Go usar valores de erro sentinela para indicar erros específicos conhecidos.
O pacote errors do Go define a função New que cria um novo valor de erro com uma mensagem específica. No entanto, muitas vezes é útil distinguir diferentes tipos de erros para tomar decisões com base neles.
Um exemplo clássico de um erro sentinela é io.EOF, que é retornado por funções de E/S quando elas encontram o fim de um fluxo de dados. Isso permite que o código do aplicativo distinga entre um erro real e o fim normal dos dados.
```

# funcao as no error
```
O operador as é usado para verificar se o valor concreto armazenado em uma interface é do tipo especificado. Ele retorna um valor booleano true se o valor concreto for do tipo especificado, caso contrário, retorna false.O operador as é usado para verificar se o valor concreto armazenado em uma interface é do tipo especificado. Ele retorna um valor booleano true se o valor concreto for do tipo especificado, caso contrário, retorna false.

err := fmt.Errorf("ocorreu um erro")
if errors.As(err, &os.PathError{}) {
    // O valor concreto de err é do tipo *os.PathError
}
```


# fun is no error
```
O operador as é usado para realizar uma asserção de tipo e extrair o valor concreto armazenado em uma interface. Se a asserção for bem-sucedida, o valor concreto é atribuído à variável especificada, e a operação retorna true. Caso contrário, a variável recebe o valor zero do tipo especificado, e a operação retorna false.

Is e comunente usado para verificar erros sentinela se ocorreu

file, err := os.Open("file.txt")
if err != nil {
    if errors.Is(err, os.ErrNotExist) {
        // Arquivo não existe
        fmt.Println("O arquivo não existe")
    } else if errors.Is(err, os.ErrPermission) {
        // Permissão negada
        fmt.Println("Permissão negada para abrir o arquivo")
    } else {
        // Outro erro
        fmt.Println("Erro ao abrir o arquivo:", err)
    }
}
```

# diferenca entre is e as
- is compara seu error por outro forneceido, este pode ser um sentinela (uso dentro de um if)
- compara seu erro pertence a alguma classe concreta a atribuido no ponteiro passado, onde pode ser utilizado dentro do bloco if 

# encapsulando varios erros
- quando tenho varias checagens aonde emito o mesmo erro com a mesma mensagem, posso usar o defer para simplificar o codigo
- lembrando que funcao defer e executada antes do return, na ordem inversa de declaracao (da ultima para primeira)
````
func doThing1() (string, error) {
	return "", errors.New("teste1")
}

func doThing2() (string, error) {
	return "", errors.New("teste2")
}

func doSomeThings() (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in doSomeThings: %w", err)
		}
	}()

	val1, err := doThing1()
	if err != nil {
		return val1, err
	}

	return doThing2()
}

func main() {
	var err error
	var result string
	result, err = doSomeThings()
	fmt.Println(err, result)
}


````

#  recover
- quando temos um falha, uma forma de se recuperar dela e utilizar o mecanismo de recover
- este deve estar dentro de um defer, pois apenas funcoes adiadas são executadas após uma falha
- o uso de recover se reserve a erros fatais e não tratativas de exceptions


# pacotes
- para importar um modulo dentro da classe (que estao em outros subdiretorios), primeiro:
 - crie o go mod init exemple.com/nome
 - no import, supondo que esta nome/package, import ficaria assim exemple.com/nome/package

````
package main

import (
	"fmt"

	"example.com/system/format"

	"example.com/system/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}

````

# pacotes internal
- quando declaramos pacote com o nome internal, dentro de outro package, as funções dentro do mesmo ficam visíveis apenas dentro do package pai e nos pacotes internos dentro do mesmo pai

# dependencia circular
- go não permite dependencia circular

# dependencia indireta
- quando um modulo que importo, importa outras dependencias que ele necessita 

# versoes das dependencias (modulos)
- para baixar as dependências, execute:
```
go get
```
- para baixar uma versão especifica de um modulo ou import
```
go get modulo@versao
```
- para verificar as versoes das dependencias, check o arquivo go.mod
- por default go pega as ultimas versoes

# alguns conceitos
- Módulo: Uma coleção de pacotes Go versionados juntos.
- Dependência de módulo: Um módulo externo que seu módulo depende e importa.

```
Então, quando você diz "eu tenho dois módulos que usam a mesma dependência, mas com versões diferentes", você está se referindo a dois módulos diferentes que dependem de versões diferentes de um terceiro módulo.
O gerenciamento dessas dependências de módulos, incluindo a escolha de versões específicas, é uma das principais tarefas do sistema de Go Modules.
```
- quando criamos um modulo com github.com/seurepo/projeto, ao dar um get ./.., go tentará baixar o modulo de la
- caso ainda não tenha enviado o modulo ap git, vá no diretório raiz e execute work init/work use, para usar o modulo localmente
- nao commite o arquivo go.work
```
go work init workspace_lib (modulo que vou precisar)
go work use workspace_use (modulo que estou usando)

```
- para importar um modulo do git, informe o mesmo como requirido no go.mod do seu modulo
- execute na raiz do seu modulo go ./get...
- caso seu modulo dependa de outro seu, e não queria utilizar local, envio ao github
- para projetos novos, go nao permite versao de dependencia acima de v1
- caso necessite, utilze o replace (obs deve estar em outro usuario)
```
replace github.com/fabriciolfj/workspace_lib => github.com/seu-usuario/workspace_lib v2.0.0

```
- ou para mudar para versao 2, crie um path no repositorio v2, coloque os arquivos dentro dele, incluindo o go.mod atualizado para repositorio/v2
- crie uma tag com esse novo path para v2.0.0
- atualize o local de uso do modulo, acrescento o sufixo v2 e a versao v2.0.2go


# verificador de codigo
- ferramenta ideal é o staticcheck, para instalar: (lembrando de mudar o local da instalação, alterando a variável de ambiente GOBIN)\
```
go install honnef.co/go/tools/cmd/staticcheck@latest
```
- para examinar, execute dentro do seu módulo ./staticcheck ./...
- outra ferramenta para verificar a qualidade do código é a revive
```
go install github.com/mgechev/revive@latest
```
  - podemos inclusive adicionar mais regras de verificação, adicionando um arquivo no local do module com nome build_in.toml
```
[rule.redefines-builtin-id] (para verificar sombreamento de variáveis)
```
- para executar a verificação, execute:
```
revive -config build_in.toml ./...
```
- para ver mais regras olhe https://revive.run/r
- uma ferramenta que possui diversas outras, é o golangci (recomendavel instalar o binario golangci-lint)
- podemos também personalizar sua verificação, colocando um arquivo no local de execução com nome .golangci.yml
```
linters:
  enable:
    - govet
	- predeclared
linters-settings:
  govet:
    check-shadowing: true
	settings:
	  shadow:
	    strict: true
    enable-all: true

```
- para executar: golangci-lint run
- para verificar alguma vulnerabildiade de segurança nas suas libs ou de tercerios, podemos utilizar o govulncheck
```
go install golang.org/x/vuln/cmd/govulncheck@latest
```
- para executar: govulncheck ./...echo