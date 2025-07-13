package model

import (
	"hash/fnv"
	"io"
	"net/http"
	"net/url"
)

// REGISTER YOU HOSTS IN LIST
// YOU CAN USE DEFAULT HOSTS AS AN EXAMPLE
var hosts = []string{"http://storage1:8080", "http://storage2:8080"}

func RunOperation(operationName string, params map[string]string) (string, error) {
	key := params["key"]
	hashAlg := fnv.New32a()
	hashAlg.Write([]byte(key))
	hash := hashAlg.Sum32()
	serverIndex := int(hash) % len(hosts)
	host := hosts[serverIndex]

	baseURL := host + "/" + operationName
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
