package main

import (
	'fmt'
	'encoding/json'
)

func main{
	s := '{"key1": 1, "key2": 2}'
	b, _ := json.Marshal(s)
	fmt.Println("b:", string(b))
}
