package controller

import (
	"fmt"
	"net/http"
	"router/model"
	"strings"
)

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

	res, err := model.RunOperation(operationName, params)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		fmt.Println("Problem with getting access to storage: ", err)
		return
	}
	// fmt.Println(string(res))
	fmt.Fprint(w, string(res))
}
