package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := map[string]any{}
	contents, err := os.ReadFile("testdata/sample.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(contents, data)
}
