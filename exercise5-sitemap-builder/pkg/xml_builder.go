package sitemap

import (
	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
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

	for _, h := range ha {
		xc.Body = append(xc.Body,
			XmlStruct{
				h.Href,
				nil,
			})
	}
}
