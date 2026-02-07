package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	//PAYWALL
	//search("www.ft.com", "https://www.ft.com/markets", ".js-teaser-heading-link", "href", "", , false)

	//Everything Else
	//search("news.ycombinator.com", "https://news.ycombinator.com/news", ".titleline > a", "href", "", false)
	//search("www.propublica.org", "https://www.propublica.org/", ".story-card__hed > a", "href", "", false)
	//search("apnews.com", "https://apnews.com/", ".PagePromo-title > a", "href", "", false)
	//search("arstechnica.com", "https://arstechnica.com/", "article > div > div > h2 > a", "href", "", false)
	//search("www.mlive.com" ,"https://www.mlive.com/news/", ".river-item > div > div > div > h2 > a", "href", "", false)
	//search("techcrunch.com", "https://techcrunch.com/latest/", ".loop-card__title > a", "href", "", false)
	//search("www.fool.com", "https://www.fool.com/trending-news/", "article > div > a", "href", "h5", true)
	//search("www.kenklippenstein.com", "https://www.kenklippenstein.com/archive?sort=new", ".container-Qnseki > div > div > a", "href", "", false)
}

func search(domain string, url string, searchParam string, matchAttr string, titleOverride string, dynamicLinking bool) {
	// Initialize the Collector
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML(searchParam, func(e *colly.HTMLElement) {
		var text string
		var link string

		if titleOverride != "" {
			text = strings.TrimSpace(e.ChildText(titleOverride))
		} else {
			text = strings.TrimSpace(e.Text)
		}

		if dynamicLinking {
			link = domain + e.Attr(matchAttr)
		} else {
			link = e.Attr(matchAttr)
		}

		fmt.Printf("Article: %s\n\tLink: %s\n", text, link)
	})

	// Request Logging
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Error Handling
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Start the crawl
	c.Visit(url)
}

/*
	// 1. Initialize the Collector
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// 2. OnHTML: Triggered when the collector finds a specific HTML element (CSS Selector)
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("Main Heading found:", e.Text)
	})

	// 3. Find and visit all links on the page
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %s\n", link)

		// To follow links recursively, you'd use:
		// e.Request.Visit(link)
	})

	// 4. Request Logging
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// 5. Error Handling
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Start the crawl
	c.Visit("https://en.wikipedia.org/wiki/Go_(programming_language)")

*/
