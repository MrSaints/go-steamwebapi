package steamwebapi

import (
//"testing"
)

var (
	client *Client
)

func setup() {
	if client == nil {
		client = NewClient("")
	}
}
