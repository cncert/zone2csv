package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func readFile(strChan chan string) {
	start1 := time.Now()
	filePath := "com.file"
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	lineScanner := bufio.NewScanner(f)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		// fmt.Println(line)
		strChan <- line
	}
	fmt.Println(time.Now().Sub(start1))
}

func writefile(strChan chan string) {
	f, err := os.OpenFile("com.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	wr := csv.NewWriter(f)

	for line := range strChan {

		newStr := strings.SplitN(line, " ", 5)
		if len(newStr) > 4 {
			row := []string{newStr[0], newStr[3], newStr[1], newStr[2], newStr[4]}
			err := wr.Write(row)
			if err != nil {
				fmt.Println(err)
			}
		}
		row := []string{newStr[0], newStr[3], newStr[1], newStr[2]}
		errWrite := wr.Write(row)
		if errWrite != nil {
			fmt.Println(errWrite)
		}
		wr.Flush()

	}

}

func main() {
	strChan := make(chan string)
	go readFile(strChan)
	writefile(strChan)
}
