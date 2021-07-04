package main

import (
	"flag"
	"fmt"
	"github.com/NamDuongkiwi/GoLab/Lec4/models"
	"log"
	"os"
)



func NewEbooks() *models.Ebooks {
	return &models.Ebooks{}
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}
	currentUrl := args[0] // Đây là biến lấy ra URL mà bạn muốn crawl data
	fmt.Println(currentUrl)
}