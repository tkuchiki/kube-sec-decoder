package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var isHideData bool
	var replaceData string
	flag.BoolVar(&isHideData, "hide-data", false, "hide .data.*")
	flag.StringVar(&replaceData, "replace-data", "<secret>", "replace .data.* to specified string")
	flag.Parse()
	decoder := NewDecoder(isHideData, replaceData)
	decoded, err := decoder.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(decoded)
}
