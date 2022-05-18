package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("укажите url")
		os.Exit(1)
	}
	req, err := http.Get(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
	} else {
		defer req.Body.Close()
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			file, err := os.Create("dev09.txt")
			if err != nil{
				fmt.Println("Unable to create file:", err) 
				os.Exit(1) 
			}
			defer file.Close() 
			file.WriteString(string(b))
		}
	}
}
	