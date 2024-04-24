package main

import "fmt"

func main() {
	words := []string{"fabricio", "lucas"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right lenght:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long world!")
		}
	}

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 1, 2, 3, 4, 5:
			fmt.Println("ok")
		case 10, 6:
		case 7:
			fmt.Println("saindo")
			break loop
		default:
			fmt.Println("invalido")
		}
	}
}
