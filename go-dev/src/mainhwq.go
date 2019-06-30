// web project main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func GetJokes() {

	file := "E://news.txt"
	fout, err := os.Create(file)
	defer fout.Close()
	if err != nil {
		fmt.Println(file, err)
		return
	}
	doc, err := goquery.NewDocument("http://news.qq.com")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".linkto").Each(func(i int, s *goquery.Selection) {
		fout.WriteString(s.Text())
		fout.WriteString("\n")
	})
}
func main() {

	GetJokes()
}
