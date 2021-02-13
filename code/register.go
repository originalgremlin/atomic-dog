package main

import (
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
		return fmt.Errorf("failed to register %s: only callable functions may be registered", name)
	}
	callables[name] = callable
	return nil
}

// Unregister remote functions.
func Unregister(name string) error {
	delete(callables, name)
	return nil
}
