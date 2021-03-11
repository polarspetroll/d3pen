package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(help)
		return
	}
	url := os.Args[1]
	if string(url[len(url)-1]) != "/" {
		url += "/"
	}
	fmt.Printf("looking for %v\n\n", url)
	wordlist, err := os.Open(os.Args[2])
	if err != nil {
		panic(err.Error())
	}
	defer wordlist.Close()
	scan := bufio.NewScanner(wordlist)
	for scan.Scan() {
		resp, err := http.Get(url + scan.Text())
		if err != nil {
			panic(err.Error())
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(body), "dependencie") == true {
			fmt.Printf("%v ⎷\n", url+scan.Text())
		} else {
			fmt.Printf("%v X\n", url+scan.Text())
		}
	}
}

var help string = `
Usage :
d3pen <url> <wordlist>
Example : d3pen http://google.com/ list.txt`
