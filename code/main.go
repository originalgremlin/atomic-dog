package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// TODO: what type is the value here?  I want "any function of any arguments and return type".  Is just "interface{}" okay?
var callables = make(map[string]interface{})

// Register functions that can be executed remotely.
//
// Names must be unique within the registering program, but need not be globally unique.
func Register(name string, callable interface{}) error {
	if reflect.TypeOf(callable).Kind() != reflect.Func {
		return fmt.Errorf("Failed to register %s. Only callable functions may be registered.", name)
	} else {
		callables[name] = callable
		return nil
	}
}

// Unregister remote functions.
func Unregister(name string) error {
	delete(callables, name)
	return nil
}

// Run this node as an independent worker in the distributed program.
func Run() error {
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
	j, _ := json.Marshal(m)

	m2 := make(map[string]interface{})
	json.Unmarshal(j, &m2)

	fmt.Println(Arguments)
	fmt.Println(Conf)
	fmt.Println(Conf.Exists("foo"))
	fmt.Println(Conf.Get("foo"))
	fmt.Println(m)
	fmt.Println(j)
	fmt.Println(m2)
}
