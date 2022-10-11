package structs

import (
)

type LiveTrade struct {
	Data []Trade
	DataType string `json:type`
}