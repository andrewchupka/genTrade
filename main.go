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
	
	for _, item := range resultList {
		for _, symbolLookupInfo := range item.GetResult() {
			fmt.Println("Display Symbol:", symbolLookupInfo.GetDisplaySymbol())
			fmt.Println("Description: ", symbolLookupInfo.GetDescription())
			fmt.Println()
		}
		fmt.Println("----------------------")
	}

	
}
