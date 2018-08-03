package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
//	"strings"
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

	flag := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		//reading particular values from file
		if len(s) == 0 {
			flag = false
		}
		if flag {
			code := s[2:5]
			name := s[37:70]
			deduc := s[70:80]
			remarks := s[84:111]
			fmt.Println(code)
			fmt.Println(name)
			fmt.Println(deduc)
			fmt.Println(remarks)
		}

		if len(s) > 2 {
			if s[2:6] == "CODE" {
				flag = true
				continue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
