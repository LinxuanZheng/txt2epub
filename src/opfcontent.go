package src

import (
	"encoding/xml"
	"strconv"
)

type Opfcontent struct {
	XMLName          xml.Name `xml:"package"`
	Version          string   `xml:"version,attr"`
	XMLNS            string   `xml:"xmlns,attr"`
	UniqueIdentifier string   `xml:"unique-identifier,attr"`

	// metadata
	MetaData MetaData `xml:"metadata"`

	// manifest
	Manifest Manifest `xml:"manifest"`

	// spine
	Spine Spine `xml:"spine"`

	// guide
	//Guide Guide `xml:"guide"`
}

type MetaData struct {
	XMLName   xml.Name `xml:"metadata"`
	XMLNS_DC  string   `xml:"xmlns:dc,attr"`
	XMLNS_OPF string   `xml:"xmlns:opf,attr"`
	XMLNS_XSI string   `xml:"xmlns:xsi,attr"`

	// dc:title
	DC_title string `xml:"dc:title"`

	// dc:creator
	DC_creator string `xml:"dc:creator"`

	// dc:description
	DC_description string `xml:"dc:description"`

	// dc:language
	DC_language string `xml:"dc:language"`

	// dc:identifier
	DC_identifier DC_identifier `xml:"dc:identifier"`

	// meta
	Meta Meta `xml:"meta"`
}

type DC_identifier struct {
	XMLName   xml.Name `xml:"dc:identifier"`
	ID        string   `xml:"id,attr"`
	OpfScheme string   `xml:"opf:scheme,attr"`
	Value     string   `xml:",innerxml"`
}

type Meta struct {
	XMLName xml.Name `xml:"meta"`
	Name    string   `xml:"name,attr"`
	Content string   `xml:"content,attr"`
}

type Manifest struct {
	XMLName xml.Name `xml:"manifest"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName   xml.Name `xml:"item"`
	Href      string   `xml:"href,attr"`
	Id        string   `xml:"id,attr"`
	MediaType string   `xml:"media-type,attr"`
}

type Spine struct {
	XMLName  xml.Name  `xml:"spine"`
	Toc      string    `xml:"toc,attr"`
	ItemRefs []ItemRef `xml:"itemref"`
}

type ItemRef struct {
	XMLName xml.Name `xml:"itemref"`
	ItemRef string   `xml:"idref,attr"`
}

type Guide struct {
	XMLName   xml.Name    `xml:"guide"`
	Reference []Reference `xml:"reference"`
}

type Reference struct {
	XMLName xml.Name `xml:"reference"`
	Href    string   `xml:"href,attr"`
	Title   string   `xml:"title,attr"`
	Type    string   `xml:"type,attr"`
}

func NewDefaultOpfContent() Opfcontent {
	return Opfcontent{
		Version:          "2.0",
		XMLNS:            "http://www.idpf.org/2007/opf",
		UniqueIdentifier: "uuid_id",
		MetaData: MetaData{
			XMLNS_DC:    "http://purl.org/dc/elements/1.1/",
			XMLNS_OPF:   "http://www.idpf.org/2007/opf",
			XMLNS_XSI:   "http://www.w3.org/2001/XMLSchema-instance",
			DC_language: "zh",
			DC_identifier: DC_identifier{
				ID:        "uuid_id",
				OpfScheme: "uuid",
			},
			Meta: Meta{
				Name: "cover",
			},
		},
		Manifest: Manifest{},
		Spine: Spine{
			Toc: "ncx",
		},
		//Guide: Guide{},
	}
}

func (o *Opfcontent) LoadManifestAndSpine(hrefs []string) {
	items := make([]Item, 0)
	itemRefs := make([]ItemRef, 0)
	cnt := 1
	for _, href := range hrefs {
		id := "id_" + strconv.Itoa(cnt)
		items = append(items, Item{
			Href:      "TEXT/" + href,
			Id:        id,
			MediaType: "application/xhtml+xml",
		})
		itemRefs = append(itemRefs, ItemRef{
			ItemRef: id,
		})
		cnt++
	}
	o.Manifest = Manifest{
		Items: items,
	}
	o.Spine = Spine{
		Toc:      "ncx",
		ItemRefs: itemRefs,
	}
}
