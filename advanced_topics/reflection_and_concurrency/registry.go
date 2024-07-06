package main

import (
	"fmt"
	"reflect"
	"sync"
)

var registry = make(map[string]reflect.Value)
var mutex sync.RWMutex

func RegisterService(name string, svc interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	registry[name] = reflect.ValueOf(svc) // this returns a Value representing the run-time data of svc and since it in of type interface therefore it is a pointer to the interface
}

func CallService(name string, methodName string, args []reflect.Value) (reflect.Value, error) {
	mutex.RLock()
	svc, ok := registry[name]
	mutex.RUnlock()

	if !ok {
		return reflect.Value{}, fmt.Errorf("service not found: %s", name)
	}

	method := svc.MethodByName(methodName)
	if !method.IsValid() {
		return reflect.Value{}, fmt.Errorf("method not found: %s", methodName)
	}
	response := method.Call(args)
	if len(response) > 1 && !response[1].IsNil() {
		return reflect.Value{}, response[1].Interface().(error)
	}
	return response[0], nil
}
