package parser

import (
	"fmt"
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

func (ha *HrefArray) String() string {
	var result string

	for _, h := range *ha {
		result += fmt.Sprintf("Href: %v, Text: %v\n", h.Href, h.Text)
	}
	return result
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
	var rec func(*html.Node)
	rec = func(n *html.Node) {
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

		if n.FirstChild != nil {
			rec(n.FirstChild)
		}
		if n.NextSibling != nil {
			rec(n.NextSibling)
		}
	}
	rec(n)
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
