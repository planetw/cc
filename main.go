package main

import (
	"github.com/planetw/cc/crawler"
)

func main() {
	crawler := crawler.NewCrawler()
	crawler.ScrapeTeam()
}
