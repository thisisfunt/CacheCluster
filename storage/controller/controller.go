package controller

import (
	"CacheCluster/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// находим и вызываем операцию
func CallOperation(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = path[1:]
	pathSplit := strings.Split(path, "?")
	if len(pathSplit) == 0 {
		http.Error(w, "", http.StatusBadRequest)
		fmt.Println("Problem with url: cannot parse")
		return
	}
	params := make(map[string]string, 5)
	for k, v := range r.URL.Query() {
		params[k] = v[0]
	}
	operationName := pathSplit[0]
	operation := model.Funcs[operationName]
	if operation == nil {
		http.Error(w, "Operation not found", http.StatusNotFound)
		fmt.Println("Operation not found")
		return
	}
	res, err := operation(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Problem with parsing JSON", http.StatusInternalServerError)
		fmt.Println("Problem with parsing JSON")
		return
	}
	fmt.Fprint(w, string(resJson))
}
