package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	res, err := http.Get("https://www.leagueoflegends.com/es-es/champions/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".style__Text-n3ovyt-3").Each(
		func(i int, s *goquery.Selection) {
			i++
			fmt.Printf("Número: %d Nombre: %v\n", i, s.Text())
		})
}
