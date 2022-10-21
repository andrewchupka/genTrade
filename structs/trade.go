package structs

import (
	"time"
	"encoding/json"
	"log"
)

type Trade struct {
	Price float64 `json:"p"`
	Timestamp time.Time `json:"t"`
	Symbol string `json:"s"`
	Volume float64 `json:"v"`
	additionalData interface{} `json:"c"`
}

func (trade *Trade) Stringify() string {
	tradeJson, err := json.Marshal(trade)
	if (err != nil) {
		log.Fatal(err.Error())
	} 
	
	return string(tradeJson)
}

func (trade *Trade) UnmarshalJSON(b []byte) error {
	m := struct {
		Price float64 `json:"p"`
		Timestamp int64 `json:"t"`
		Symbol string `json:"s"`
		Volume float64 `json:"v"`
		additionalData interface{} `json:"c"`
	}{}

	err := json.Unmarshal(b, &m)
	if (err != nil) { return err }

	trade.Price = m.Price
	trade.Timestamp = time.UnixMilli(m.Timestamp)
	trade.Symbol = m.Symbol
	trade.Volume = m.Volume
	trade.additionalData = m.additionalData

	return nil
}

