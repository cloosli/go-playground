package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type item struct {
	StoryURL string
	Source   string
	Comments string
	Title    string
}

func main() {
	stories := []item{}
	users := map[string]bool{}

	outputDir := filepath.Join("music", "bendsound")
	os.MkdirAll(outputDir, os.ModePerm)

	u := colly.NewCollector(
		colly.MaxDepth(1),
		colly.CacheDir("./_cache"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"),
	)
	// Set max Parallelism and introduce a Random Delay
	u.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	u.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.Path)
	})

	u.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	u.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		if strings.Index(r.Headers.Get("Content-Type"), "application/octet-stream") > -1 {
			fmt.Println("download music")
			u := r.Request.URL

			filename := u.Path[strings.LastIndex(u.Path, "/"):]
			r.Save(outputDir + "/" + filename)

		}
	})

	u.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasSuffix(link, ".mp3") {
			temp := item{}
			temp.StoryURL = link
			stories = append(stories, temp)
			if !users[link] {
				users[link] = true
				u.Visit(e.Request.AbsoluteURL(link))
			}
		}
	})

	u.OnHTML("#products_grid > div.pagenavi2 > div > a.current2", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		u.Visit(e.Request.AbsoluteURL(link))
	})

	u.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL.Path)
	})

	u.Visit("https://www.bensound.com/royalty-free-music?sort=p.date_added&order=DESC")

	u.Wait()
}
