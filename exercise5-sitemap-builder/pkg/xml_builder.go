package sitemap

import (
	"log"
	"net/http"
	"strings"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
)

type XmlStruct struct {
	Name string   `xml:"loc"`
	Map  XmlArray `xml:"url"`
}

type XmlArray []XmlStruct

type XmlContent struct {
	XmlName string   `xml:"urlset,attr"`
	Body    XmlArray `xml:"url"`
}

func (xc *XmlContent) Build(ha mypkg.HrefArray, name string, depth int) {
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
					searchForMoreLinks(name, h.Href, depth+1, finder),
				})
		}
	}
}

func searchForMoreLinks(link string, name string, depth int, finder map[string]struct{}) XmlArray {
	if depth > 5 {
		return nil
	}

	resp, err := http.Get(link+name)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(link+name)
		log.Printf("Get code: %v\n", resp.StatusCode)
		return nil
	}

	newContent := mypkg.NewParser()
	newContent.Parse(resp.Body)

	xs := &XmlStruct{}

	for _, h := range *newContent {
		if !strings.HasPrefix(h.Href, "/") {
			continue
		}

		if _, ok := finder[h.Href]; !ok {
			finder[h.Href] = struct{}{}
			if name == "/" {
				name = ""
			}

			xs.Map = append(xs.Map,
				XmlStruct{
					h.Href,
					searchForMoreLinks(link, name + strings.TrimRight(h.Href, "/"), depth+1, finder),
				})
		}
	}

	return xs.Map
}
