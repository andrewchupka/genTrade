package structs

import (
)

type LiveTrade struct {
	Data []Trade `json:"data"`
	Type string `json:"type"`
}