package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

type item struct {
	Date     string
	Name     string
	Distance string
	City     string
}

func main() {
	events := []item{}

	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.CacheDir("./_cache"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"),
	)
	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 1,
		RandomDelay: 5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Visited", r.Request.URL)
	})

	//next page
	// c.OnHTML("div.ea_eb > div.ea_eh > div.ea_ah > div > div.u9_n0.u9_va.u9_vb > a[rel=next]", func(e *colly.HTMLElement) {
	c.OnHTML("a[rel=next].u9_vc", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		u, _ := url.Parse(e.Request.AbsoluteURL(nextPage))
		q := u.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page > 1 {
			fmt.Println("found next page", page, u.Path)
			c.Visit(e.Request.AbsoluteURL(nextPage))
		}
	})
	// events
	c.OnHTML("li[data-testing-id=event-card].v4_g", func(e *colly.HTMLElement) {
		temp := item{}
		date, _ := time.Parse("2006-1-2", e.ChildAttr("div > div.v4_v7.v4_v5", "content"))
		temp.Date = date.Format("2006-01-02")
		temp.Name = e.ChildText("h3[itemprop=name]")
		temp.Distance = e.ChildText("div > div.v4_wa.v4_v5 > div:nth-child(3)")
		temp.City = e.ChildText("span[itemprop=addressLocality]")
		events = append(events, temp)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.letsdothis.com/gb/running-events?geolocation=false&utm_source=runners_world_gb&resultsPerPage=24&page=1&selectedLocationName=United%20Kingdom&viewportNorthEastLat=60.9&viewportNorthEastLong=2.1&viewportSouthWestLat=49.8&viewportSouthWestLong=-8.9&sort=%7B%22date%22%3A%22asc%22%7D&sortOption=date-0")

	c.Wait()
	// fmt.Println("\n\n\n events:")
	// for _, story := range events {
	// 	fmt.Printf("%#v\n", story)
	// }
	writeToFile(events)
}

var header = []string{"Date", "Name", "Distance", "City"}

func writeToFile(events []item) {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(header)
	for _, value := range events {
		array := []string{value.Date, value.Name, value.Distance, value.City}
		err = writer.Write(array)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
