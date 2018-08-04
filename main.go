package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	files, err := ioutil.ReadDir("./input/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		readFile(f.Name())
	}
	
	readDict()
}

func readDict() error {
	var mySlice [][][]string
	var value string
	mySlice, err := xlsx.FileToSlice("dict.xlsx")
	if err != nil {
		log.Println(err)
		return err
	}
	value = mySlice[0][0][0]
	_ = value
	fmt.Println(value)
	return nil
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

		if len(s) > 5 {
			if s[2:6] == "CODE" {
				flag = true
				continue
			}
		}

		if len(s) == 0 {
			flag = false
			continue
		}
		if len(s) < 82 {
			flag = false
		}

		//reading particular values from file
		if flag {
			code := strings.Trim(s[2:5], " ")
			name := strings.Trim(s[37:70], " ")
			deduc := strings.Trim(s[70:80], " ")
			remarks := strings.Trim(s[84:111], " ")
			fmt.Println(code)
			fmt.Println(name)
			fmt.Println(deduc)
			fmt.Println(remarks)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
