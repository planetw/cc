package crawler

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Team struct {
	Name string
	Url  string
}

func (c *Crawler) ScrapeTeam() {
	doc, err := c.Load("https://www.cpbl.com.tw")
	if err != nil {
		log.Fatal(err)
	}

	doc.
		Find("#Footer div.teams ul li a").
		Each(func(i int, s *goquery.Selection) {

			t := &Team{
				Name: utils.getAttr("title"),
				Url:  utils.getAttr("href"),
			}

			fmt.Printf("%v", t)
		})
}
