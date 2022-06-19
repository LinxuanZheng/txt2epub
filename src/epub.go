package src

import (
	"encoding/xml"
	"github.com/klarkxy/gohtml"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

const MIMETYPE = `application/epub+zip`
const META_CONTAINER = `<?xml version="1.0" encoding="UTF-8" ?>
<container version="1.0" xmlns="urn:oasis:names:tc:opendocument:xmlns:container">
  <rootfiles> <rootfile full-path="content.opf" media-type="application/oebps-package+xml"/> </rootfiles>
</container>`
const CoverSource = `data:image/png;base64,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`

type Epub struct {
	Book       *Book
	OutputPath string
}

func (e *Epub) GenMimetype() {
	if Exists(e.OutputPath + "/mimetype") {
		err := os.Remove(e.OutputPath + "/mimetype")
		if err != nil {
			panic(err)
		}
	}

	err := ioutil.WriteFile(e.OutputPath+"/mimetype", []byte(MIMETYPE), 0444)
	if err != nil {
		panic(err)
	}
}

func (e *Epub) GenContainerXML() {
	metaDir := e.OutputPath + "/META-INF"
	if _, err := os.Stat(metaDir); os.IsNotExist(err) {
		err := os.Mkdir(metaDir, 0777)
		if err != nil {
			panic(err)
		}
		err = os.Chmod(metaDir, 0777)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	if Exists(e.OutputPath + "/META-INF/container.xml") {
		err := os.Remove(e.OutputPath + "/META-INF/container.xml")
		if err != nil {
			panic(err)
		}
	}

	err := ioutil.WriteFile(e.OutputPath+"/META-INF/container.xml", []byte(META_CONTAINER), 0444)
	if err != nil {
		panic(err)
	}
}

func (e *Epub) GenOpfContent() {
	opf := NewDefaultOpfContent()

	opf.MetaData.DC_title = e.Book.Config.Title
	opf.MetaData.DC_creator = e.Book.Config.Author
	opf.MetaData.DC_description = e.Book.Config.Desc
	opf.MetaData.DC_identifier.Value = e.Book.Config.UUID
	opf.MetaData.Meta.Content = filepath.Base(e.Book.Config.Cover)

	hrefs := make([]string, 0)
	for _, chapter := range e.Book.Chapters {
		hrefs = append(hrefs, chapter.FileName)
	}
	opf.LoadManifestAndSpine(hrefs)

	out, err := xml.MarshalIndent(opf, "  ", "    ")
	if err != nil {
		panic(err)
	}
	output := []byte(xml.Header)
	output = append(output, out...)

	if Exists(e.OutputPath + "/content.opf") {
		err := os.Remove(e.OutputPath + "/content.opf")
		if err != nil {
			panic(err)
		}
	}

	err = ioutil.WriteFile(e.OutputPath+"/content.opf", output, 0444)
	if err != nil {
		panic(err)
	}
}

func (e *Epub) CopyCover() {
	if !Exists(e.Book.Config.Cover) {
		panic("封面不存在")
	}

	err := CopyFile(e.Book.Config.Cover, e.OutputPath+"/"+filepath.Base(e.Book.Config.Cover))
	if err != nil {
		panic(err)
	}
}

func (e *Epub) GenChapters() {
	if Exists(e.OutputPath + "/TEXT") {
		err := os.RemoveAll(e.OutputPath + "/TEXT")
		if err != nil {
			panic(err)
		}
	}

	err := os.Mkdir(e.OutputPath+"/TEXT", 0777)
	if err != nil {
		panic(err)
	}

	baseDir := e.OutputPath + "/TEXT/"
	for _, chapter := range e.Book.Chapters {
		if Exists(baseDir + chapter.FileName) {
			err := os.Remove(baseDir + chapter.FileName)
			if err != nil {
				panic(err)
			}
		}
		output := chapter.GenChapterHTML()
		err := ioutil.WriteFile(baseDir+chapter.FileName, []byte(output), 0444)
		if err != nil {
			panic(err)
		}
	}
}

func (e *Epub) GenNCX() {
	ncx := gohtml.Tag("ncx").
		Attr("xmlns", "http://www.daisy.org/z3986/2005/ncx/").
		Attr("version", "2005-1")

	head := ncx.Head()
	head.Tag("meta").Attr("name", "dtb:uid").Attr("content", e.Book.Config.UUID)
	head.Tag("meta").Attr("name", "dtb:depth").Attr("content", "1")
	head.Tag("meta").Attr("name", "dtb:totalPageCount").Attr("content", "0")
	head.Tag("meta").Attr("name", "dtb:maxPageNumber").Attr("content", "0")

	ncx.Tag("docTitle").Tag("text").Text(e.Book.Config.Title)

	navMap := ncx.Tag("navMap")

	cnt := 1
	for _, chapter := range e.Book.Chapters {
		navPoint := navMap.Tag("navPoint").
			Attr("id", "id_"+strconv.Itoa(cnt)).
			Attr("playOrder", strconv.Itoa(cnt)).
			Attr("class", "chapter")
		navPoint.Tag("navLabel").Tag("text").Text(chapter.Title)
		navPoint.Tag("content").Attr("src", "TEXT/"+chapter.FileName)
		cnt++
	}

	output := xml.Header + ncx.String()

	if Exists(e.OutputPath + "/toc.ncx") {
		err := os.Remove(e.OutputPath + "/toc.ncx")
		if err != nil {
			panic(err)
		}
	}

	err := ioutil.WriteFile(e.OutputPath+"/toc.ncx", []byte(output), 0444)
	if err != nil {
		panic(err)
	}
}

func (e *Epub) GenZip() {
	files, dirs, err := GetFilesAndDirs(e.OutputPath)
	if err != nil {
		panic(err)
	}

	files = append(files, dirs...)
	openFiles := make([]*os.File, 1)
	for _, file := range files {
		fi, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		if path.Base(fi.Name()) == "mimetype" {
			openFiles[0] = fi
		} else {
			openFiles = append(openFiles, fi)
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = ZipCompress(openFiles, wd+"/"+e.Book.Config.Title+".zip")
	if err != nil {
		panic(err)
	}

	for i, _ := range openFiles {
		openFiles[i].Close()
	}

	err = os.Rename(wd+"/"+e.Book.Config.Title+".zip", wd+"/"+e.Book.Config.Title+".epub")
	if err != nil {
		panic(err)
	}
}
