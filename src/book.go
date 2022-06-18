package src

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Book struct {
	File     *os.File
	Config   Config
	Chapters []Chapter
}

func NewBook(file *os.File) *Book {
	return &Book{
		File:     file,
		Config:   LoadConfig(),
		Chapters: nil,
	}
}

func (book *Book) LoadBook() error {
	br := bufio.NewReader(book.File)
	reg := regexp.MustCompile(book.Config.Regex)
	chapter := Chapter{
		FileName:   "text00000.html",
		Title:      "前言",
		Paragraphs: []string{"\r\n", "\r\n", "\r\n", "\r\n"},
	}
	book.Chapters = make([]Chapter, 0)

	cnt := 0

	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}

		if reg.Match([]byte(line)) {
			book.Chapters = append(book.Chapters, chapter)
			cnt++

			chapter = Chapter{
				FileName:   "text" + fmt.Sprintf("%05d", cnt) + ".html",
				Title:      strings.Replace(line, "\r\n", "", -1),
				Paragraphs: []string{"\r\n", "\r\n", "\r\n", "\r\n"},
			}
		} else {
			line = strings.Replace(line, "\r\n", "", -1)
			line = strings.Replace(line, " ", "", -1)
			if len(line) > 0 {
				chapter.Paragraphs = append(chapter.Paragraphs, line)
			}
		}
	}

	return nil
}

func (book *Book) GenBook() {

}
