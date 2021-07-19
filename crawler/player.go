package crawler

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type player struct {
	Name string
}

type team struct {
	Name    string
	Players []player
}

func (c *Crawler) ScrapePlayer() {
	doc, err := c.Load("https://www.cpbl.com.tw/player")
	if err != nil {
		log.Fatal(err)
	}

	doc.
		Find(".PlayersList dl").
		Each(func(i int, s *goquery.Selection) {
			t := &team{}

			s.
				Children().
				Each(func(i int, s *goquery.Selection) {
					switch node := goquery.NodeName(s); node {
					case "dt":
						t.Name = s.Text()
						break
					case "dd":
						p := player{
							Name: s.Text(),
						}

						t.Players = append(t.Players, p)
						break
					}
				})

			fmt.Printf("%v", t)
		})
}
