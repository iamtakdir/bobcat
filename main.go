package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) !=1 {
		fileName := os.Args[1]

		fmt.Println(fileName,"\n")
		readAnything(fileName)
	}else {
		fmt.Println("File Name can't be empty, bobcat<space>File Name ")
	}
}

func readAnything (fileName string){

	content, err := os.ReadFile(fileName)
	if err !=nil{
		log.Fatal("Error in Reading file ", err)
	}

	fmt.Println(string(content))
}