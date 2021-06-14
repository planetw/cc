package crawler

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func (c *Crawler) ScrapePlayer() {
	doc, err := c.Load("https://www.cpbl.com.tw/player")
	if err != nil {
		log.Fatal(err)
	}

	doc.
		Find(".PlayersList dl").
		Children().
		Each(func(i int, s *goquery.Selection) {
			switch node := goquery.NodeName(s); node {
			case "dt":
				fmt.Printf("Team Name: %s", s.Text())
			case "dd":
				fmt.Printf("Player Name: %s", s.Text())
			}
		})
}
