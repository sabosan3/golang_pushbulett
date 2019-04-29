package main

import (
	"./lib"
	set "github.com/deckarep/golang-set"
)

func main() {
	conf := lib.LoadConfig()

	doc := lib.GetHTML(conf.Targetsite.URL)
	links := lib.GetLinkFromDoc(doc, conf.Targetsite.URL)

	setLinks := set.NewSet()
	for _, link := range links {
		setLinks.Add(link)
	}

	contents := []string{}
	var title string
	for _, link := range setLinks.ToSlice() {
		if lib.IsTodayNews(lib.GetHTML(link.(string))) {
			title = lib.GetTitleFromURL(link.(string))
			contents = append(contents, link.(string))
			lib.PostMessage(conf.Pushbulett.Token, title, contents...)
		}
	}
}
