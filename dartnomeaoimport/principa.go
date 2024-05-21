package main

import (
	crand "crypt/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])

	if err != nil {
		panic("error")
	}

	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}

func main() {
	result := seedRand()
	fmt.Println(result)
}
