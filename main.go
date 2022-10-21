package main

import (
	"encoding/json"
	"fmt"
	constants "genTrade/helpers"
	finnIntegration "genTrade/integrations"

	"github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func main() {
	
	finn := finnIntegration.NewFinnhub()
	// doLookup(*finn)
	// getFinancials(*finn)

	// hard coded to getting bitcoin
	finishedWriting := make (chan string, len(constants.CRYPTO_LIST))
	for _, cryptoSymbol := range(constants.CRYPTO_LIST) {
		go finn.TradeLookup(cryptoSymbol, finishedWriting)
	}

	for _, item := range constants.CRYPTO_LIST{
		<-finishedWriting
		fmt.Println("Finished writing", item)
	}
	
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
		data, _ := json.Marshal(item)
		fmt.Println("Data: ", string(data))
		fmt.Println()
		
		fmt.Println("----------------------")
	}

	return resultList
}


