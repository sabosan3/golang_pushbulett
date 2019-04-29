package lib

import (
	"math"
	"net/url"
	"os"
	"strings"
	"time"

	goquery "github.com/PuerkitoBio/goquery"
)

// TimeFormat is const
const TimeFormat = "2006年1月2日"

/*
GetHTML return *goquery.Document
*/
func GetHTML(url string) *goquery.Document {

	logger := InitLogging()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		logger.Printf("error: failed to get document from " + url)
		logger.Printf(err.Error())
		os.Exit(1)
	}

	return doc
}

/*
GetLinkFromDoc return []string
*/
func GetLinkFromDoc(doc *goquery.Document, burl string) []string {

	baseurl, _ := url.Parse(burl)

	links := []string{}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		fullLink := toAbsURL(baseurl, link)

		if strings.Contains(fullLink, "articles") {
			links = append(links, fullLink)
		}

	})

	return links
}

/*
GetTitleFromURL return string
*/
func GetTitleFromURL(url string) string {

	doc := GetHTML(url)
	title := doc.Find("#sb-site > div > div > div.col.col-sm-8.content-left > div > div.article-header > h1")

	return title.Text()
}

/*
toAbsURL return string
*/
func toAbsURL(baseurl *url.URL, weburl string) string {
	relurl, err := url.Parse(weburl)
	if err != nil {
		return ""
	}
	absurl := baseurl.ResolveReference(relurl)
	return absurl.String()
}

/*
IsTodayNews return bool
*/
func IsTodayNews(doc *goquery.Document) bool {

	date := doc.Find("#sb-site > div > div > div.col.col-sm-8.content-left > div > div.article-header > div.date")
	strDate := date.Text()
	strDate = strings.Split(strDate, " ")[0]

	timeDate, _ := time.Parse(TimeFormat, strDate)
	//timeDate, _ := time.Parse(TimeFormat, "2019年4月29日")

	now := time.Now()
	//now := time.Date(2019, 4, 2, 0, 0, 0, 0, time.Local)

	sub := timeDate.Sub(now)

	return math.Abs(sub.Hours()) <= 24
}
