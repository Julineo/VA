package main

import (
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"strings"

	"github.com/tealeg/xlsx"
)

var vs [][][]string

func main() {
	files, err := ioutil.ReadDir("./input/")
	if err != nil {
		log.Fatal(err)
	}

	vs, _ = readDict()

	for _, f := range files {
		readFile(f.Name())
	}


}

func readDict() ([][][]string, error) {
	var mySlice [][][]string
	mySlice, err := xlsx.FileToSlice("dict.xlsx")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return mySlice, nil
}

func readFile(s string) {

	file, err := os.Open("./input/" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//creating xlsx file
	out := xlsx.NewFile()
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	sheet, err = out.AddSheet("Sheet1")
	if err != nil {
		log.Println(err.Error())
	}

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "Id"
	cell = row.AddCell()
	cell.Value = "LAST NAME"
	cell = row.AddCell()
	cell.Value = "FIRST NAME"
	cell = row.AddCell()
	cell.Value = "MI"
	cell = row.AddCell()
	cell.Value = "AMT DED"
	cell = row.AddCell()
	cell.Value = "REMARKS"

	flag := false
	var code string

	scanner := bufio.NewScanner(file)
	for {
		start:
		next := scanner.Scan()
		if !next {
			break
		}
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
			code = strings.Trim(s[2:5], " ")
			name := strings.Trim(s[37:70], " ")
			deduc := strings.Trim(s[70:80], " ")
			remarks := strings.Trim(s[84:111], " ")
			for _, v := range vs[0] {
				if code == v[0] && name == v[4]	{
					row = sheet.AddRow()
					cell = row.AddCell()
					cell.Value = v[13]
					cell = row.AddCell()
					cell.Value = v[9]
					cell = row.AddCell()
					cell.Value = v[10]
					cell = row.AddCell()
					cell.Value = v[11]
					cell = row.AddCell()
					cell.Value = deduc
					cell = row.AddCell()
					cell.Value = remarks
					goto start
				}
			}
			row = sheet.AddRow()
			cell = row.AddCell()
			for _, s := range strings.Fields(name) {
				cell = row.AddCell()
				cell.Value = s
			}
			cell = row.AddCell()
			cell.Value = deduc
			cell = row.AddCell()
			cell.Value = remarks

		}
	}

	err = out.Save("./output/" + code + ".xlsx")
	if err != nil {
		log.Println(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
