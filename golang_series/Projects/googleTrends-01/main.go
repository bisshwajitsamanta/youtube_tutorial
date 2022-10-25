package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title    string `xml:"title"`
	Link     string `xml:"link"`
	Traffic  string `xml:"approx_traffic"`
	NewsItem []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadLineLink string `xml:"news_item_url"`
}

func GetGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func ReadGoogleTrends() []byte {
	resp := GetGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func main() {
	var r RSS
	data := ReadGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Below are the Google Search for today !!")
	fmt.Println(strings.Repeat("__", 30))
	for index, i := range r.Channel.ItemList {
		rank := (index + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term:", i.Title)
		fmt.Println("Link to the trend:", i.Link)
		for _, j := range i.NewsItem {
			fmt.Println("Headline:", j.Headline)
			fmt.Println("Link to the article:", j.HeadLineLink)
		}
		fmt.Println(strings.Repeat("__", 60))
	}
}
