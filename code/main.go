package main

import (
	"encoding/json"
	"fmt"

	zmq "github.com/go-zeromq/zmq4"
)

func main() {
	m := make(map[string]interface{})
	m["one"] = 1
	m["two"] = "two"
	m["three"] = []int{1, 2, 3}
	m["four"] = 4.4
	m["five"] = true
	m["six"] = nil
	j, _ := json.Marshal(m)

	m2 := make(map[string]interface{})
	json.Unmarshal(j, &m2)

	fmt.Println(Arguments)
	fmt.Println(Conf)
	fmt.Println(Conf.Exists("foo"))
	fmt.Println(Conf.Get("foo"))
	fmt.Println(zmq.CmdCancel)
	fmt.Println(m)
	fmt.Println(j)
	fmt.Println(m2)
}
