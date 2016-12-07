package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "&s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/proverb", handler)
	http.ListenAndServe(":3000", nil)
}
