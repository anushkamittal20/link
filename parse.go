package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

//the struct contains a string part for href and a text part for text within
//the anchor tag without formatting
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

//deals with the href content in the node
func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

//deals with text part of the link node
func text(n *html.Node) string {
	// a simple text node can be returned without any addition
	if n.Type == html.TextNode {
		return n.Data
	}
	//if the node isnt an element node further parsing of it wont be required
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	//we look through the parse tree to find any relevant text that we can
	//return as text part of the node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")

}
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
