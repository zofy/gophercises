package mapper

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/zofy/linkParser/cmd/linkParser/parser"
)

const XMLns = "http://www.sitemaps.org/schemas/sitemap/0.9"

// SiteMap -
type SiteMap map[string][]string

type empty struct{}

type loc struct {
	Val     []string `xml:"loc"`
	Address string   `xml:"address,attr"`
}

type urlSet struct {
	URLs  []loc  `xml:"url"`
	XMLns string `xml:"xmlns,attr"`
}

func readBody(url string) ([]byte, error) {
	var body []byte
	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

func getBaseURL(domain string) (string, error) {
	resp, err := http.Get(domain)
	if err != nil {
		return "", err
	}
	baseURL := &url.URL{
		Scheme: resp.Request.URL.Scheme,
		Host:   resp.Request.URL.Host,
	}
	return baseURL.String(), nil
}

func get(base, urlStr string) []string {
	var hrefs []string
	body, err := readBody(urlStr)
	if err != nil {
		return hrefs
	}
	links, err := parser.Parse(bytes.NewReader(body))
	if err != nil {
		return hrefs
	}
	for _, l := range links {
		if strings.HasPrefix(l.Href, "/") {
			hrefs = append(hrefs, base+l.Href)
		} else if strings.HasPrefix(l.Href, "http") {
			hrefs = append(hrefs, l.Href)
		}
	}
	return filter(hrefs, withBase(base))
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withBase(base string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, base)
	}
}

func bfs(baseURL string, maxDepth int) SiteMap {
	siteMap := make(SiteMap)
	seen := make(map[string]empty)
	var q map[string]empty
	nq := map[string]empty{
		baseURL: empty{},
	}
	var currURL string
	var depth int
	for depth <= maxDepth {
		q, nq = nq, make(map[string]empty)
		if len(q) == 0 {
			break
		}
		for currURL = range q {
			seen[currURL] = empty{}
			pages := get(baseURL, currURL)
			siteMap[currURL] = append(siteMap[currURL], pages...)
			for _, href := range pages {
				if _, ok := seen[href]; !ok {
					nq[href] = empty{}
				}
			}
		}
		depth++
	}
	return siteMap
}

// BuildMap -
func BuildMap(urlStr string, maxDepth int) (SiteMap, error) {
	baseURL, err := getBaseURL(urlStr)
	if err != nil {
		return SiteMap{}, err
	}
	return bfs(baseURL, maxDepth), nil
}

// ToXML -
func (sm SiteMap) ToXML() error {
	urlset := urlSet{XMLns: XMLns}
	for page, chs := range sm {
		urlset.URLs = append(urlset.URLs, loc{Address: page, Val: chs})
	}
	enc := xml.NewEncoder(os.Stdout)
	fmt.Print(xml.Header)
	enc.Indent("", "  ")
	if err := enc.Encode(urlset); err != nil {
		return err
	}
	return nil
}

// Print - prints sitemap
func (sm SiteMap) Print() {
	fmt.Println("Site map")
	for k, v := range sm {
		fmt.Println(k)
		for _, x := range v {
			fmt.Printf("\t %s\n", x)
		}
	}
}
