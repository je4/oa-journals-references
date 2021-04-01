package main

import (
	"encoding/xml"
	"fmt"
	"github.com/je4/oa-journals-references/v2/pkg/oai"
	"log"
	"time"
)

type DOAJ_Journal struct {
	XMLName     xml.Name `xml:"oai_dc dc"`
	Title       []string `xml:"dc title"`
	Creator     []string `xml:"dc creator"`
	Subject     []string `xml:"dc subject"`
	Description []string `xml:"dc description"`
	Publisher   []string `xml:"dc publisher"`
	Contributor []string `xml:"dc contributor"`
	Date        []string `xml:"dc date"`
	Type        []string `xml:"dc type"`
	Format      []string `xml:"dc format"`
	Identifier  []string `xml:"dc identifier"`
	Source      []string `xml:"dc source"`
	Language    []string `xml:"dc language"`
	Relation    []string `xml:"dc relation"`
	Coverage    []string `xml:"dc coverage"`
	Rights      []string `xml:"dc rights"`
}

type DOAJ_Article_Subject struct {
	Type string `xml:"xsi type,attr"`
	Data string `xml:",chardata"`
}

type DOAJ_Article struct {
	XMLName     xml.Name               `xml:"oai_dc dc"`
	Title       []string               `xml:"dc title"`
	Identifier  []string               `xml:"dc identifier"`
	Date        []time.Time            `xml:"dc date"`
	Relation    []string               `xml:"dc relation"`
	Description []string               `xml:"dc description"`
	Creator     []string               `xml:"dc creator"`
	Publisher   []string               `xml:"dc publisher"`
	Type        []string               `xml:"dc type"`
	Subject     []DOAJ_Article_Subject `xml:"dc subject"`
	Language    []string               `xml:"dc language"`
	Source      []string               `xml:"dc source"`
}

func main() {
	/*
		(&oai.Request{
			BaseURL: "https://oai.datacite.org/oai",
		}).HarvestSets(func(set *oai.Set) {
			fmt.Printf("[%s] %s: %s\n", set.SetSpec, set.SetName, set.SetDescription.GoString())
		})
	*/

	/*
		counter := 0
		(&oai.Request{
			BaseURL: "https://doaj.org/oai",
			Set: "TENDOkFydHMgaW4gZ2VuZXJhbA~~",
			MetadataPrefix: "oai_dc",
		}).HarvestRecords(func(record *oai.Record) {
			counter++

			journal := &DOAJ_Journal{} // OAI_dc_base
			if err := xml.Unmarshal(record.Metadata.Body, journal); err != nil {
				log.Fatalf("cannot unmarshal %s: %v", record.Metadata.GoString(), err)
			}

			fmt.Printf("%6d [%s] %s\n", counter, record.Header.Identifier, journal.Identifier)
		})
	*/

	counter := 0
	(&oai.Request{
		BaseURL:        "https://doaj.org/oai.article",
		Set:            "TENDOkFydHMgaW4gZ2VuZXJhbA~~",
		MetadataPrefix: "oai_dc",
	}).HarvestRecords(func(record *oai.Record) {
		counter++

		article := &DOAJ_Article{} // OAI_dc_base
		if err := xml.Unmarshal(record.Metadata.Body, article); err != nil {
			log.Fatalf("cannot unmarshal %s: %v", record.Metadata.GoString(), err)
		}

		fmt.Printf("%6d [%s]\n", counter, record.Header.Identifier)
	})

}
