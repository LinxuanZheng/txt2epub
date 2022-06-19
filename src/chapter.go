package src

import (
	"github.com/klarkxy/gohtml"
	"html"
)

const HTMLHeader = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`

type Chapter struct {
	FileName   string
	Title      string
	Paragraphs []string
}

func (c Chapter) GenChapterHTML() string {
	page := gohtml.Html().Attr("xmlns", "http://www.w3.org/1999/xhtml")
	head := page.Head()
	head.Title().Text(c.Title)
	head.Style().Type("text/css").Text(`p { text-indent: 2em; }
    	h1 {margin-top: 1em}
    	h2 {margin: 1em 0; text-align: center; font-size: 2em;}
    	h3 {margin: 0 0 2em; font-weight: normal; text-align:center; font-size: 1.5em; font-style: italic;}
    	.center { text-align: center; }`)
	content := page.Body().Attr("xml:lang", "zh-CN").Attr("lang", "zh-CN")

	content.H2().Text(c.Title)

	for _, para := range c.Paragraphs {
		content.P().Text(html.EscapeString("  " + para))
	}

	return HTMLHeader + page.String()
}
