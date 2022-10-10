package main

import "fmt"

type Client struct {
	id   string // variable start with small case is private
	name string
}

func NewClient(id string, name string) *Client {
	return &Client{
		id:   id,
		name: name,
	}
}

func main() {
	client := NewClient("22", "RRRR")
	fmt.Println(client)
}
