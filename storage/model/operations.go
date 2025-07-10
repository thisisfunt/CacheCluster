package model

import (
	"encoding/json"
	"errors"
)

var storage = make(map[string]interface{})

func GetObject(params map[string]string) (interface{}, error) {
	return storage[params["key"]], nil
}

func SetObject(params map[string]string) (interface{}, error) {
	key := params["key"]
	value := params["value"]
	structName := params["struct"]
	factory := Structs[structName]()
	if factory == nil {
		return nil, errors.New("Struct don't exist")
	}
	err := json.Unmarshal([]byte(value), &factory)
	if err != nil {
		return nil, errors.New("Not correct json")
	}
	storage[key] = factory
	return factory, nil
}

func DeleteObject(params map[string]string) (interface{}, error) {
	delete(storage, params["key"])
	return true, nil
}

func ExistsObjectByKey(params map[string]string) (interface{}, error) {
	_, isExist := storage[params["key"]]
	return isExist, nil
}

// CREATE YOU FUNCTIONS HERE!
// DATA CONTAINS IN "storage"
// GET PARAMETERS CONTAIN in "params"
// YOU CAN USE DEFAULT FUNCTIONS AS AN EXAMPLE

// REGISTER YOU FUNCTIONS IN LIST
// YOU CAN USE DEFAULT FUNCTIONS AS AN EXAMPLE
var Funcs = map[string]func(map[string]string) (interface{}, error){
	"get":    GetObject,
	"set":    SetObject,
	"delete": DeleteObject,
	"exist":  ExistsObjectByKey,
	//  "urlPath": functionName,
}
