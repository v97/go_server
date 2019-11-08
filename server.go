package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving on port 8000")
}

func page_1_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page 1")
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

/*
<sitemap><loc>https://www.scmp.com/sitemap_article.xml</loc><lastmod>2019-11-09T04:00Z</lastmod></sitemap>
<sitemap><loc>https://www.scmp.com/sitemap_gallery.xml</loc><lastmod>2019-11-09T04:00Z</lastmod></sitemap>
<sitemap><loc>https://www.scmp.com/sitemap_video.xml</loc><lastmod>2019-11-09T04:00Z</lastmod></sitemap>
*/

func main() {
	//resp, _ := http.Get("https://rss.nytimes.com/services/xml/rss/nyt/Technology.xml/")
	resp, _ := http.Get("https://www.scmp.com/sitemap.xml")
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println("Starting server")
	fmt.Println(s.Locations)

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/page_1/", page_1_handler)

	//http.ListenAndServe(":8000", nil)
}