package crawler

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Crawler struct {
}

func (c *Crawler) Load(url string) (*goquery.Document, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}

	req, err := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return goquery.NewDocumentFromReader(res.Body)
}

func NewCrawler() *Crawler {
	return &Crawler{}
}
