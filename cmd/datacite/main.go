package main

import (
	"fmt"
	"github.com/je4/oa-journals-references/v2/pkg/oai"
)

func main() {
	/*
		(&oai.Request{
			BaseURL: "https://oai.datacite.org/oai",
		}).HarvestSets(func(set *oai.Set) {
			fmt.Printf("[%s] %s: %s\n", set.SetSpec, set.SetName, set.SetDescription.GoString())
		})
	*/

	counter := 0
	(&oai.Request{
		BaseURL:        "https://oai.datacite.org/oai",
		Set:            "ETHZ.FHNW",
		MetadataPrefix: "oai_dc",
	}).HarvestRecords(func(record *oai.Record) {
		counter++
		fmt.Printf("%6d [%s] %s\n", counter, record.Header.Identifier, record.Metadata.GoString())
	})

}
