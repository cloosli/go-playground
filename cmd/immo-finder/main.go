package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type comparisQuery struct {
	LocationSearchString string
	RootPropertyTypes    []int
	PriceFrom            int
	PriceTo              int
	RoomsFrom            int
	Sort                 int
}

type item struct {
	Type              string
	Title             string
	ShortDesctription string
	LongDesctription  string
	Place             string
	Price             string
	Images            []string
	URL               string
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
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		fmt.Println("Body", string(r.Body))
	})

	//next page
	// data-css-selector="pagination-item-next"
	c.OnHTML("ul > li[data-css-selector=pagination-item-next] > a[href]", func(e *colly.HTMLElement) {
		print("FOUND NEXT PAGE")
		nextPage := e.Attr("href")
		u, _ := url.Parse(e.Request.AbsoluteURL(nextPage))
		q := u.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page > 1 {
			fmt.Println("found next page", page, u.Path)
			c.Visit(e.Request.AbsoluteURL(nextPage))
		}
		c.Visit(e.Request.AbsoluteURL(nextPage))
	})
	// events
	c.OnHTML("#__next > div > div.css-67pwck.excbu0j1 div a", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if !strings.HasPrefix(e.Attr("href"), "/immobilien/marktplatz/details/show") {
			return
		}
		// date, _ := time.Parse("2006-1-2", e.ChildAttr("div > div.v4_v7.v4_v5", "content"))
		// temp.Date = date.Format("2006-01-02")
		temp := item{}
		temp.URL = e.Request.AbsoluteURL(url)
		temp.Type = e.ChildText("p.css-dsqae5")
		temp.Title = e.ChildText("h3")
		temp.Place = e.ChildText("div > div.v4_wa.v4_v5 > div:nth-child(3)")
		temp.Price = e.ChildText("span.css-1ladu04")
		temp.ShortDesctription = e.ChildText("div>p.css-ejje9")
		events = append(events, temp)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	//https://www.comparis.ch/immobilien/api/v1/singlepage/resultitems?requestObject={"Header":{"Language":"de"}}"
	//c.Visit("https://www.comparis.ch/immobilien/api/v1/singlepage/resultitems?requestObject=%7B%22Header%22:%7B%22Language%22:%22de%22%7D%7D")
	c.Visit("https://www.comparis.ch/immobilien/result/list?requestobject=%7B%22DealType%22%3A%2210%22%2C%22LocationSearchString%22%3A%22%22%2C%22RootPropertyTypes%22%3A%5B%220%22%5D%2C%22PriceTo%22%3A%222200%22%2C%22RoomsFrom%22%3A%224.5%22%2C%22Sort%22%3A%2211%22%2C%22AdAgeMax%22%3A-1%2C%22ComparisPointsMin%22%3A-1%2C%22SiteId%22%3A-1%7D&sort=11")

	c.Wait()
	fmt.Println("\n\n\n properties:")
	for _, story := range events {
		fmt.Printf("%#v\n", story)
	}
	writeToFile(events)
}

var header = []string{"Type", "Title", "Place", "Desc", "URL"}

func writeToFile(events []item) {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(header)
	for _, value := range events {
		array := []string{value.Type, value.Title, value.Place, value.ShortDesctription, value.URL}
		err = writer.Write(array)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
