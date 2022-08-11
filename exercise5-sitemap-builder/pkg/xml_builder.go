package sitemap

import (
	"strings"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
)

var (
	depth int
)

type XmlStruct struct {
	Name string      `xml:"loc"`
	Map  []XmlStruct `xml:"url"`
}

type XmlContent struct {
	XmlName string      `xml:"urlset,attr"`
	Body    []XmlStruct `xml:"url"`
}

func (xc *XmlContent) Build(ha mypkg.HrefArray, name string) {
	xc.XmlName = name
	finder := make(map[string]struct{})

	for _, h := range ha {
		if !strings.HasPrefix(h.Href, "/") {
			continue
		}

		if _, ok := finder[h.Href]; !ok {
			finder[h.Href] = struct{}{}
			xc.Body = append(xc.Body,
				XmlStruct{
					h.Href,
					searchForMoreLinks(name + h.Href, 0),
				})
		}
	}
}

func searchForMoreLinks(link string, depth int) []XmlStruct {

	return nil
}
