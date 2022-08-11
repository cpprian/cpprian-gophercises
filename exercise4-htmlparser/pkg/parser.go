package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type HrefStruct struct {
	Href string `xml:"loc"`
	Text string `xml:"tag"`
}

type HrefArray []HrefStruct

func NewParser() *HrefArray {
	return &HrefArray{}
}

func (ha *HrefArray) Parse(r io.Reader) error {
	z, err := html.Parse(r)
	if err != nil {
		return err
	}

	ha.parseElementNode(z)
	return nil
}

func (ha *HrefArray) parseElementNode(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && a.Val != "#" {
				*ha = append(*ha, HrefStruct{
					Href: a.Val,
					Text: getContentFromTag(n),
				})
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ha.parseElementNode(c)
	}
}

func getContentFromTag(n *html.Node) string {
	result := ""

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			if c.FirstChild != nil {
				result += " " + getContentFromTag(c)
			}
		} else if c.Type == html.TextNode {
			result += strings.TrimSpace(c.Data)
		}
	}
	return result
}
