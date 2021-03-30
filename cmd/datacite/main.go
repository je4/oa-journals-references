package main

import (
	"fmt"
	"github.com/je4/oa-journals-references/pkg/oai"
)

func main() {
	(&oai.Request{
		BaseURL: "https://oai.datacite.org/oai",
	}).HarvestSets(func(record *oai.Record) {
		fmt.Printf("%s\n\n", record.Metadata.Body[0:500])
	})
}
