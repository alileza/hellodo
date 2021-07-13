package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":9090", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Tamara 👋")
	})); err != nil {
		panic(err)
	}
}
