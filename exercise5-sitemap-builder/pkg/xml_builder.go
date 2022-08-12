package sitemap

import (
	"log"
	"net/http"
	"strings"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
)

var (
	Maxdepth int
	Website  string
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

func (xc *XmlContent) Build(ha mypkg.HrefArray, name string, depth, maxdepth int) {
	xc.XmlName = name
	finder := make(map[string]struct{})

	for _, h := range ha {
		data := strings.TrimPrefix(h.Href, name)
		if !strings.HasPrefix(data, "/") {
			continue
		}

		if _, ok := finder[data]; !ok {
			finder[data] = struct{}{}
			xc.Body = append(xc.Body,
				XmlStruct{
					data,
					searchForMoreLinks(name, data, finder, depth+1, maxdepth),
				})
		}
	}
}

func searchForMoreLinks(link string, name string, finder map[string]struct{}, depth, maxdepth int) XmlArray {
	if depth > maxdepth {
		return nil
	}

	resp, err := http.Get(link + name)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(link + name)
		log.Printf("Get code: %v\n", resp.StatusCode)
		return nil
	}

	newContent := mypkg.NewParser()
	newContent.Parse(resp.Body)

	xs := &XmlStruct{}

	for _, h := range *newContent {
		data := strings.TrimPrefix(h.Href, name)
		if !strings.HasPrefix(data, "/") {
			continue
		}

		if _, ok := finder[data]; !ok {
			finder[data] = struct{}{}
			if name == "/" {
				name = ""
			}

			xs.Map = append(xs.Map,
				XmlStruct{
					data,
					searchForMoreLinks(link, data, finder, depth+1, maxdepth),
				})
		}
	}

	return xs.Map
}
