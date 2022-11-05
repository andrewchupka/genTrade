package main

import (
	"encoding/json"
	"fmt"
	"genTrade/helpers"
	finnIntegration "genTrade/integrations"
	"genTrade/structs"
	"github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func main() {
	
	finn := finnIntegration.NewFinnhub()
	// doLookup(*finn)
	// getFinancials(*finn)

	// hard coded to getting bitcoin
	finishedWriting := make (chan string, len(helpers.CRYPTO_LIST))
	for _, cryptoSymbol := range(helpers.CRYPTO_LIST) {
		pool := structs.MakeGenePool(cryptoSymbol)
		// go pool.Process()
		go finn.LiveTradeFeed(cryptoSymbol, finishedWriting, pool.TradeChannel)

	}

	for _, item := range helpers.CRYPTO_LIST{
		<-finishedWriting
		fmt.Println("Finished writing", item)
	}

}

func doLookup(finn finnIntegration.Finnhub) []finnhub.SymbolLookup {
	resultList := finn.ListLookup(helpers.SYMBOL_LIST[:])

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
	resultList := finn.ListBasicFinancials(helpers.SYMBOL_LIST[:])

	for _, item := range resultList {
		fmt.Println("Display Symbol:", item.GetSymbol())
		data, _ := json.Marshal(item)
		fmt.Println("Data: ", string(data))
		fmt.Println()

		fmt.Println("----------------------")
	}

	return resultList
}
