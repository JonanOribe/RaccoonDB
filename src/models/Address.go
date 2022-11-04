package models

import (
	"encoding/json"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}
