package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

var (
	stateJSON     = ""
	containerJSON = ""
)

func state(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, stateJSON)
}

func containers(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, containerJSON)
}

func init() {
	s, err := ioutil.ReadFile("state.json")
	if err != nil {
		panic(err)
	}
	stateJSON = string(s)

	c, err := ioutil.ReadFile("container.json")
	if err != nil {
		panic(err)
	}
	containerJSON = string(c)
}

func main() {
	http.HandleFunc("/state", state)
	http.HandleFunc("/containers", containers)
	http.ListenAndServe(":5051", nil)
}
