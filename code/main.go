package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

// Start this node as an independent worker in the distributed program.
func Start() error {
	// TODO: launch all the sockets and kick off the execution cycle
	return nil
}

func main() {
	m := make(map[string]interface{})
	m["one"] = 1
	m["two"] = "two"
	m["three"] = []int{1, 2, 3}
	m["four"] = 4.4
	m["five"] = true
	m["six"] = nil
	j, _ := msgpack.Marshal(m)

	m2 := make(map[string]interface{})
	msgpack.Unmarshal(j, &m2)

	fmt.Println(Arguments)
	fmt.Println(Conf)
	fmt.Println(Conf.Exists("foo"))
	fmt.Println(Conf.Get("foo"))
	fmt.Println(m)
	fmt.Println(j)
	fmt.Println(m2)
}
