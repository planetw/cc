package utils

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/planetw/cc/crawler/utils"
)

func getAttr(attribute string) string {
	attr, exist := s.Attr(attribute)

	if !exist {
		log.Fatal(attribute)
	}

	return attr
}
