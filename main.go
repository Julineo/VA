package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
)

func main() {
	files, err := ioutil.ReadDir("./input/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		readFile(f.Name())
	}
}

func readFile(s string) {
	fmt.Println("./input/" + s)
	file, err := os.Open("./input/" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
