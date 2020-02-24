package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var strToFind string = "go"

//----------------------getCount-------------------------------
func getCount(url string, c chan int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	kolvo := strings.Count(string(body), strToFind)
	fmt.Println("Count for", url, "-", kolvo)
	c <- kolvo
}

//---------------------------main--------------------------
func main() {

	totalKolvo := 0

	c := make(chan int)

	file, err := os.Open("url.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go getCount(scanner.Text(), c)
		tmp := <-c
		totalKolvo = totalKolvo + tmp
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total: ", totalKolvo)
	//
	var input string
	fmt.Scanln(&input)
}

//---------------------------------------------------------------
