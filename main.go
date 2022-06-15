package main

import (
	"fmt"
	constants "genTrade/helpers"
	finnIntegration "genTrade/integrations"
	

	// "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func main() {
	
	finn := finnIntegration.NewFinnhub()
	resultList := finn.ListLookup(constants.SYMBOL_LIST[:])
	fmt.Println("done")
	
	for val := range resultList {
		fmt.Println(val)
	}

	
}
