package main

import (
	"fmt"
	"reflect"
	"encoding/json"
)

type Student struct {
    Name    string
    Age     int
    Guake   bool
    Classes []string
    Price   float32
}

func main() {

	 st := &Student {
	     "Xiao Ming",
	     16,
	     true,
	     []string{"Math", "English", "Chinese"},
	      9.99,
	 }
	 fmt.Println("type: ", reflect.TypeOf(st))

	 b, _ := json.Marshal(st)

	 fmt.Println("b: ", string(b))

 }
