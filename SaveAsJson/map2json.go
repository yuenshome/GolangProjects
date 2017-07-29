package main

import (
    "encoding/json"
    "fmt"
	"reflect"
)

type ChannelOperator struct {
	Oid      string
	Aid      string
	Name     string
	Isonline bool
	Msgchan  chan string  `json:"-"`
}

type ChannelOperators struct {
	Op []ChannelOperator
}

func main() {
	fmt.Println("Hello, 世界")
	mm := make(map[string]*ChannelOperator)
	mm["123"] = &ChannelOperator{Oid: "12312"}
	mm["5678"] = &ChannelOperator{Oid: "12312"}

	b, _ := json.Marshal(mm)
	fmt.Println(string(b))
	fmt.Println("_:", string(_))
	fmt.Println("b  type:", reflect.TypeOf(b))
	fmt.Println("mm type:", reflect.TypeOf(mm))
}


