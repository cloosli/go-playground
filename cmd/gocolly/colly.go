package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type item struct {
	StoryURL string
	Source   string
	Comments string
	Title    string
}

func main() {
	stories := []item{}
	outputDir := filepath.Join("pictures", "inventables")
	os.MkdirAll(outputDir, os.ModePerm)

	c := colly.NewCollector(
		// colly.AllowedDomains("inventables.com"),
		colly.MaxDepth(1),
		colly.CacheDir("./coursera_cache"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"),
	)
	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	detailCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	detailCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
			u := r.Request.URL
			q := u.Query()
			path := q.Get("frompath")
			if path != "" {
				newPath := filepath.Join(outputDir, path)
				os.MkdirAll(newPath, os.ModePerm)
				filename := u.Path[strings.LastIndex(u.Path, "/"):]
				r.Save(newPath + "/" + filename)
				return
			}
		}
	})

	c.OnHTML("a[href].collection-container", func(e *colly.HTMLElement) {
		temp := item{}
		temp.Title = e.ChildText("span.collection-title")
		temp.StoryURL = e.Request.AbsoluteURL(e.Attr("href"))
		temp.Source = e.ChildAttr("img.collection-thumb-large", "src")
		stories = append(stories, temp)
		c.Visit(temp.StoryURL)
	})

	// collections
	c.OnHTML("div.collection-container", func(e *colly.HTMLElement) {
		detailUrl := e.ChildAttr("a[href]", "href")
		detailCollector.Visit(detailUrl)
	})
	// detail
	detailCollector.OnHTML("div.item > img", func(e *colly.HTMLElement) {
		u, err := url.Parse(e.Attr("src"))
		if err != nil {
			log.Fatal(err)
		}
		q := u.Query()
		q.Set("frompath", e.Request.URL.Path)
		u.RawQuery = q.Encode()
		c.Visit(u.String())
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.inventables.com/projects")

	c.Wait()
	// for _,story := range stories {
	// 	fmt.Printf("Story: %v\n", story)
	// }
}

func downloadImage(path string, filename string, url string) {
	img, _ := os.Create(filepath.Join(path, filename))
	defer img.Close()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	b, _ := io.Copy(img, resp.Body)
	fmt.Println("File size: ", b)
}
