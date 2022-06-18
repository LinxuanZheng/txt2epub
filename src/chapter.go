package src

import (
	"github.com/klarkxy/gohtml"
)

const HTMLHeader = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`

type Chapter struct {
	FileName   string
	Title      string
	Paragraphs []string
}

func (c Chapter) GenChapterHTML() string {
	html := gohtml.Html().Attr("xmlns", "http://www.w3.org/1999/xhtml")
	html.Head().Title()
	content := html.Body().Attr("xml:lang", "zh-CN").Attr("lang", "zh-CN")

	content.H2().Text(c.Title)

	content.P().Text("&nbsp;")

	for _, para := range c.Paragraphs {
		content.P().Text(para)
	}

	return HTMLHeader + html.String()
}
