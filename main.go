package main

import (
	"fmt"
	constants "genTrade/helpers"
	finnIntegration "genTrade/integrations"
	

	"github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func main() {
	
	finn := finnIntegration.NewFinnhub()
	doLookup(*finn)
	getFinancials(*finn)
}

func doLookup(finn finnIntegration.Finnhub) []finnhub.SymbolLookup{
	resultList := finn.ListLookup(constants.SYMBOL_LIST[:])
	
	for _, item := range resultList {
		for _, symbolLookupInfo := range item.GetResult() {
			fmt.Println("Display Symbol:", symbolLookupInfo.GetDisplaySymbol())
			fmt.Println("Description: ", symbolLookupInfo.GetDescription())
			fmt.Println()
		}
		fmt.Println("----------------------")
	}

	return resultList
}

func getFinancials(finn finnIntegration.Finnhub) []finnhub.BasicFinancials {
	resultList := finn.ListBasicFinancials(constants.SYMBOL_LIST[:])
	
	for _, item := range resultList {
		fmt.Println("Display Symbol:", item.GetSymbol())
		data, _ := item.MarshalJSON()
		fmt.Println("Data: ", string(data))
		fmt.Println()
		
		fmt.Println("----------------------")
	}

	return resultList
}


