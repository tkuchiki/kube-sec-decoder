package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	decoder := NewDecoder()
	decoded, err := decoder.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(decoded)
}
