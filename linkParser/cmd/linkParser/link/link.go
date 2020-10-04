package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link - represents link in html doc
type Link struct {
	Href string
	Text string
}

func getHref(n *html.Node) string {
	var href string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			href = attr.Val
		}
	}
	return href
}

func text(n *html.Node) string {
	var ret string
	if n.Type == html.TextNode {
		return n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func dfs(n *html.Node) []Link {
	var links []Link
	var curr *html.Node
	stack := []*html.Node{n}
	for len(stack) > 0 {
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
		for curr != nil {
			if curr.Type == html.ElementNode && curr.Data == "a" {
				links = append(links, Link{Href: getHref(curr), Text: text(curr)})
			}
			if curr.NextSibling != nil {
				stack = append(stack, curr.NextSibling)
			}
			curr = curr.FirstChild
		}
	}
	return links
}

// Parse - Returns slice of parsed links
func Parse(r io.Reader) ([]Link, error) {
	var links []Link
	doc, err := html.Parse(r)
	if err != nil {
		return links, err
	}
	return dfs(doc), nil
}
