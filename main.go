package main

import (
	finnhub "genTrade/integrations";
)

func main() {
	finn := finnhub.NewFinnhub() 

	finn.Lookup("APPL")
	finn.Lookup("DKS")
	finn.Lookup("BTC")
	finn.Lookup("HSGD")
}