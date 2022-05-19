package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Реализовать утилиту wget с возможностью скачивать сайты целиком.

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("укажите url") 
		os.Exit(1)
	}
	req, err := http.Get(flag.Args()[0]) // по заданному урлу
	if err != nil {
		fmt.Println(err)
	} else {
		defer req.Body.Close()
		b, err := ioutil.ReadAll(req.Body)// считываем боди нашего реквеста GET
		if err != nil {
			fmt.Println(err)
		} else {
			file, err := os.Create("dev09.html") // записываем в файл 
			if err != nil{
				fmt.Println("Unable to create file:", err) 
				os.Exit(1) 
			}
			defer file.Close() 
			file.WriteString(string(b))
		}
	}
}
	