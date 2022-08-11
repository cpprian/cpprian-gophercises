package test_parser

import (
	"os"
	"testing"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
)

const (
	ex1 string = "./ex1.html"
	ex2 string = "./ex2.html"
	ex3 string = "./ex3.html"
	ex4 string = "./ex4.html"
)

func TestParser(t *testing.T) {

	t.Run("basic scenario a tagname", func(t *testing.T) {
		f, err := os.Open(ex1)
		IsErrorOccured(t, err)
		defer f.Close()

		href := mypkg.NewParser()
		href.Parse(f)

		want := &mypkg.HrefStruct{
			Href: "/other-page",
			Text: "A link to another page",
		}
		got := (*href)[0]

		AssertStringComparison(t, want.Href, got.Href)
		AssertStringComparison(t, want.Text, got.Text)
	})

	t.Run("page with extra space with extra tag", func(t *testing.T) {
		f, err := os.Open(ex2)
		IsErrorOccured(t, err)
		defer f.Close()

		got := mypkg.NewParser()
		got.Parse(f)

		want := &mypkg.HrefArray{
			{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github Hello world!",
			},
		}

		CompareHrefArray(t, *want, *got)
	})

	t.Run("page with empty links and extra attributes", func(t *testing.T) {
		f, err := os.Open(ex3)
		IsErrorOccured(t, err)
		defer f.Close()

		got := mypkg.NewParser()
		got.Parse(f)

		want := &mypkg.HrefArray{
			{
				Href: "/lost",
				Text: "Lost? Need help?",
			},
			{
				Href: "https://twitter.com/marcusolsson",
				Text: "@marcusolsson",
			},
		}

		CompareHrefArray(t, *want, *got)
	})

	t.Run("data from tag with comment", func(t *testing.T) {
		f, err := os.Open(ex4)
		IsErrorOccured(t, err)
		defer f.Close()

		got := mypkg.NewParser()
		got.Parse(f)

		want := &mypkg.HrefArray{
			{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		}

		CompareHrefArray(t, *want, *got)
	})
}

func CompareHrefArray(t testing.TB, want, got mypkg.HrefArray) {
	if len(want) != len(got) {
		t.Errorf("expected len %v, but actually is %v", len(want), len(got))
	}

	for i := 0; i < len(want); i++ {
		AssertStringComparison(t, want[i].Href, got[i].Href)
		AssertStringComparison(t, want[i].Text, got[i].Text)
	}
}

func AssertStringComparison(t testing.TB, want, got string) {
	if want != got {
		t.Errorf("want %v, got %v\n", want, got)
	}
}

func IsErrorOccured(t testing.TB, err error) {
	if err != nil {
		t.Error(err)
	}
}
