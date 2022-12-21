package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		var pdf bool = strings.Contains(file.Name(), ".pdf")
		if pdf {
			fmt.Println("Removing:", file.Name())
			err := os.Remove(file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
